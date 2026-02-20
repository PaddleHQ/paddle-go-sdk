package paddle_test

import (
	"context"
	"fmt"
	"net/http"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v5"
)

type spyTransport struct {
	called bool
	base   http.RoundTripper
}

// RoundTrip implements the http.RoundTripper interface and spies on usage.
func (s *spyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	s.called = true
	return s.base.RoundTrip(req)
}

// Demonstrates how to get an existing entity.
func Example_client() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{transaction}}})

	// Our custom transport to assert usage.
	spy := &spyTransport{
		base: http.DefaultTransport.(*http.Transport).Clone(),
	}

	// Example of how to create a new HTTP client with custom settings to use within the Paddle SDK.
	httpClient := &http.Client{
		Transport: spy,
	}

	// Create a new Paddle client.
	client, err := paddle.New(
		os.Getenv("PADDLE_API_KEY"),
		paddle.WithBaseURL(s.URL), // Uses the mock server, you will not need this in your integration.
		paddle.WithClient(httpClient),
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

	// Asserting that our custom client transport was used.
	if !spy.called {
		fmt.Println("Expected custom transport to be called")
	}

	fmt.Println(res.ID, err)
	// Output: txn_01hv8m0mnx3sj85e7gxc6kga03 <nil>
}
