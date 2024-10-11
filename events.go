// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"
	"encoding/json"

	paddlenotification "github.com/PaddleHQ/paddle-go-sdk/pkg/paddlenotification"
)

// Event: Represents an event entity.
type Event interface{}

// GenericEvent is an generic implementation for Event
type GenericEvent struct {
	Event
	// EventID: Unique Paddle ID for this event, prefixed with `evt_`.
	EventID string `json:"event_id,omitempty"`
	// EventType: Type of event sent by Paddle, in the format `entity.event_type`.
	EventType EventTypeName `json:"event_type,omitempty"`
	// OccurredAt: RFC 3339 datetime string of when this event occurred.
	OccurredAt string `json:"occurred_at,omitempty"`
	// Data: New or changed entity.
	Data paddlenotification.NotificationPayload `json:"data,omitempty"`
}

// AddressCreatedEvent represents an Event implementation for address.created event.
type AddressCreatedEvent struct {
	GenericEvent
	Data paddlenotification.AddressNotification `json:"data"`
}

// AddressImportedEvent represents an Event implementation for address.imported event.
type AddressImportedEvent struct {
	GenericEvent
	Data paddlenotification.AddressNotification `json:"data"`
}

// AddressUpdatedEvent represents an Event implementation for address.updated event.
type AddressUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.AddressNotification `json:"data"`
}

// AdjustmentCreatedEvent represents an Event implementation for adjustment.created event.
type AdjustmentCreatedEvent struct {
	GenericEvent
	Data paddlenotification.AdjustmentNotification `json:"data"`
}

// AdjustmentUpdatedEvent represents an Event implementation for adjustment.updated event.
type AdjustmentUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.AdjustmentNotification `json:"data"`
}

// BusinessCreatedEvent represents an Event implementation for business.created event.
type BusinessCreatedEvent struct {
	GenericEvent
	Data paddlenotification.BusinessNotification `json:"data"`
}

// BusinessImportedEvent represents an Event implementation for business.imported event.
type BusinessImportedEvent struct {
	GenericEvent
	Data paddlenotification.BusinessNotification `json:"data"`
}

// BusinessUpdatedEvent represents an Event implementation for business.updated event.
type BusinessUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.BusinessNotification `json:"data"`
}

// CustomerCreatedEvent represents an Event implementation for customer.created event.
type CustomerCreatedEvent struct {
	GenericEvent
	Data paddlenotification.CustomerNotification `json:"data"`
}

// CustomerImportedEvent represents an Event implementation for customer.imported event.
type CustomerImportedEvent struct {
	GenericEvent
	Data paddlenotification.CustomerNotification `json:"data"`
}

// CustomerUpdatedEvent represents an Event implementation for customer.updated event.
type CustomerUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.CustomerNotification `json:"data"`
}

// DiscountCreatedEvent represents an Event implementation for discount.created event.
type DiscountCreatedEvent struct {
	GenericEvent
	Data paddlenotification.DiscountNotification `json:"data"`
}

// DiscountImportedEvent represents an Event implementation for discount.imported event.
type DiscountImportedEvent struct {
	GenericEvent
	Data paddlenotification.DiscountNotification `json:"data"`
}

// DiscountUpdatedEvent represents an Event implementation for discount.updated event.
type DiscountUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.DiscountNotification `json:"data"`
}

// PayoutCreatedEvent represents an Event implementation for payout.created event.
type PayoutCreatedEvent struct {
	GenericEvent
	Data paddlenotification.PayoutNotification `json:"data"`
}

// PayoutPaidEvent represents an Event implementation for payout.paid event.
type PayoutPaidEvent struct {
	GenericEvent
	Data paddlenotification.PayoutNotification `json:"data"`
}

// PriceCreatedEvent represents an Event implementation for price.created event.
type PriceCreatedEvent struct {
	GenericEvent
	Data paddlenotification.PriceNotification `json:"data"`
}

// PriceImportedEvent represents an Event implementation for price.imported event.
type PriceImportedEvent struct {
	GenericEvent
	Data paddlenotification.PriceNotification `json:"data"`
}

// PriceUpdatedEvent represents an Event implementation for price.updated event.
type PriceUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.PriceNotification `json:"data"`
}

// ProductCreatedEvent represents an Event implementation for product.created event.
type ProductCreatedEvent struct {
	GenericEvent
	Data paddlenotification.ProductNotification `json:"data"`
}

// ProductImportedEvent represents an Event implementation for product.imported event.
type ProductImportedEvent struct {
	GenericEvent
	Data paddlenotification.ProductNotification `json:"data"`
}

// ProductUpdatedEvent represents an Event implementation for product.updated event.
type ProductUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.ProductNotification `json:"data"`
}

// ReportCreatedEvent represents an Event implementation for report.created event.
type ReportCreatedEvent struct {
	GenericEvent
	Data paddlenotification.ReportNotification `json:"data"`
}

// ReportUpdatedEvent represents an Event implementation for report.updated event.
type ReportUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.ReportNotification `json:"data"`
}

// SubscriptionActivatedEvent represents an Event implementation for subscription.activated event.
type SubscriptionActivatedEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionCanceledEvent represents an Event implementation for subscription.canceled event.
type SubscriptionCanceledEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionCreatedEvent represents an Event implementation for subscription.created event.
type SubscriptionCreatedEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionCreatedNotification `json:"data"`
}

