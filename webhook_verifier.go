// ----------------------------------------------------------------------------
// Paddle Webhook Verifier
//
// Added Security Improvements:
// - Timestamp validation to prevent replay attacks (±5 min window)
// - Strict SHA256 hex regex to validate signature format
// - Request body size limit using actual ResponseWriter (prevents silent DoS)
// - Proper rehydration of req.Body for downstream use
// - Clear and specific HTTP status responses on failure 
//
// Based on original Paddle logic
// ----------------------------------------------------------------------------

package paddle

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"strings"
)

var (
	ErrMissingSignature       = errors.New("missing signature")
	ErrInvalidSignatureFormat = errors.New("invalid signature format")
	ErrReplayAttack           = errors.New("timestamp is too old or too far in the future")
)

// enforce a strict signature format: ts=<timestamp>;h1=<64-char hex>
var signatureRegexp = regexp.MustCompile(`^ts=(\d+);h1=([a-f0-9]{64})$`)

// WebhookVerifier validates signed Paddle webhook requests
type WebhookVerifier struct {
	secretKey []byte
}

// NewWebhookVerifier creates a new instance with the provided shared secret.
func NewWebhookVerifier(secretKey string) *WebhookVerifier {
	return &WebhookVerifier{secretKey: []byte(secretKey)}
}

// Verify checks signature, timestamp   and body integrity
// it now accepts the http.ResponseWriter so MaxBytesReader can work properly
func (wv *WebhookVerifier) Verify(w http.ResponseWriter, req *http.Request) (bool, error) {

	sigHeader :=  req.Header.Get("Paddle-Signature") 

	sig := strings.TrimSpace(sigHeader) // trim  spaces

	if sig == "" {
		return false, ErrMissingSignature
	}

	matches := signatureRegexp.FindStringSubmatch(sig)
	if len(matches) != 3 {
		return false, ErrInvalidSignatureFormat
	}

	tsStr := matches[1]
	h1 := matches[2]

	tsInt, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		return false, ErrInvalidSignatureFormat
	}

	// prevent replay attacks with ±5 min timestamp window
	if math.Abs(float64(time.Now().Unix()-tsInt)) > 300 {
		return false, ErrReplayAttack
	}

	const maxBodySize = 2 << 20 // 2 MB

	// Original  (it will be ineffective): as ResponseWriter is set nil, the body limit will  not be enforced
	// limited := http.MaxBytesReader(nil, req.Body, maxBodySize)

	// Fixed: use  ResponseWriter, so that  413 error can be sent if needed
	limited := http.MaxBytesReader(w, req.Body, maxBodySize)

	body, err := io.ReadAll(limited)
	if err != nil {
		return false, err
	}

	// Rehydrate body for downstream handlers 
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	// Construct HMAC using ts:body format as per Paddle docs
	mac := hmac.New(sha256.New, wv.secretKey)
	mac.Write([]byte(tsStr))
	mac.Write([]byte(":"))
	mac.Write(body)

	expectedMAC := mac.Sum(nil)

	receivedMAC, err := hex.DecodeString(h1)
	if err != nil {
		return false, err
	}

	return hmac.Equal(receivedMAC, expectedMAC), nil
}

// Middleware verifies the request before passing it to the next handler.
func (wv *WebhookVerifier) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//  Pass ResponseWriter to Verify, so that the  body size limiting works
		ok, err := wv.Verify(w, r)
		if err != nil {
			switch {
			case errors.Is(err, ErrMissingSignature),
				errors.Is(err, ErrInvalidSignatureFormat),
				errors.Is(err, ErrReplayAttack):
				http.Error(w, err.Error(), http.StatusBadRequest)
			default:
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		if !ok {
			http.Error(w, "signature mismatch", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
