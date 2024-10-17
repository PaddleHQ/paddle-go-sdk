package paddle_test

import (
	"context"
	"fmt"
	"os"
	"time"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v2"
)

// Demonstrates how to create a new entity.
func Example_create() {
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

	// Create a transaction.
	res, err := client.CreateTransaction(ctx, &paddle.CreateTransactionRequest{
		Items: []paddle.CreateTransactionItems{
			*paddle.NewCreateTransactionItemsTransactionItemFromCatalog(&paddle.TransactionItemFromCatalog{
				Quantity: 20,
				PriceID:  "pri_01gsz91wy9k1yn7kx82aafwvea",
			}),
			*paddle.NewCreateTransactionItemsTransactionItemFromCatalog(&paddle.TransactionItemFromCatalog{
				Quantity: 1,
				PriceID:  "pri_01gsz96z29d88jrmsf2ztbfgjg",
			}),
			*paddle.NewCreateTransactionItemsTransactionItemFromCatalog(&paddle.TransactionItemFromCatalog{
				Quantity: 1,
				PriceID:  "pri_01gsz98e27ak2tyhexptwc58yk",
			}),
		},
		CustomerID:     paddle.PtrTo("ctm_01hv6y1jedq4p1n0yqn5ba3ky4"),
		AddressID:      paddle.PtrTo("add_01hv8gq3318ktkfengj2r75gfx"),
		CurrencyCode:   paddle.PtrTo(paddle.CurrencyCodeUSD),
		CollectionMode: paddle.PtrTo(paddle.CollectionModeManual),
		BillingDetails: &paddle.BillingDetails{
			EnableCheckout:      false,
			PurchaseOrderNumber: "PO-123",
			PaymentTerms: paddle.Duration{
				Interval:  paddle.IntervalDay,
				Frequency: 14,
			},
		},
		BillingPeriod: &paddle.TimePeriod{
			StartsAt: time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
			EndsAt:   time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
		},
	})

	fmt.Println(res.ID, err)
	// Output: txn_01hv8m0mnx3sj85e7gxc6kga03 <nil>
}
