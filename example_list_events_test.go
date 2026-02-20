package paddle_test

import (
	"context"
	"encoding/json"
	"fmt"
	paddle "github.com/PaddleHQ/paddle-go-sdk/v5"
	"os"
)

// Demonstrates how to fetch a list of events and iterate over the provided results.
func Example_listEvents() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{events}}})

	// Create a new Paddle client.
	client, err := paddle.New(
		os.Getenv("PADDLE_API_KEY"),
		paddle.WithBaseURL(s.URL), // Uses the mock server, you will not need this in your integration.
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	// Optionally set a transit ID on the context. This is useful to link your
	// own request IDs to Paddle API requests.
	ctx = paddle.ContextWithTransitID(ctx, "sdk-testing-request-1")

	// Get a collection of events.
	res, err := client.ListEvents(ctx, &paddle.ListEventsRequest{})

	// Iterate the events.
	err = res.Iter(ctx, func(e paddle.Event) (bool, error) {
		switch v := e.(type) {
		case *paddle.TransactionCompletedEvent:
			// here v could be used as concrete type TransactionCompletedEvent
			fmt.Println(v.EventID)
			fmt.Println(v.Data.ID)
		case *paddle.TransactionUpdatedEvent:
			// here v could be used as concrete type TransactionUpdatedEvent
			fmt.Println(v.EventID)
			fmt.Println(v.Data.ID)
		default:
			// Unhandled event, we could log and error or covert to GenericEvent
			ge, err := toGenericEvent(v)
			if err != nil {
				return false, err
			}
			fmt.Println(ge.EventID)
		}

		return true, nil
	})
	fmt.Println(err)

	// Output:
	//evt_01hywqk7y8qfzj69z3pdvz34qt
	//txn_01hywqfe6yxhxcsfb4ays8mqt3
	//evt_01hywqfn8b1na40vyarxaxqa9t
	//txn_01hywqfervkjg6hkk035wy24gt
	//evt_01hv9771tccgcm4y810d8zbceh
	//<nil>
}

func toGenericEvent(e paddle.Event) (ge *paddle.GenericEvent, err error) {
	t, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(t, &ge)
	if err != nil {
		return nil, err
	}

	return ge, nil
}
