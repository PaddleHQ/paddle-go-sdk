package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v4"
)

// Demonstrates how to update an existing entity.
func Example_update() {
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

	// Update a transaction.
	res, err := client.UpdateTransaction(ctx, &paddle.UpdateTransactionRequest{
		DiscountID: paddle.NewPtrPatchField("dsc_01gtgztp8fpchantd5g1wrksa3"),
		Items: paddle.NewPatchField([]paddle.UpdateTransactionItems{
			*paddle.NewUpdateTransactionItemsTransactionItemFromCatalog(&paddle.TransactionItemFromCatalog{
				Quantity: 50,
				PriceID:  "pri_01gsz91wy9k1yn7kx82aafwvea",
			}),
			*paddle.NewUpdateTransactionItemsTransactionItemFromCatalog(&paddle.TransactionItemFromCatalog{
				Quantity: 1,
				PriceID:  "pri_01gsz96z29d88jrmsf2ztbfgjg",
			}),
			*paddle.NewUpdateTransactionItemsTransactionItemFromCatalog(&paddle.TransactionItemFromCatalog{
				Quantity: 1,
				PriceID:  "pri_01gsz98e27ak2tyhexptwc58yk",
			}),
		}),
	})

	fmt.Println(res.ID, err)
	// Output: txn_01hv8m0mnx3sj85e7gxc6kga03 <nil>
}
