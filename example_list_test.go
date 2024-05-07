package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk"
)

// Demonstrates how to fetch a list and iterate over the provided results.
func Example_list() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{transactions}}})

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

	// Get a collection of transactions.
	res, err := client.ListTransactions(ctx, &paddle.ListTransactionsRequest{})

	// Iterate the transactions.
	err = res.Iter(ctx, func(v *paddle.TransactionIncludes) (bool, error) {
		fmt.Println(v.ID)
		return true, nil
	})
	fmt.Println(err)

	// Output:
	//txn_01hv8xxw3etar07vaxsqbyqasy
	//txn_01hv8xbtmb6zc7c264ycteehth
	//txn_01hv8wptq8987qeep44cyrewp9
	//<nil>
}

// Demonstrates how to fetch a list and iterate over the provided results, including the automatic pagination.
func Example_pagination() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{
		transactionsPaginatedPg1,
		transactionsPaginatedPg2,
	}}})

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

	// Get a collection of transactions.
	res, err := client.ListTransactions(ctx, &paddle.ListTransactionsRequest{})

	// Iterate the transactions which will internally paginate to the next page.
	err = res.Iter(ctx, func(v *paddle.TransactionIncludes) (bool, error) {
		fmt.Println(v.ID)
		return true, nil
	})
	fmt.Println(err)

	// Output:
	//txn_01hv8xxw3etar07vaxsqbyqasy
	//txn_01hv8xbtmb6zc7c264ycteehth
	//txn_01hv8wptq8987qeep44cyrewp9
	//txn_01hv8wnvvtedwjrhfhpr9vkq9w
	//txn_01hv8m0mnx3sj85e7gxc6kga03
	//txn_01hv8kxg3hxyxs9t471ms9kfsz
	//<nil>
}
