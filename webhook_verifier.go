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
	"strings"
	"time"
)

var (
	ErrMissingSignature       = errors.New("missing signature")
	ErrInvalidSignatureFormat = errors.New("invalid signature format")
	ErrReplayAttack           = errors.New("timestamp is too old or too far in the future")
)

// enforce a strict signature format: ts=<timestamp>;h1=<64-char hex>
var signatureRegexp = regexp.MustCompile(`^ts=(\d+);h1=([a-f0-9]{64})$`)

const TOLERANCE_DEFAULT_VALUE = 5 * time.Minute

// WebhookVerifier validates signed Paddle webhook requests
type WebhookVerifier struct {
	secretKey []byte
	tolerance time.Duration
	resWriter http.ResponseWriter
	rpFlag	  bool
}

type VerifierOption func(*WebhookVerifier)

// NewWebhookVerifier creates a new instance with the provided shared secret.
func NewWebhookVerifier(secretKey string, opts ...VerifierOption) *WebhookVerifier {
	wv := &WebhookVerifier{
		secretKey: []byte(secretKey),
		tolerance: TOLERANCE_DEFAULT_VALUE,
		rpFlag:    true,
	}
	for _, opt := range opts {
		opt(wv)
	}
	return wv
}

// VerifierWithResponseWriter optionally enables response writer support for MaxBytesReader.
func VerifierWithResponseWriter(w http.ResponseWriter) VerifierOption {
	return func(wv *WebhookVerifier) {
		wv.resWriter = w
	}
}

func VerifierWithReplayPrevention(flag bool) VerifierOption {
	return func(wv *WebhookVerifier){
		wv.rpFlag = flag
	}
}

// VerifierWithTimestampTolerance optionally sets the allowed timestamp skew for replay prevention.
func VerifierWithTimestampTolerance(d time.Duration) VerifierOption {
	return func(wv *WebhookVerifier) {
		wv.tolerance = d
	}
}

// Verify checks signature, timestamp, and body integrity.
func (wv *WebhookVerifier) Verify(req *http.Request) (bool, error) {
	sigHeader := strings.TrimSpace(req.Header.Get("Paddle-Signature"))
	if sigHeader == "" {
		return false, ErrMissingSignature
	}

	matches := signatureRegexp.FindStringSubmatch(sigHeader)
	if len(matches) != 3 {
		return false, ErrInvalidSignatureFormat
	}

	tsStr := matches[1]
	h1 := matches[2]

	tsInt, err := strconv.ParseInt(tsStr, 10, 64)
	if err != nil {
		return false, ErrInvalidSignatureFormat
	}

	now := time.Now().Unix()
	if math.Abs(float64(now-tsInt)) > wv.tolerance.Seconds() {
		return false, ErrReplayAttack
	}

	const maxBodySize = 2 << 20 // 2 MB
	limited := http.MaxBytesReader(wv.resWriter, req.Body, maxBodySize)

	body, err := io.ReadAll(limited)
	if err != nil {
		return false, err
	}

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
		// Set the ResponseWriter for body size limiting
		wv.resWriter = w

		ok, err := wv.Verify(r)
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
