package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to run an explore metrics query.
func Example_metrics_explores_run() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{metricsExplore}}})

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

	res, err := client.RunExploreMetricsQuery(ctx, &paddle.RunExploreMetricsQueryRequest{
		Entity: paddle.MetricEntityTransactionsCompleted,
		From:   "2026-05-01",
		To:     "2026-08-01",
		Measures: []paddle.MetricsExploreMeasure{
			{Field: "gross_revenue", Agg: paddle.PtrTo(paddle.MetricsAggregationSum)},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.Entity, len(res.Series))
	// Output:
	// transactions.completed 1
}
