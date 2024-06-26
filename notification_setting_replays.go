// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import "context"

// NotificationSettingReplaysClient is a client for the Notification setting replays resource.
type NotificationSettingReplaysClient struct {
	doer Doer
}

// ReplayNotificationRequest is given as an input to ReplayNotification.
type ReplayNotificationRequest struct {
	// URL path parameters.
	NotificationID string `in:"path=notification_id" json:"-"`
}

// ReplayNotification performs the POST operation on a Notification setting replays resource.
func (c *NotificationSettingReplaysClient) ReplayNotification(ctx context.Context, req *ReplayNotificationRequest) (err error) {
	if err := c.doer.Do(ctx, "POST", "/notifications/{notification_id}/replay", req, nil); err != nil {
		return err
	}

	return nil
}
