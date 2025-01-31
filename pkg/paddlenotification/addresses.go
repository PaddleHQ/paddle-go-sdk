// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// AddressCreated represents the address.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type AddressCreated struct {
	GenericNotificationEvent
	Data AddressNotification `json:"data"`
}

// AddressImported represents the address.imported event.
// See https://developer.paddle.com/webhooks/overview for more information.
type AddressImported struct {
	GenericNotificationEvent
	Data AddressNotification `json:"data"`
}

// AddressUpdated represents the address.updated event.
// See https://developer.paddle.com/webhooks/overview for more information.
type AddressUpdated struct {
	GenericNotificationEvent
	Data AddressNotification `json:"data"`
}

// AddressNotification: New or changed entity.
type AddressNotification struct {
	NotificationPayload `json:"-"`

	// ID: Unique Paddle ID for this address entity, prefixed with `add_`.
	ID string `json:"id"`
	// CustomerID: Paddle ID for the customer related to this address, prefixed with `cus_`.
	CustomerID string `json:"customer_id"`
	// Description: Memorable description for this address.
	Description *string `json:"description"`
	// FirstLine: First line of this address.
	FirstLine *string `json:"first_line"`
	// SecondLine: Second line of this address.
	SecondLine *string `json:"second_line"`
	// City: City of this address.
	City *string `json:"city"`
	// PostalCode: ZIP or postal code of this address. Required for some countries.
	PostalCode *string `json:"postal_code"`
	// Region: State, county, or region of this address.
	Region *string `json:"region"`
	// CountryCode: Supported two-letter ISO 3166-1 alpha-2 country code for this address.
	CountryCode CountryCode `json:"country_code"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data"`
	// Status: Whether this entity can be used in Paddle.
	Status Status `json:"status"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta"`
}
