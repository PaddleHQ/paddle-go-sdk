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
	// ErrMissingSignature is returned when the signature is missing.
	ErrMissingSignature = errors.New("missing signature")

	// ErrInvalidSignatureFormat is returned when the signature format is invalid.
	ErrInvalidSignatureFormat = errors.New("invalid signature format")

	// ErrReplayAttack is returned when the webhook timestamp is too far in the past or future.
	ErrReplayAttack = errors.New("timestamp is too old or too far in the future")
)

// signatureRegexp matches the Paddle-Signature header format, e.g.:
//
//	ts=1671552777;h1=eb4d0dc8853be92b7f063b9f3ba5233eb920a09459b6e6b2c26705b4364db151
//
// More information can be found here: https://developer.paddle.com/webhooks/signature-verification.
var signatureRegexp = regexp.MustCompile(`^ts=(\d+);h1=([a-f0-9]{64})$`)

// WebhookVerifier is used to verify webhook requests from Paddle.
type WebhookVerifier struct {
	secretKey          []byte
	timestampTolerance *time.Duration
}

// VerifierOption defines a functional option for configuring the verifier.
type VerifierOption func(*WebhookVerifier)

// NewWebhookVerifier creates a new WebhookVerifier with the given secret key and optional config.
func NewWebhookVerifier(secretKey string, opts ...VerifierOption) *WebhookVerifier {
	wv := &WebhookVerifier{secretKey: []byte(secretKey)}
	for _, opt := range opts {
		opt(wv)
	}
	return wv
}

// VerifierWithTimestampTolerance optionally enables replay prevention by enforcing timestamp validity.
// If set, it rejects signatures outside the given time skew.
func VerifierWithTimestampTolerance(d time.Duration) VerifierOption {
	return func(wv *WebhookVerifier) {
		wv.timestampTolerance = &d
	}
}

// Verify verifies the signature and (optionally) the timestamp of a webhook request.
func (wv *WebhookVerifier) Verify(req *http.Request) (bool, error) {
	sig := strings.TrimSpace(req.Header.Get("Paddle-Signature"))
	if sig == "" {
		return false, ErrMissingSignature
	}

	matches := signatureRegexp.FindStringSubmatch(sig)
	if len(matches) != 3 {
		return false, ErrInvalidSignatureFormat
	}

	tsStr := matches[1]
	h1 := matches[2]

	// Optional: validate timestamp for replay protection if tolerance is configured
	if wv.timestampTolerance != nil {
		tsInt, err := strconv.ParseInt(tsStr, 10, 64)
		if err != nil {
			return false, ErrInvalidSignatureFormat
		}
		now := time.Now().Unix()
		if math.Abs(float64(now-tsInt)) > wv.timestampTolerance.Seconds() {
			return false, ErrReplayAttack
		}
	}

	const maxBodySize = 2 << 20 // 2 MB
	req.Body = http.MaxBytesReader(nil, req.Body, maxBodySize)

	body, err := io.ReadAll(req.Body)
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

// Middleware returns a middleware that verifies the signature of a webhook request.
func (wv *WebhookVerifier) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ok, err := wv.Verify(r)
		if err != nil && (errors.Is(err, ErrMissingSignature) || errors.Is(err, ErrInvalidSignatureFormat) || errors.Is(err, ErrReplayAttack)) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !ok {
			http.Error(w, "signature mismatch", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
