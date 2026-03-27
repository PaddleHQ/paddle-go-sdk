package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v5"
)

// Demonstrates how to get monthly recurring revenue metrics.
func Example_metrics_monthly_recurring_revenue() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{metricsMonthlyRecurringRevenue}}})

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

	res, err := client.GetMetricsMonthlyRecurringRevenue(ctx, &paddle.GetMetricsMonthlyRecurringRevenueRequest{
		From: "2025-09-01",
		To:   "2025-09-05",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.CurrencyCode, res.Interval)
	for _, dp := range res.Timeseries {
		fmt.Println(dp.Timestamp, dp.Amount)
	}
	// Output:
	// USD day
	// 2025-09-01T00:00:00Z 1286023068
	// 2025-09-02T00:00:00Z 1345678901
	// 2025-09-03T00:00:00Z 1398765432
	// 2025-09-04T00:00:00Z 1420987654
}
