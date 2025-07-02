package paddle_test

import (
	"context"
	"fmt"
	"os"
	"time"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v4"
	"github.com/PaddleHQ/paddle-go-sdk/v4/pkg/paddlenotification"
)

// Demonstrates how to create a Simulation with Payload and read the Payload back out of the response
func Example_simulation_create() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulation}}})

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

	simulation, err := client.CreateSimulation(ctx, paddle.NewCreateSimulationRequestSimulationSingleEventCreate(&paddle.SimulationSingleEventCreate{
		NotificationSettingID: "ntfset_01j84xydheq48n3btebwf6ndn6",
		Name:                  "Go SDK Test with Payload",
		Type:                  "customer.created",
		Payload: paddlenotification.CustomerNotification{
			ID:               "ctm_01j870snka0xdp6szgyxze6d6d",
			Name:             paddle.PtrTo("John Doe"),
			Email:            "john.doe@paddle.com",
			MarketingConsent: false,
			Status:           paddlenotification.StatusActive,
			CustomData:       nil,
			Locale:           "en",
			CreatedAt:        time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
			UpdatedAt:        time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
			ImportMeta:       nil,
		},
	}))
	if err != nil {
		return
	}

	// Use type assertion on the response for payload
	payload, ok := (simulation.Payload).(*paddlenotification.CustomerNotification)
	if !ok {
		fmt.Println("Payload is unexpected type")
		return
	}

	fmt.Println(simulation.ID)
	fmt.Println(payload.ID)
	// Output:
	//ntfsim_01j9y0jwekrcyezscgkehvdmd6
	//ctm_01j870snka0xdp6szgyxze6d6d
}

// Demonstrates how to create a scenario Simulation
func Example_simulation_create_scenario_simulation() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulationScenario}}})

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

	simulation, err := client.CreateSimulation(ctx, paddle.NewCreateSimulationRequestSimulationScenarioCreate(&paddle.SimulationScenarioCreate{
		NotificationSettingID: "ntfset_01j84xydheq48n3btebwf6ndn6",
		Name:                  "Go SDK Test without Payload",
		Type:                  "subscription_creation",
		Config: paddle.NewSimulationScenarioCreateConfigSimulationSubscriptionCreation(&paddle.SimulationSubscriptionCreation{
			SubscriptionCreation: paddle.SimulationSubscriptionCreationConfigCreate{
				Entities: *paddle.NewSimulationSubscriptionCreationConfigCreateEntitiesSimulationSubscriptionCreationConfigNoPrices(
					&paddle.SimulationSubscriptionCreationConfigNoPrices{
						CustomerID: paddle.PtrTo("ctm_01j870snka0xdp6szgyxze6d6d"),
					},
				),
				Options: paddle.SimulationSubscriptionCreationConfigOptions{
					CustomerSimulatedAs: paddle.CustomerSimulatedAsNew,
				},
			},
		}),
	}))
	if err != nil {
		return
	}

	fmt.Println(simulation.ID)
	fmt.Println(simulation.Type)
	fmt.Println(simulation.Name)
	fmt.Println(simulation.NotificationSettingID)
	fmt.Println(simulation.Payload)
	// Output:
	//ntfsim_01j9y0jwekrcyezscgkehvdmd6
	//subscription_creation
	//Go SDK Test without Payload
	//ntfset_01j84xydheq48n3btebwf6ndn6
	//<nil>
}

