package paddle_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testSignature = `ts=1710929255;h1=6c05ef8fa83c44d751be6d259ec955ce5638e2c54095bf128e408e2fce1589c8`
	testPayload   = `{"data":{"id":"pri_01hsdn96k2hxjzsq5yerecdj9j","name":null,"status":"active","quantity":{"maximum":999999,"minimum":1},"tax_mode":"account_setting","product_id":"pro_01hsdn8qp7yydry3x1yeg6a9rv","unit_price":{"amount":"1000","currency_code":"USD"},"custom_data":null,"description":"testing","import_meta":null,"trial_period":null,"billing_cycle":{"interval":"month","frequency":1},"unit_price_overrides":[]},"event_id":"evt_01hsdn97563968dy0szkmgjwh3","event_type":"price.created","occurred_at":"2024-03-20T10:07:35.590857Z","notification_id":"ntf_01hsdn977e920kbgzt6r6c9rqc"}`
	testSecretKey = `pdl_ntfset_01hsdn8d43dt7mezr1ef2jtbaw_hKkRiCGyyRhbFwIUuqiTBgI7gnWoV0Gr`
)

func TestVerifier_Verify(t *testing.T) {
	type args struct {
		name      string
		payload   string
		signature string
		secretKey string

		expectedResult bool
		expectedError  error
	}

	cases := []args{
		{
			name: "valid signature",

			payload:   testPayload,
			signature: testSignature,
			secretKey: testSecretKey,

			expectedResult: true,
		},
		{
			name: "missing signature",

			payload:   testPayload,
			secretKey: testSecretKey,

			expectedResult: false,
			expectedError:  paddle.ErrMissingSignature,
		},
		{
			name: "invalid payload",

			payload:   `{}`,
			signature: testSignature,
			secretKey: testSecretKey,

			expectedResult: false,
		},
		{
			name: "invalid signature format",

			payload:   testPayload,
			signature: `ts=x;h1=y`,
			secretKey: testSecretKey,

			expectedResult: false,
			expectedError:  paddle.ErrInvalidSignatureFormat,
		},
		{
			name: "invalid secret key",

			payload:   testPayload,
			signature: testSignature,
			secretKey: `abc`,

			expectedResult: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.payload))
			req.Header.Set("Paddle-Signature", c.signature)

			ok, err := paddle.NewWebhookVerifier(c.secretKey).Verify(req)
			assert.Equal(t, c.expectedResult, ok)
			if c.expectedError == nil {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, c.expectedError)
			}

			// Assert that the body was put back onto the request after reading:
			body, err := io.ReadAll(req.Body)
			require.NoError(t, err)

			assert.Equal(t, c.payload, string(body))
		})
	}
}

func TestVerifier_Middleware(t *testing.T) {
	type args struct {
		name string

		payload   string
		signature string
		secretKey string

		expectedStatus int
		expectedBody   string
	}

	cases := []args{
		{
			name: "valid signature",

			payload:   testPayload,
			signature: testSignature,
			secretKey: testSecretKey,

			expectedStatus: http.StatusOK,
		},
		{
			name: "missing signature",

			payload:   testPayload,
			secretKey: testSecretKey,

			expectedStatus: http.StatusBadRequest,
			expectedBody:   "missing signature\n",
		},
		{
			name: "invalid payload",

			payload:   `{}`,
			signature: testSignature,
			secretKey: testSecretKey,

			expectedStatus: http.StatusForbidden,
			expectedBody:   "signature mismatch\n",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(c.payload))
			req.Header.Set("Paddle-Signature", c.signature)

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
