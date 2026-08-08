package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to get a checkout domain.
func Example_checkout_domains_get() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{checkoutDomain}}})

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

	res, err := client.GetCheckoutDomain(ctx, &paddle.GetCheckoutDomainRequest{
		DomainID: "chedom_01kkertpke0gv2t61p1pq8v23x",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Domain, res.Status)
	// Output:
	// example.com approved
}
