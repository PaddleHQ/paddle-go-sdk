// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// DiscountCreated represents the discount.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type DiscountCreated struct {
	GenericNotificationEvent
	Data DiscountNotification `json:"data"`
}

// DiscountImported represents the discount.imported event.
// See https://developer.paddle.com/webhooks/overview for more information.
type DiscountImported struct {
	GenericNotificationEvent
	Data DiscountNotification `json:"data"`
}

// DiscountUpdated represents the discount.updated event.
// See https://developer.paddle.com/webhooks/overview for more information.
type DiscountUpdated struct {
	GenericNotificationEvent
	Data DiscountNotification `json:"data"`
}

// DiscountStatus: Whether this entity can be used in Paddle. `expired` and `used` are set automatically by Paddle..
type DiscountStatus string

const (
	DiscountStatusActive   DiscountStatus = "active"
	DiscountStatusArchived DiscountStatus = "archived"
	DiscountStatusExpired  DiscountStatus = "expired"
	DiscountStatusUsed     DiscountStatus = "used"
)

// Type: Type of discount. Determines how this discount impacts the checkout or transaction total..
type Type string

const (
	TypeFlat        Type = "flat"
	TypeFlatPerSeat Type = "flat_per_seat"
	TypePercentage  Type = "percentage"
)

// DiscountNotification: New or changed entity.
type DiscountNotification struct {
	// ID: Unique Paddle ID for this discount, prefixed with `dsc_`.
	ID string `json:"id,omitempty"`
	// Status: Whether this entity can be used in Paddle. `expired` and `used` are set automatically by Paddle.
	Status DiscountStatus `json:"status,omitempty"`
	// Description: Short description for this discount for your reference. Not shown to customers.
	Description string `json:"description,omitempty"`
	// EnabledForCheckout: Whether this discount can be redeemed by customers at checkout (`true`) or not (`false`).
	EnabledForCheckout bool `json:"enabled_for_checkout,omitempty"`
	// Code: Unique code that customers can use to redeem this discount at checkout.
	Code *string `json:"code,omitempty"`
	// Type: Type of discount. Determines how this discount impacts the checkout or transaction total.
	Type Type `json:"type,omitempty"`
	// Amount: Amount to discount by. For `percentage` discounts, must be an amount between `0.01` and `100`. For `flat` and `flat_per_seat` discounts, amount in the lowest denomination for a currency.
	Amount string `json:"amount,omitempty"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code. Required where discount type is `flat` or `flat_per_seat`.
	CurrencyCode *CurrencyCode `json:"currency_code,omitempty"`
	// Recur: Whether this discount applies for multiple subscription billing periods (`true`) or not (`false`).
	Recur bool `json:"recur,omitempty"`
	/*
	   MaximumRecurringIntervals: Number of subscription billing periods that this discount recurs for. Requires `recur`. `null` if this discount recurs forever.

	   Subscription renewals, midcycle changes, and one-time charges billed to a subscription aren't considered a redemption. `times_used` is not incremented in these cases.
	*/
	MaximumRecurringIntervals *int `json:"maximum_recurring_intervals,omitempty"`
	/*
	   UsageLimit: Maximum number of times this discount can be redeemed. This is an overall limit for this discount, rather than a per-customer limit. `null` if this discount can be redeemed an unlimited amount of times.

	   Paddle counts a usage as a redemption on a checkout, transaction, or the initial application against a subscription. Transactions created for subscription renewals, midcycle changes, and one-time charges aren't considered a redemption.
	*/
	UsageLimit *int `json:"usage_limit,omitempty"`
	// RestrictTo: Product or price IDs that this discount is for. When including a product ID, all prices for that product can be discounted. `null` if this discount applies to all products and prices.
	RestrictTo []string `json:"restrict_to,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
	/*
	   ExpiresAt: RFC 3339 datetime string of when this discount expires. Discount can no longer be redeemed after this date has elapsed. `null` if this discount can be redeemed forever.

	   Expired discounts can't be redeemed against transactions or checkouts, but can be applied when updating subscriptions.
	*/
	ExpiresAt *string `json:"expires_at,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
}