// SubscriptionPastDueEvent represents an Event implementation for subscription.past_due event.
type SubscriptionPastDueEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionImportedEvent represents an Event implementation for subscription.imported event.
type SubscriptionImportedEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionPausedEvent represents an Event implementation for subscription.paused event.
type SubscriptionPausedEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionResumedEvent represents an Event implementation for subscription.resumed event.
type SubscriptionResumedEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionTrialingEvent represents an Event implementation for subscription.trialing event.
type SubscriptionTrialingEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// SubscriptionUpdatedEvent represents an Event implementation for subscription.updated event.
type SubscriptionUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.SubscriptionNotification `json:"data"`
}

// TransactionBilledEvent represents an Event implementation for transaction.billed event.
type TransactionBilledEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionCanceledEvent represents an Event implementation for transaction.canceled event.
type TransactionCanceledEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionCompletedEvent represents an Event implementation for transaction.completed event.
type TransactionCompletedEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionCreatedEvent represents an Event implementation for transaction.created event.
type TransactionCreatedEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionPaidEvent represents an Event implementation for transaction.paid event.
type TransactionPaidEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionPastDueEvent represents an Event implementation for transaction.past_due event.
type TransactionPastDueEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionPaymentFailedEvent represents an Event implementation for transaction.payment_failed event.
type TransactionPaymentFailedEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionReadyEvent represents an Event implementation for transaction.ready event.
type TransactionReadyEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// TransactionUpdatedEvent represents an Event implementation for transaction.updated event.
type TransactionUpdatedEvent struct {
	GenericEvent
	Data paddlenotification.TransactionNotification `json:"data"`
}

// unmarshalEvent unmarshals JSON data to the correct Event implementation
func unmarshalEvent(data []byte) (Event, error) {
	e := &GenericEvent{}
	if err := json.Unmarshal(data, e); err != nil {
		return nil, err
	}

	var t Event
	switch e.EventType {
	case "address.created":
		t = &AddressCreatedEvent{}
	case "address.imported":
		t = &AddressImportedEvent{}
	case "address.updated":
		t = &AddressUpdatedEvent{}
	case "adjustment.created":
		t = &AdjustmentCreatedEvent{}
	case "adjustment.updated":
		t = &AdjustmentUpdatedEvent{}
	case "business.created":
		t = &BusinessCreatedEvent{}
	case "business.imported":
		t = &BusinessImportedEvent{}
	case "business.updated":
		t = &BusinessUpdatedEvent{}
	case "customer.created":
		t = &CustomerCreatedEvent{}
	case "customer.imported":
		t = &CustomerImportedEvent{}
	case "customer.updated":
		t = &CustomerUpdatedEvent{}
	case "discount.created":
		t = &DiscountCreatedEvent{}
	case "discount.imported":
		t = &DiscountImportedEvent{}
	case "discount.updated":
		t = &DiscountUpdatedEvent{}
	case "payout.created":
		t = &PayoutCreatedEvent{}
	case "payout.paid":
		t = &PayoutPaidEvent{}
	case "price.created":
		t = &PriceCreatedEvent{}
	case "price.imported":
		t = &PriceImportedEvent{}
	case "price.updated":
		t = &PriceUpdatedEvent{}
	case "product.created":
		t = &ProductCreatedEvent{}
	case "product.imported":
		t = &ProductImportedEvent{}
	case "product.updated":
		t = &ProductUpdatedEvent{}
	case "report.created":
		t = &ReportCreatedEvent{}
	case "report.updated":
		t = &ReportUpdatedEvent{}
	case "subscription.activated":
		t = &SubscriptionActivatedEvent{}
	case "subscription.canceled":
		t = &SubscriptionCanceledEvent{}
	case "subscription.created":
		t = &SubscriptionCreatedEvent{}
	case "subscription.past_due":
		t = &SubscriptionPastDueEvent{}
	case "subscription.imported":
		t = &SubscriptionImportedEvent{}
	case "subscription.paused":
		t = &SubscriptionPausedEvent{}
	case "subscription.resumed":
		t = &SubscriptionResumedEvent{}
	case "subscription.trialing":
		t = &SubscriptionTrialingEvent{}
	case "subscription.updated":
		t = &SubscriptionUpdatedEvent{}
	case "transaction.billed":
		t = &TransactionBilledEvent{}
	case "transaction.canceled":
		t = &TransactionCanceledEvent{}
	case "transaction.completed":
		t = &TransactionCompletedEvent{}
	case "transaction.created":
		t = &TransactionCreatedEvent{}
	case "transaction.paid":
		t = &TransactionPaidEvent{}
	case "transaction.past_due":
		t = &TransactionPastDueEvent{}
	case "transaction.payment_failed":
		t = &TransactionPaymentFailedEvent{}
	case "transaction.ready":
		t = &TransactionReadyEvent{}
	case "transaction.updated":
		t = &TransactionUpdatedEvent{}
	default:
		return e, nil
	}

	err := json.Unmarshal(data, t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

// EventsClient is a client for the Events resource.
type EventsClient struct {
	doer Doer
}

// ListEventsRequest is given as an input to ListEvents.
type ListEventsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `id` (for `event_id`).
	*/
	OrderBy *string `in:"query=order_by;omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `50`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page;omitempty" json:"-"`
}

// ListEvents performs the GET operation on a Events resource.
func (c *EventsClient) ListEvents(ctx context.Context, req *ListEventsRequest) (res *Collection[Event], err error) {
	if err := c.doer.Do(ctx, "GET", "/events", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
