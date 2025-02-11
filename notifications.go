// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"
	"encoding/json"

	paddleerr "github.com/PaddleHQ/paddle-go-sdk/v3/pkg/paddleerr"
	paddlenotification "github.com/PaddleHQ/paddle-go-sdk/v3/pkg/paddlenotification"
)

// ErrNotificationMaximumActiveSettingsReached represents a `notification_maximum_active_settings_reached` error.
// See https://developer.paddle.com/errors/notifications/notification_maximum_active_settings_reached for more information.
var ErrNotificationMaximumActiveSettingsReached = &paddleerr.Error{
	Code: "notification_maximum_active_settings_reached",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrNotificationCannotReplay represents a `notification_cannot_replay` error.
// See https://developer.paddle.com/errors/notifications/notification_cannot_replay for more information.
var ErrNotificationCannotReplay = &paddleerr.Error{
	Code: "notification_cannot_replay",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrURLNotificationSettingIncorrect represents a `url_notification_setting_incorrect` error.
// See https://developer.paddle.com/errors/notifications/url_notification_setting_incorrect for more information.
var ErrURLNotificationSettingIncorrect = &paddleerr.Error{
	Code: "url_notification_setting_incorrect",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrEmailNotificationSettingIncorrect represents a `email_notification_setting_incorrect` error.
// See https://developer.paddle.com/errors/notifications/email_notification_setting_incorrect for more information.
var ErrEmailNotificationSettingIncorrect = &paddleerr.Error{
	Code: "email_notification_setting_incorrect",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrNotificationReplayInvalidOriginType represents a `notification_replay_invalid_origin_type` error.
// See https://developer.paddle.com/errors/notifications/notification_replay_invalid_origin_type for more information.
var ErrNotificationReplayInvalidOriginType = &paddleerr.Error{
	Code: "notification_replay_invalid_origin_type",
	Type: paddleerr.ErrorTypeRequestError,
}

// NotificationStatus: Status of this notification..
type NotificationStatus string

const (
	NotificationStatusNotAttempted NotificationStatus = "not_attempted"
	NotificationStatusNeedsRetry   NotificationStatus = "needs_retry"
	NotificationStatusDelivered    NotificationStatus = "delivered"
	NotificationStatusFailed       NotificationStatus = "failed"
)

// NotificationOrigin: Describes how this notification was created..
type NotificationOrigin string

const (
	NotificationOriginEvent  NotificationOrigin = "event"
	NotificationOriginReplay NotificationOrigin = "replay"
)

// Notification: Represents a notification entity.
type Notification struct {
	// ID: Unique Paddle ID for this notification, prefixed with `ntf_`.
	ID string `json:"id,omitempty"`
	// Type: Type of event sent by Paddle, in the format `entity.event_type`.
	Type EventTypeName `json:"type,omitempty"`
	// Status: Status of this notification.
	Status NotificationStatus `json:"status,omitempty"`
	// Payload: Notification payload. Includes the new or changed event.
	Payload paddlenotification.NotificationEvent `json:"payload,omitempty"`
	// OccurredAt: RFC 3339 datetime string of when this notification occurred.
	OccurredAt string `json:"occurred_at,omitempty"`
	// DeliveredAt: RFC 3339 datetime string of when this notification was delivered. `null` if not yet delivered successfully.
	DeliveredAt *string `json:"delivered_at,omitempty"`
	// ReplayedAt: RFC 3339 datetime string of when this notification was replayed. `null` if not replayed.
	ReplayedAt *string `json:"replayed_at,omitempty"`
	// Origin: Describes how this notification was created.
	Origin NotificationOrigin `json:"origin,omitempty"`
	// LastAttemptAt: RFC 3339 datetime string of when this notification was last attempted.
	LastAttemptAt *string `json:"last_attempt_at,omitempty"`
	// RetryAt: RFC 3339 datetime string of when this notification is scheduled to be retried.
	RetryAt *string `json:"retry_at,omitempty"`
	// TimesAttempted: How many times delivery of this notification has been attempted. Automatically incremented by Paddle after an attempt.
	TimesAttempted int `json:"times_attempted,omitempty"`
	// NotificationSettingID: Unique Paddle ID for this notification setting, prefixed with `ntfset_`.
	NotificationSettingID string `json:"notification_setting_id,omitempty"`
}

// UnmarshalJSON implements the json.Unmarshaler interface for Notification
func (n *Notification) UnmarshalJSON(data []byte) error {
	type alias Notification
	if err := json.Unmarshal(data, (*alias)(n)); err != nil {
		return err
	}

	var t paddlenotification.NotificationEvent
	switch n.Type {
	case "address.created":
		t = &paddlenotification.AddressCreated{}
	case "address.imported":
		t = &paddlenotification.AddressImported{}
	case "address.updated":
		t = &paddlenotification.AddressUpdated{}
	case "adjustment.created":
		t = &paddlenotification.AdjustmentCreated{}
	case "adjustment.updated":
		t = &paddlenotification.AdjustmentUpdated{}
	case "business.created":
		t = &paddlenotification.BusinessCreated{}
	case "business.imported":
		t = &paddlenotification.BusinessImported{}
	case "business.updated":
		t = &paddlenotification.BusinessUpdated{}
	case "customer.created":
		t = &paddlenotification.CustomerCreated{}
	case "customer.imported":
		t = &paddlenotification.CustomerImported{}
	case "customer.updated":
		t = &paddlenotification.CustomerUpdated{}
	case "discount.created":
		t = &paddlenotification.DiscountCreated{}
	case "discount.imported":
		t = &paddlenotification.DiscountImported{}
	case "discount.updated":
		t = &paddlenotification.DiscountUpdated{}
	case "payment_method.saved":
		t = &paddlenotification.PaymentMethodSaved{}
	case "payment_method.deleted":
		t = &paddlenotification.PaymentMethodDeleted{}
	case "payout.created":
		t = &paddlenotification.PayoutCreated{}
	case "payout.paid":
		t = &paddlenotification.PayoutPaid{}
	case "price.created":
		t = &paddlenotification.PriceCreated{}
	case "price.imported":
		t = &paddlenotification.PriceImported{}
	case "price.updated":
		t = &paddlenotification.PriceUpdated{}
	case "product.created":
		t = &paddlenotification.ProductCreated{}
	case "product.imported":
		t = &paddlenotification.ProductImported{}
	case "product.updated":
		t = &paddlenotification.ProductUpdated{}
	case "report.created":
		t = &paddlenotification.ReportCreated{}
	case "report.updated":
		t = &paddlenotification.ReportUpdated{}
	case "subscription.activated":
		t = &paddlenotification.SubscriptionActivated{}
	case "subscription.canceled":
		t = &paddlenotification.SubscriptionCanceled{}
	case "subscription.created":
		t = &paddlenotification.SubscriptionCreated{}
	case "subscription.past_due":
		t = &paddlenotification.SubscriptionPastDue{}
	case "subscription.imported":
		t = &paddlenotification.SubscriptionImported{}
	case "subscription.paused":
		t = &paddlenotification.SubscriptionPaused{}
	case "subscription.resumed":
		t = &paddlenotification.SubscriptionResumed{}
	case "subscription.trialing":
		t = &paddlenotification.SubscriptionTrialing{}
	case "subscription.updated":
		t = &paddlenotification.SubscriptionUpdated{}
	case "transaction.billed":
		t = &paddlenotification.TransactionBilled{}
	case "transaction.canceled":
		t = &paddlenotification.TransactionCanceled{}
	case "transaction.completed":
		t = &paddlenotification.TransactionCompleted{}
	case "transaction.created":
		t = &paddlenotification.TransactionCreated{}
	case "transaction.paid":
		t = &paddlenotification.TransactionPaid{}
	case "transaction.past_due":
		t = &paddlenotification.TransactionPastDue{}
	case "transaction.payment_failed":
		t = &paddlenotification.TransactionPaymentFailed{}
	case "transaction.ready":
		t = &paddlenotification.TransactionReady{}
	case "transaction.updated":
		t = &paddlenotification.TransactionUpdated{}
	case "transaction.revised":
		t = &paddlenotification.TransactionRevised{}
	default:
		t = &paddlenotification.GenericNotificationEvent{}
	}

	rawT, err := json.Marshal(n.Payload)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(rawT, t); err != nil {
		return err
	}

	n.Payload = t

	return nil
}

// NotificationsClient is a client for the Notifications resource.
type NotificationsClient struct {
	doer Doer
}

// ListNotificationsRequest is given as an input to ListNotifications.
type ListNotificationsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// NotificationSettingID is a query parameter.
	// Return entities related to the specified notification destination. Use a comma-separated list to specify multiple notification destination IDs.
	NotificationSettingID []string `in:"query=notification_setting_id;omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `id`.
	*/
	OrderBy *string `in:"query=order_by;omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `50`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page;omitempty" json:"-"`
	// Search is a query parameter.
	// Return entities that match a search query. Searches `id` and `type` fields.
	Search *string `in:"query=search;omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status;omitempty" json:"-"`
	// Filter is a query parameter.
	// Return entities that contain the Paddle ID specified. Pass a transaction, customer, or subscription ID.
	Filter *string `in:"query=filter;omitempty" json:"-"`
	// To is a query parameter.
	// Return entities up to a specific time. Pass an RFC 3339 datetime string.
	To *string `in:"query=to;omitempty" json:"-"`
	// From is a query parameter.
	// Return entities from a specific time. Pass an RFC 3339 datetime string.
	From *string `in:"query=from;omitempty" json:"-"`
}

// ListNotifications performs the GET operation on a Notifications resource.
func (c *NotificationsClient) ListNotifications(ctx context.Context, req *ListNotificationsRequest) (res *Collection[*Notification], err error) {
	if err := c.doer.Do(ctx, "GET", "/notifications", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetNotificationRequest is given as an input to GetNotification.
type GetNotificationRequest struct {
	// URL path parameters.
	NotificationID string `in:"path=notification_id" json:"-"`
}

// GetNotification performs the GET operation on a Notifications resource.
func (c *NotificationsClient) GetNotification(ctx context.Context, req *GetNotificationRequest) (res *Notification, err error) {
	if err := c.doer.Do(ctx, "GET", "/notifications/{notification_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
