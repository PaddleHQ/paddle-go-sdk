package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to list history entries for a subscription and work with the
// action-specific detail variants.
func Example_subscription_history_list() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{subscriptionHistory}}})

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

	// Get a collection of history entries for a subscription.
	res, err := client.ListSubscriptionHistory(ctx, &paddle.ListSubscriptionHistoryRequest{
		SubscriptionID: "sub_01hv959anj4zrw503h2acawb3p",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Iterate the history entries. The Detail field holds a concrete variant
	// selected by the entry's action.
	err = res.Iter(ctx, func(v *paddle.SubscriptionHistory) (bool, error) {
		switch d := v.Detail.(type) {
		case *paddle.SubscriptionHistoryDetailSubscriptionCanceled:
			fmt.Println(d.Action, *d.EffectiveFrom)
		case *paddle.SubscriptionHistoryDetailSubscriptionActivated:
			fmt.Println(d.Action, d.TransactionID)
		case map[string]any:
			// Actions this SDK version doesn't know about preserve their raw data.
			fmt.Println("unknown:", d["action"])
		}
		return true, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.EstimatedTotal(), res.HasMore())
	// Output:
	// subscription_canceled next_billing_period
	// subscription_activated txn_01h89231k3a1q8mn6p4r5s7t9v
	// unknown: subscription_shipping_address_verified
	// 3 false
}
