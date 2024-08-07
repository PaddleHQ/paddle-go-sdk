// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// CustomerCreated represents the customer.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type CustomerCreated struct {
	GenericNotificationEvent
	Data CustomerNotification `json:"data"`
}

// CustomerImported represents the customer.imported event.
// See https://developer.paddle.com/webhooks/overview for more information.
type CustomerImported struct {
	GenericNotificationEvent
	Data CustomerNotification `json:"data"`
}

// CustomerUpdated represents the customer.updated event.
// See https://developer.paddle.com/webhooks/overview for more information.
type CustomerUpdated struct {
	GenericNotificationEvent
	Data CustomerNotification `json:"data"`
}

// CustomerNotification: New or changed entity.
type CustomerNotification struct {
	// ID: Unique Paddle ID for this customer entity, prefixed with `ctm_`.
	ID string `json:"id,omitempty"`
	// Name: Full name of this customer. Required when creating transactions where `collection_mode` is `manual` (invoices).
	Name *string `json:"name,omitempty"`
	// Email: Email address for this customer.
	Email string `json:"email,omitempty"`
	/*
	   MarketingConsent: Whether this customer opted into marketing from you. `false` unless customers check the marketing consent box
	   when using Paddle Checkout. Set automatically by Paddle.
	*/
	MarketingConsent bool `json:"marketing_consent,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status Status `json:"status,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// Locale: Valid IETF BCP 47 short form locale tag.
	Locale string `json:"locale,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
}
