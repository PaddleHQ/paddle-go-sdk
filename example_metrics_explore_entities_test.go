package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

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

	res, err := client.ListExploreMetricEntities(ctx, &paddle.ListExploreMetricEntitiesRequest{})
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
