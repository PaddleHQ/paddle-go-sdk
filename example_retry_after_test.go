package paddle_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v4"
	"github.com/PaddleHQ/paddle-go-sdk/v4/pkg/paddleerr"
)

// Demonstrates how to handle rate limiting and retry scenarios using the Retry-After header.
// This example shows a realistic pattern for scheduling retries in a background job queue
// or processing system when the API returns a 429 Too Many Requests response with a
// Retry-After header. The RetryAfter information is available on paddleerr.Error for
// known API errors such as too_many_requests.
func Example_retryAfter() {
	// Create a mock HTTP server that simulates rate limiting with Retry-After.
	s := mockServerForExample(mockServerResponse{
		stub:       &stub{paths: []stubPath{tooManyRequestsError}},
		statusCode: http.StatusTooManyRequests,
		headers: http.Header{
			"Retry-After": []string{"120"},
		},
	})

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

	// Make an API call which will return a 429 response.
	// For this example we'll ignore the response
	_, err = client.GetTransaction(ctx, &paddle.GetTransactionRequest{
		TransactionID: "txn_01hv8m0mnx3sj85e7gxc6kga03",
	})

	var apiErr *paddleerr.Error
	if errors.As(err, &apiErr) && apiErr.RetryAfter != nil {
		fmt.Printf("Error indicates total delay: ~%ds\n", int(apiErr.RetryAfter.TotalDelay().Seconds()))
		fmt.Printf("RetryAt expired check: %t\n", apiErr.RetryAfter.IsExpired())

		// Here you would add logic to retry this call later, e.g. re-enqueuing with a delay.
		// You can utilise RetryAfter.WaitTime() to get a relative delay from when the request was parsed.
	}

	// Output:
	// Error indicates total delay: ~120s
	// RetryAt expired check: false
}
