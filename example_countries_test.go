package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to list countries.
func Example_countries_list() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{countries}}})

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

	res, err := client.ListCountries(ctx, &paddle.ListCountriesRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = res.Iter(ctx, func(v *paddle.CountryDefinition) (bool, error) {
		fmt.Println(v.Iso, v.Name)
		return true, nil
	})
	fmt.Println(err)

	// Output:
	// US UNITED STATES
	// <nil>
}
