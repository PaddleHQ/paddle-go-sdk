package paddle_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v5"
	"github.com/PaddleHQ/paddle-go-sdk/v5/pkg/paddlenotification"
)

const (
	exampleWebhookSignature = `ts=1710929255;h1=6c05ef8fa83c44d751be6d259ec955ce5638e2c54095bf128e408e2fce1589c8`
	exampleWebhookPayload   = `{"data":{"id":"pri_01hsdn96k2hxjzsq5yerecdj9j","name":null,"status":"active","quantity":{"maximum":999999,"minimum":1},"tax_mode":"account_setting","product_id":"pro_01hsdn8qp7yydry3x1yeg6a9rv","unit_price":{"amount":"1000","currency_code":"USD"},"custom_data":null,"description":"testing","import_meta":null,"trial_period":null,"billing_cycle":{"interval":"month","frequency":1},"unit_price_overrides":[]},"event_id":"evt_01hsdn97563968dy0szkmgjwh3","event_type":"price.created","occurred_at":"2024-03-20T10:07:35.590857Z","notification_id":"ntf_01hsdn977e920kbgzt6r6c9rqc"}`
	exampleWebhookSecretKey = `pdl_ntfset_01hsdn8d43dt7mezr1ef2jtbaw_hKkRiCGyyRhbFwIUuqiTBgI7gnWoV0Gr`
)

// Demonstrates how to unmarshal webhooks to their notification type
func Example_webhookUnmarshal() {
	// Create a WebhookVerifier with your secret key
	// You should keep your secret outside the src, e.g. as an env variable
	verifier := paddle.NewWebhookVerifier(exampleWebhookSecretKey)

	// Webhook is a small definition of what we want to initially read before processing the entire payload
	type Webhook struct {
		EventID   string                           `json:"event_id"`
		EventType paddlenotification.EventTypeName `json:"event_type"`
	}

	// We're utilising the Middleware verification method
	handler := verifier.Middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The webhook is verified at this point, we can safely process it
		defer r.Body.Close()

		rawBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		// Initially read only the event id and event type from the request
		type Webhook struct {
			EventID   string                           `json:"event_id"`
			EventType paddlenotification.EventTypeName `json:"event_type"`
		}

		var webhook Webhook

		if err := json.Unmarshal(rawBody, &webhook); err != nil {
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		// Optionally check you've not processed this event_id before in your system

		// Handle each notification based on the webhook.EventType
		// In this case we're going to return a string ID from the corresponding notification type
		var entityID string

		switch webhook.EventType {
		case paddlenotification.EventTypeNameAddressCreated:
			address := &paddlenotification.AddressCreated{}
			if err := json.Unmarshal(rawBody, address); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// You can safely proceed with address as a paddlenotification.AddressCreated type
			entityID = address.Data.ID
		case paddlenotification.EventTypeNamePriceCreated:
			price := &paddlenotification.PriceCreated{}
			if err := json.Unmarshal(rawBody, price); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// You can safely proceed with address as a paddlenotification.PriceCreated type
			entityID = price.Data.ID
		default:
			generic := &paddlenotification.GenericNotificationEvent{}
			if err := json.Unmarshal(rawBody, generic); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			// You can attempt to process this even though the event type was not recognised in the switch
			// In this case we'll simply respond with the event id
			entityID = generic.EventID
		}

		// Respond as soon as possible with a 200 OK
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"ID": "%s"}`, entityID)))
	}))

	// We're simulating a call to the server, everything below can be skipped in your implementation

	req, err := http.NewRequest(http.MethodPost, "localhost:8081", strings.NewReader(exampleWebhookPayload))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Paddle-Signature", exampleWebhookSignature)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	body, err := io.ReadAll(rr.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body), err)
	// Output: {"ID": "pri_01hsdn96k2hxjzsq5yerecdj9j"} <nil>
}
