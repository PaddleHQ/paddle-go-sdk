package paddle_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v2"
)

const (
	exampleSignature = `ts=1710929255;h1=6c05ef8fa83c44d751be6d259ec955ce5638e2c54095bf128e408e2fce1589c8`
	examplePayload   = `{"data":{"id":"pri_01hsdn96k2hxjzsq5yerecdj9j","name":null,"status":"active","quantity":{"maximum":999999,"minimum":1},"tax_mode":"account_setting","product_id":"pro_01hsdn8qp7yydry3x1yeg6a9rv","unit_price":{"amount":"1000","currency_code":"USD"},"custom_data":null,"description":"testing","import_meta":null,"trial_period":null,"billing_cycle":{"interval":"month","frequency":1},"unit_price_overrides":[]},"event_id":"evt_01hsdn97563968dy0szkmgjwh3","event_type":"price.created","occurred_at":"2024-03-20T10:07:35.590857Z","notification_id":"ntf_01hsdn977e920kbgzt6r6c9rqc"}`
	exampleSecretKey = `pdl_ntfset_01hsdn8d43dt7mezr1ef2jtbaw_hKkRiCGyyRhbFwIUuqiTBgI7gnWoV0Gr`
)

// Demonstrates how to verify a webhook.
func ExampleWebhookVerifier_Verify() {
	// Create a WebhookVerifier with your secret key
	// You should keep your secret outside the src, e.g. as an env variable
	verifier := paddle.NewWebhookVerifier(exampleSecretKey)

	// Create a server to receive the webhook
	// Note: you may have this in place already
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request with the verifier
		ok, err := verifier.Verify(r)

		// Check no error occurred during verification and return an appropriate response
		if err != nil && (errors.Is(err, paddle.ErrMissingSignature) || errors.Is(err, paddle.ErrInvalidSignatureFormat)) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if verification was successful
		if !ok {
			// Return an appropriate response to let Paddle know
			http.Error(w, "signature mismatch", http.StatusForbidden)
			return
		}

		// Best practice here is to check if you have processed this webhook already using the event id
		// At this point you should store for async processing
		// For example a local queue or db entry

		// Respond as soon as possible with a 200 OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))

	// We're simulating a call to the server, everything below can be skipped in your implementation

	req, err := http.NewRequest(http.MethodPost, s.URL, strings.NewReader(examplePayload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Paddle-Signature", exampleSignature)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	fmt.Println(string(body), err)
	// Output: {"success": true} <nil>
}

// Demonstrates how to use the WebhookVerifier as a middleware.
func ExampleWebhookVerifier_Middleware() {
	// Create a WebhookVerifier with your secret key
	// You should keep your secret outside the src, e.g. as an env variable
	verifier := paddle.NewWebhookVerifier(exampleSecretKey)

	// Wrap your handler with the verifier.Middleware method
	handler := verifier.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The request making it this far means the webhook was verified
		// Best practice here is to check if you have processed this webhook already using the event id
		// At this point you should store for async processing
		// For example a local queue or db entry

		// Respond as soon as possible with a 200 OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"success": true}`))
	}))

	// We're simulating a call to the server, everything below can be skipped in your implementation

	req, err := http.NewRequest(http.MethodPost, "localhost:8081", strings.NewReader(examplePayload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Paddle-Signature", exampleSignature)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body), err)
	// Output: {"success": true} <nil>
}
