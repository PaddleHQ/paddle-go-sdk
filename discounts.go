// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"

	paddleerr "github.com/PaddleHQ/paddle-go-sdk/pkg/paddleerr"
)

// ErrDiscountExpired represents a `discount_expired` error.
// See https://developer.paddle.com/errors/discounts/discount_expired for more information.
var ErrDiscountExpired = &paddleerr.Error{
	Code: "discount_expired",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrDiscountUsageLimitExceeded represents a `discount_usage_limit_exceeded` error.
// See https://developer.paddle.com/errors/discounts/discount_usage_limit_exceeded for more information.
var ErrDiscountUsageLimitExceeded = &paddleerr.Error{
	Code: "discount_usage_limit_exceeded",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrDiscountCodeConflict represents a `discount_code_conflict` error.
// See https://developer.paddle.com/errors/discounts/discount_code_conflict for more information.
var ErrDiscountCodeConflict = &paddleerr.Error{
	Code: "discount_code_conflict",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrDiscountRestrictedProductNotActive represents a `discount_restricted_product_not_active` error.
// See https://developer.paddle.com/errors/discounts/discount_restricted_product_not_active for more information.
var ErrDiscountRestrictedProductNotActive = &paddleerr.Error{
	Code: "discount_restricted_product_not_active",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrDiscountRestrictedProductPriceNotActive represents a `discount_restricted_product_price_not_active` error.
// See https://developer.paddle.com/errors/discounts/discount_restricted_product_price_not_active for more information.
var ErrDiscountRestrictedProductPriceNotActive = &paddleerr.Error{
	Code: "discount_restricted_product_price_not_active",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrDiscountUsageLimitLessThanTimesUsed represents a `discount_usage_limit_less_than_times_used` error.
// See https://developer.paddle.com/errors/discounts/discount_usage_limit_less_than_times_used for more information.
var ErrDiscountUsageLimitLessThanTimesUsed = &paddleerr.Error{
	Code: "discount_usage_limit_less_than_times_used",
	Type: paddleerr.ErrorTypeRequestError,
}

// DiscountsClient is a client for the Discounts resource.
type DiscountsClient struct {
	doer Doer
}

// ListDiscountsRequest is given as an input to ListDiscounts.
type ListDiscountsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// Code is a query parameter.
	// Return entities that match the discount code. Use a comma-separated list to specify multiple discount codes.
	Code []string `in:"query=code;omitempty" json:"-"`
	// ID is a query parameter.
	// Return only the IDs specified. Use a comma-separated list to get multiple entities.
	ID []string `in:"query=id;omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `created_at` and `id`.
	*/
	OrderBy *string `in:"query=order_by;omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `50`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page;omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status;omitempty" json:"-"`
}

// ListDiscounts performs the GET operation on a Discounts resource.
func (c *DiscountsClient) ListDiscounts(ctx context.Context, req *ListDiscountsRequest) (res *Collection[*Discount], err error) {
	if err := c.doer.Do(ctx, "GET", "/discounts", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateDiscountRequest is given as an input to CreateDiscount.
type CreateDiscountRequest struct {
	// Amount: Amount to discount by. For `percentage` discounts, must be an amount between `0.01` and `100`. For `flat` and `flat_per_seat` discounts, amount in the lowest denomination for a currency.
	Amount string `json:"amount,omitempty"`
	// Description: Short description for this discount for your reference. Not shown to customers.
	Description string `json:"description,omitempty"`
	// Type: Type of discount.
	Type DiscountType `json:"type,omitempty"`
	// EnabledForCheckout: Whether this discount can be applied by a customer at checkout.
	EnabledForCheckout *bool `json:"enabled_for_checkout,omitempty"`
	// Code: Unique code that customers can use to apply this discount at checkout. Use letters and numbers only, up to 16 characters. Paddle generates a random 10-character code if a code is not provided and `enabled_for_checkout` is `true`.
	Code *string `json:"code,omitempty"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code. Required where discount type is `flat` or `flat_per_seat`.
	CurrencyCode *CurrencyCode `json:"currency_code,omitempty"`
	// Recur: Whether this discount applies for multiple billing periods.
	Recur *bool `json:"recur,omitempty"`
	// MaximumRecurringIntervals: Amount of subscription billing periods that this discount recurs for. Requires `recur`. `null` if this discount recurs forever.
	MaximumRecurringIntervals *int `json:"maximum_recurring_intervals,omitempty"`
	// UsageLimit: Maximum amount of times this discount can be used. This is an overall limit, rather than a per-customer limit. `null` if this discount can be used an unlimited amount of times.
	UsageLimit *int `json:"usage_limit,omitempty"`
	// RestrictTo: Product or price IDs that this discount is for. When including a product ID, all prices for that product can be discounted. `null` if this discount applies to all products and prices.
	RestrictTo []string `json:"restrict_to,omitempty"`
	// ExpiresAt: RFC 3339 datetime string of when this discount expires. Discount can no longer be applied after this date has elapsed. `null` if this discount can be applied forever.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
}

// CreateDiscount performs the POST operation on a Discounts resource.
func (c *DiscountsClient) CreateDiscount(ctx context.Context, req *CreateDiscountRequest) (res *Discount, err error) {
	if err := c.doer.Do(ctx, "POST", "/discounts", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetDiscountRequest is given as an input to GetDiscount.
type GetDiscountRequest struct {
	// URL path parameters.
	DiscountID string `in:"path=discount_id" json:"-"`
}

// GetDiscount performs the GET operation on a Discounts resource.
func (c *DiscountsClient) GetDiscount(ctx context.Context, req *GetDiscountRequest) (res *Discount, err error) {
	if err := c.doer.Do(ctx, "GET", "/discounts/{discount_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateDiscountRequest is given as an input to UpdateDiscount.
type UpdateDiscountRequest struct {
	// URL path parameters.
	DiscountID string `in:"path=discount_id" json:"-"`

	// Status: Whether this entity can be used in Paddle. `expired` and `used` are set automatically by Paddle.
	Status *PatchField[DiscountStatus] `json:"status,omitempty"`
	// Description: Short description for this discount for your reference. Not shown to customers.
	Description *PatchField[string] `json:"description,omitempty"`
	// EnabledForCheckout: Whether this discount can be applied by a customer at checkout.
	EnabledForCheckout *PatchField[bool] `json:"enabled_for_checkout,omitempty"`
	// Code: Unique code that customers can use to apply this discount at checkout. Use letters and numbers only, up to 16 characters. Paddle generates a random 10-character code if a code is not provided and `enabled_for_checkout` is `true`.
	Code *PatchField[*string] `json:"code,omitempty"`
	// Type: Type of discount.
	Type *PatchField[DiscountType] `json:"type,omitempty"`
	// Amount: Amount to discount by. For `percentage` discounts, must be an amount between `0.01` and `100`. For `flat` and `flat_per_seat` discounts, amount in the lowest denomination for a currency.
	Amount *PatchField[string] `json:"amount,omitempty"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code. Required where discount type is `flat` or `flat_per_seat`.
	CurrencyCode *PatchField[*CurrencyCode] `json:"currency_code,omitempty"`
	// Recur: Whether this discount applies for multiple billing periods.
	Recur *PatchField[bool] `json:"recur,omitempty"`
	// MaximumRecurringIntervals: Amount of subscription billing periods that this discount recurs for. Requires `recur`. `null` if this discount recurs forever.
	MaximumRecurringIntervals *PatchField[*int] `json:"maximum_recurring_intervals,omitempty"`
	// UsageLimit: Maximum amount of times this discount can be used. This is an overall limit, rather than a per-customer limit. `null` if this discount can be used an unlimited amount of times.
	UsageLimit *PatchField[*int] `json:"usage_limit,omitempty"`
	// RestrictTo: Product or price IDs that this discount is for. When including a product ID, all prices for that product can be discounted. `null` if this discount applies to all products and prices.
	RestrictTo *PatchField[[]string] `json:"restrict_to,omitempty"`
	// ExpiresAt: RFC 3339 datetime string of when this discount expires. Discount can no longer be applied after this date has elapsed. `null` if this discount can be applied forever.
	ExpiresAt *PatchField[*string] `json:"expires_at,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData *PatchField[CustomData] `json:"custom_data,omitempty"`
}

// UpdateDiscount performs the PATCH operation on a Discounts resource.
func (c *DiscountsClient) UpdateDiscount(ctx context.Context, req *UpdateDiscountRequest) (res *Discount, err error) {
	if err := c.doer.Do(ctx, "PATCH", "/discounts/{discount_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