// Demonstrates how to update a Simulation with Payload and read the Payload back out of the response
func Example_simulation_update() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulation}}})

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

	simulation, err := client.UpdateSimulation(
		ctx,
		paddle.NewUpdateSimulationRequestSimulationSingleEventUpdate(
			"ntfsim_01j9y0jwekrcyezscgkehvdmd6",
			&paddle.SimulationSingleEventUpdate{
				NotificationSettingID: paddle.NewPatchField("ntfset_01j84xydheq48n3btebwf6ndn6"),
				Name:                  paddle.NewPatchField("Go SDK Test with Payload"),
				Type:                  paddle.NewPatchField(paddle.EventTypeNameCustomerCreated),

				Payload: paddle.NewPtrPatchField(paddlenotification.NotificationPayload(&paddlenotification.CustomerNotification{
					ID:               "ctm_01j870snka0xdp6szgyxze6d6d",
					Name:             paddle.PtrTo("John Doe"),
					Email:            "jane.doe@paddle.com",
					MarketingConsent: false,
					Status:           paddlenotification.StatusActive,
					CustomData:       nil,
					Locale:           "en",
					CreatedAt:        time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
					UpdatedAt:        time.Date(2024, 4, 12, 0, 0, 0, 0, time.UTC).Format(time.RFC3339),
					ImportMeta:       nil,
				})),
			},
		),
	)
	if err != nil {
		return
	}

	// Use type assertion on the response for payload
	payload, ok := (simulation.Payload).(*paddlenotification.CustomerNotification)
	if !ok {
		fmt.Println("Payload is unexpected type")
		return
	}

	fmt.Println(simulation.ID)
	fmt.Println(payload.ID)
	// Output:
	//ntfsim_01j9y0jwekrcyezscgkehvdmd6
	//ctm_01j870snka0xdp6szgyxze6d6d
}

// // Demonstrates how to list Simulations with Payload and read the Payload back out of the response
func Example_simulation_list() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulations}}})

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

	simulations, err := client.ListSimulations(ctx, &paddle.ListSimulationsRequest{})
	if err != nil {
	}

	simulations.Iter(ctx, func(s *paddle.Simulation) (bool, error) {
		switch p := s.Payload.(type) {
		case *paddlenotification.AddressNotification:
			// here v could be used as concrete type AddressNotification
			fmt.Println(p.CustomerID)
		case *paddlenotification.CustomerNotification:
			// here v could be used as concrete type CustomerNotification
			fmt.Println(p.Email)
		}

		return true, nil
	})
	// Output:
	//john.doe+blackhole@paddle.com
	//ctm_01hv6y1jedq4p1n0yqn5ba3ky4
}

// Demonstrates how to get a Simulation with Payload and read the Payload back out of the response
func Example_simulation_get() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulation}}})

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

	simulation, err := client.GetSimulation(ctx, &paddle.GetSimulationRequest{SimulationID: "ntfsim_01j9y0jwekrcyezscgkehvdmd6"})
	if err != nil {
		return
	}

	// Use type assertion on the response for payload
	payload, ok := (simulation.Payload).(*paddlenotification.CustomerNotification)
	if !ok {
		fmt.Println("Payload is unexpected type")
		return
	}

	fmt.Println(simulation.ID)
	fmt.Println(payload.ID)
	// Output:
	//ntfsim_01j9y0jwekrcyezscgkehvdmd6
	//ctm_01j870snka0xdp6szgyxze6d6d
}

// Demonstrates how to run a Simulation
func Example_simulation_run() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulationRun}}})

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

	simulationRun, err := client.CreateSimulationRun(ctx, &paddle.CreateSimulationRunRequest{
		SimulationID: "ntfsim_01j9y0jwekrcyezscgkehvdmd6",
	})
	if err != nil {
		return
	}

	fmt.Println(simulationRun.ID)
	// Output:
	//ntfsimrun_01j9yq3yspewy5r8zr05vkeekd
}

// Demonstrates how to get a SimulationRun with included SimulationRunEvents
func Example_simulation_run_get() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{simulationRunWithEvents}}})

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

	simulationRun, err := client.GetSimulationRun(ctx, &paddle.GetSimulationRunRequest{
		SimulationID:    "ntfsim_01j9y0jwekrcyezscgkehvdmd6",
		SimulationRunID: "ntfsimrun_01j9yq3yspewy5r8zr05vkeekd",
		IncludeEvents:   true,
	})
	if err != nil {
		return
	}

	// Use type assertion on the response for payload
	for _, event := range simulationRun.Events {
		payload, ok := (event.Payload).(*paddlenotification.CustomerNotification)
		if !ok {
			fmt.Println("Payload is unexpected type")
			return
		}
		fmt.Println(payload.ID)
		fmt.Println(event.Response.StatusCode)
	}

	fmt.Println(simulationRun.ID)
	// Output:
	//ctm_01j870snka0xdp6szgyxze6d6d
	//200
	//ntfsimrun_01j9yq3yspewy5r8zr05vkeekd
}
