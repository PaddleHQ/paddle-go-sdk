package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to list currencies.
func Example_currencies_list() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{currencies}}})

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

	res, err := client.ListCurrencies(ctx, &paddle.ListCurrenciesRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = res.Iter(ctx, func(v *paddle.CurrencyDefinition) (bool, error) {
		fmt.Println(v.Code, v.Name)
		return true, nil
	})
	fmt.Println(err)

	// Output:
	// GBP Pound Sterling
	// <nil>
}
