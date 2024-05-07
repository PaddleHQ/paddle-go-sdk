// Code generated by the Paddle SDK Generator; DO NOT EDIT.
package paddle

import "context"

// Event: Represents an event entity.
type Event struct {
	// EventID: Unique Paddle ID for this event, prefixed with `evt_`.
	EventID string `json:"event_id,omitempty"`
	// EventType: Type of event sent by Paddle, in the format `entity.event_type`.
	EventType string `json:"event_type,omitempty"`
	// OccurredAt: RFC 3339 datetime string of when this event occurred.
	OccurredAt string `json:"occurred_at,omitempty"`
	// Data: New or changed entity.
	Data any `json:"data,omitempty"`
}

// EventsClient is a client for the Events resource.
type EventsClient struct {
	doer Doer
}

// ListEventsRequest is given as an input to ListEvents.
type ListEventsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after,omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `id` (for `event_id`).
	*/
	OrderBy *string `in:"query=order_by,omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `50`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page,omitempty" json:"-"`
}

// ListEvents performs the GET operation on a Events resource.
func (c *EventsClient) ListEvents(ctx context.Context, req *ListEventsRequest) (res *Collection[*Event], err error) {
	if err := c.doer.Do(ctx, "GET", "/events", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
