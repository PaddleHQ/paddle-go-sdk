package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
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

	res, err := client.ExploreMetrics(ctx, &paddle.ExploreMetricsRequest{
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

// Demonstrates how to list explore metric entities.
func Example_metrics_explore_entities_list() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{metricsExploreEntities}}})

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

	res, err := client.ListExploreEntities(ctx, &paddle.ListExploreEntitiesRequest{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = res.Iter(ctx, func(v *paddle.MetricEntityDefinition) (bool, error) {
		fmt.Println(v.Entity, v.TimeDimension.MaxLookbackDays)
		return true, nil
	})
	fmt.Println(err)

	// Output:
	// subscriptions 1096
	// <nil>
}
