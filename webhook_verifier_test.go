package paddle_test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	testPayload = `{"data":{"id":"pri_01hsdn96k2hxjzsq5yerecdj9j","name":null,"status":"active","quantity":{"maximum":999999,"minimum":1},"tax_mode":"account_setting","product_id":"pro_01hsdn8qp7yydry3x1yeg6a9rv","unit_price":{"amount":"1000","currency_code":"USD"},"custom_data":null,"description":"testing","import_meta":null,"trial_period":null,"billing_cycle":{"interval":"month","frequency":1},"unit_price_overrides":[]},"event_id":"evt_01hsdn97563968dy0szkmgjwh3","event_type":"price.created","occurred_at":"2024-03-20T10:07:35.590857Z","notification_id":"ntf_01hsdn977e920kbgzt6r6c9rqc"}`
	testSecretKey = `pdl_ntfset_01hsdn8d43dt7mezr1ef2jtbaw_hKkRiCGyyRhbFwIUuqiTBgI7gnWoV0Gr`
	testSignature string
)

func init() {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	raw := timestamp + ":" + testPayload
	mac := hmac.New(sha256.New, []byte(testSecretKey))
	mac.Write([]byte(raw))
	digest := hex.EncodeToString(mac.Sum(nil))
	testSignature = fmt.Sprintf("ts=%s;h1=%s", timestamp, digest)
}

func TestVerifier_Verify(t *testing.T) {
	type args struct {
		name           string
		payload        string
		signature      string
		secretKey      string
		expectedResult bool
		expectedError  error
	}

	// Simulate a replay attack
	oldTS := strconv.FormatInt(time.Now().Add(-10*time.Minute).Unix(), 10)
	raw := oldTS + ":" + testPayload
	mac := hmac.New(sha256.New, []byte(testSecretKey))
	mac.Write([]byte(raw))
	replaySig := hex.EncodeToString(mac.Sum(nil))
	replayHeader := fmt.Sprintf("ts=%s;h1=%s", oldTS, replaySig)

	cases := []args{
		{
			name:           "valid signature",
			payload:        testPayload,
			signature:      testSignature,
			secretKey:      testSecretKey,
			expectedResult: true,
		},
		{
			name:           "missing signature",
			payload:        testPayload,
			secretKey:      testSecretKey,
			expectedResult: false,
			expectedError:  paddle.ErrMissingSignature,
		},
		{
			name:           "invalid payload",
			payload:        `{}`,
			signature:      testSignature,
			secretKey:      testSecretKey,
			expectedResult: false,
		},
		{
			name:           "invalid signature format",
			payload:        testPayload,
			signature:      `ts=x;h1=y`,
			secretKey:      testSecretKey,
			expectedResult: false,
			expectedError:  paddle.ErrInvalidSignatureFormat,
		},
		{
			name:           "invalid secret key",
			payload:        testPayload,
			signature:      testSignature,
			secretKey:      `wrong`,
			expectedResult: false,
		},
		{
			name:           "replay attack",
			payload:        testPayload,
			signature:      replayHeader,
			secretKey:      testSecretKey,
			expectedResult: false,
			expectedError:  paddle.ErrReplayAttack,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.payload))
			if c.signature != "" {
				req.Header.Set("Paddle-Signature", c.signature)

				t.Logf("🔏 Signature: %s", c.signature)
				parts := strings.Split(c.signature, ";")
				if len(parts) > 0 && strings.HasPrefix(parts[0], "ts=") {
					t.Logf("🔐 Timestamp: %s", parts[0][3:])
				}
			}
			t.Logf("📦 Payload size: %d bytes", len(c.payload))

			rr := httptest.NewRecorder()
			ok, err := paddle.NewWebhookVerifier(c.secretKey).Verify(rr, req)

			assert.Equal(t, c.expectedResult, ok)
			if c.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, c.expectedError)
			}

			// Confirm body is reusable
			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)
			assert.Equal(t, c.payload, string(body))
		})
	}
}

func TestVerifier_Middleware(t *testing.T) {
	type args struct {
		name           string
		payload        string
		signature      string
		secretKey      string
		expectedStatus int
		expectedBody   string
	}

	cases := []args{
		{
			name:           "valid signature",
			payload:        testPayload,
			signature:      testSignature,
			secretKey:      testSecretKey,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "missing signature",
			payload:        testPayload,
			secretKey:      testSecretKey,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing signature\n",
		},
		{
			name:           "invalid payload",
			payload:        `{}`,
			signature:      testSignature,
			secretKey:      testSecretKey,
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   "signature mismatch\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.payload))
			req.Header.Set("Paddle-Signature", c.signature)

			t.Logf("🧪 Middleware: %s", c.name)
			t.Logf("📦 Payload size: %d", len(c.payload))

			rr := httptest.NewRecorder()
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			paddle.NewWebhookVerifier(c.secretKey).Middleware(next).ServeHTTP(rr, req)

			assert.Equal(t, c.expectedStatus, rr.Code)
			assert.Equal(t, c.expectedBody, rr.Body.String())
		})
	}
}

func TestVerifier_RejectsOversizedPayload(t *testing.T) {
	secret := testSecretKey
	largePayload := strings.Repeat("A", 3<<20) // 3MB

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	raw := timestamp + ":" + largePayload
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(raw))
	sig := hex.EncodeToString(mac.Sum(nil))
	header := fmt.Sprintf("ts=%s;h1=%s", timestamp, sig)

	t.Logf("🧱 Oversized Payload: %d bytes", len(largePayload))
	t.Logf("🔐 Timestamp: %s", timestamp)
	t.Logf("🔏 Signature (truncated): %s...", header[:80])

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(largePayload))
	req.Header.Set("Paddle-Signature", header)
	rr := httptest.NewRecorder()

	ok, err := paddle.NewWebhookVerifier(secret).Verify(rr, req)

	assert.False(t, ok)
	require.Error(t, err)
	assert.EqualError(t, err, "http: request body too large")
}
