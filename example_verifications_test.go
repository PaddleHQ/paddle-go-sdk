package paddle_test

import (
	"context"
	"fmt"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to get a verification.
func Example_verifications_get() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{stub: &stub{paths: []stubPath{verification}}})

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

	res, err := client.GetVerification(ctx, &paddle.GetVerificationRequest{
		VerificationID: "slrvrf_01jjr5d4c14h4y3lqlk9bedkw7",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.VerificationType, res.Status)
	// Output:
	// onboarding action_required
}
