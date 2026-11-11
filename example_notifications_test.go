package paddle_test

import (
	"context"
	"fmt"
	"net/http"
	"os"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
)

// Demonstrates how to replay a notification.
func Example_notifications_replay() {
	// Create a mock HTTP server for this example - skip over this bit!
	s := mockServerForExample(mockServerResponse{
		stub:       &stub{paths: []stubPath{notificationReplay}},
		statusCode: http.StatusAccepted,
	})

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

	res, err := client.ReplayNotification(ctx, &paddle.ReplayNotificationRequest{
		NotificationID: "ntf_01h46h1s2zabpkdks7yt4vkgkc",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res.NotificationID)
	// Output:
	// ntf_01h46h1s2zabpkdks7yt4vkgkc
}
