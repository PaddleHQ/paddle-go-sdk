package paddle_test

import (
	"context"
	"fmt"
	paddle "github.com/PaddleHQ/paddle-go-sdk/v5"
	"os"
)

// Demonstrates how to get an existing entity.
func Example_get() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{transaction}}})

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

	// Get a transaction.
	res, err := client.GetTransaction(ctx, &paddle.GetTransactionRequest{
		TransactionID: "txn_01hv8m0mnx3sj85e7gxc6kga03",
	})

	fmt.Println(res.ID, err)
	// Output: txn_01hv8m0mnx3sj85e7gxc6kga03 <nil>
}
