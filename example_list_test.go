package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v3"
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
	err = res.Iter(ctx, func(v *paddle.Transaction) (bool, error) {
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
	err = res.Iter(ctx, func(v *paddle.Transaction) (bool, error) {
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

func Example_listWithoutPagination() {
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{
		event_types,
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
	// Get a collection of transactions.
	res, err := client.ListEventTypes(ctx, &paddle.ListEventTypesRequest{})

	// Iterate the transactions which will internally paginate to the next page.
	err = res.Iter(ctx, func(v *paddle.EventType) (bool, error) {
		fmt.Println(v.Name)
		return true, nil
	})
	fmt.Println(err)

	// Output:
	//transaction.billed
	//transaction.canceled
	//transaction.completed
	//transaction.created
	//transaction.paid
	//transaction.past_due
	//transaction.payment_failed
	//transaction.ready
	//transaction.updated
	//subscription.activated
	//subscription.canceled
	//subscription.created
	//subscription.imported
	//subscription.past_due
	//subscription.paused
	//subscription.resumed
	//subscription.trialing
	//subscription.updated
	//product.created
	//product.imported
	//product.updated
	//price.created
	//price.imported
	//price.updated
	//customer.created
	//customer.imported
	//customer.updated
	//address.created
	//address.imported
	//address.updated
	//business.created
	//business.imported
	//business.updated
	//adjustment.created
	//adjustment.updated
	//payout.created
	//payout.paid
	//discount.created
	//discount.imported
	//discount.updated
	//report.created
	//report.updated
	//<nil>
}
