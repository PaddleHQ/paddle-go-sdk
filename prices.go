// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"

	paddleerr "github.com/PaddleHQ/paddle-go-sdk/pkg/paddleerr"
)

// ErrPriceTrialPeriodMissingFields represents a `price_trial_period_missing_fields` error.
// See https://developer.paddle.com/errors/prices/price_trial_period_missing_fields for more information.
var ErrPriceTrialPeriodMissingFields = &paddleerr.Error{
	Code: "price_trial_period_missing_fields",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrPriceTrialPeriodRequiresBillingCycle represents a `price_trial_period_requires_billing_cycle` error.
// See https://developer.paddle.com/errors/prices/price_trial_period_requires_billing_cycle for more information.
var ErrPriceTrialPeriodRequiresBillingCycle = &paddleerr.Error{
	Code: "price_trial_period_requires_billing_cycle",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrPriceBillingCycleFrequencyBelow1 represents a `price_billing_cycle_frequency_below_1` error.
// See https://developer.paddle.com/errors/prices/price_billing_cycle_frequency_below_1 for more information.
var ErrPriceBillingCycleFrequencyBelow1 = &paddleerr.Error{
	Code: "price_billing_cycle_frequency_below_1",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrPriceTrialPeriodFrequencyBelow1 represents a `price_trial_period_frequency_below_1` error.
// See https://developer.paddle.com/errors/prices/price_trial_period_frequency_below_1 for more information.
var ErrPriceTrialPeriodFrequencyBelow1 = &paddleerr.Error{
	Code: "price_trial_period_frequency_below_1",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrPriceDuplicateCurrencyOverrideForCountry represents a `price_duplicate_currency_override_for_country` error.
// See https://developer.paddle.com/errors/prices/price_duplicate_currency_override_for_country for more information.
var ErrPriceDuplicateCurrencyOverrideForCountry = &paddleerr.Error{
	Code: "price_duplicate_currency_override_for_country",
	Type: paddleerr.ErrorTypeRequestError,
}

// PriceIncludes: Represents a price entity with included entities.
type PriceIncludes struct {
	// ID: Unique Paddle ID for this price, prefixed with `pri_`.
	ID string `json:"id,omitempty"`
	// ProductID: Paddle ID for the product that this price is for, prefixed with `pro_`.
	ProductID string `json:"product_id,omitempty"`
	// Description: Internal description for this price, not shown to customers. Typically notes for your team.
	Description string `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type string `json:"type,omitempty"`
	// Name: Name of this price, shown to customers at checkout and on invoices. Typically describes how often the related product bills.
	Name *string `json:"name,omitempty"`
	// BillingCycle: How often this price should be charged. `null` if price is non-recurring (one-time).
	BillingCycle *Duration `json:"billing_cycle,omitempty"`
	// TrialPeriod: Trial period for the product related to this price. The billing cycle begins once the trial period is over. `null` for no trial period. Requires `billing_cycle`.
	TrialPeriod *Duration `json:"trial_period,omitempty"`
	// TaxMode: How tax is calculated for this price.
	TaxMode string `json:"tax_mode,omitempty"`
	// UnitPrice: Base price. This price applies to all customers, except for customers located in countries where you have `unit_price_overrides`.
	UnitPrice Money `json:"unit_price,omitempty"`
	// UnitPriceOverrides: List of unit price overrides. Use to override the base price with a custom price and currency for a country or group of countries.
	UnitPriceOverrides []UnitPriceOverride `json:"unit_price_overrides,omitempty"`
	// Quantity: Limits on how many times the related product can be purchased at this price. Useful for discount campaigns.
	Quantity PriceQuantity `json:"quantity,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status string `json:"status,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Product: Related product for this price. Returned when the `include` parameter is used with the `product` value.
	Product Product `json:"product,omitempty"`
}

// PricesClient is a client for the Prices resource.
type PricesClient struct {
	doer Doer
}

// ListPricesRequest is given as an input to ListPrices.
type ListPricesRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// ID is a query parameter.
	// Return only the IDs specified. Use a comma-separated list to get multiple entities.
	ID []string `in:"query=id;omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `billing_cycle.frequency`, `billing_cycle.interval`, `id`, `product_id`, `quantity.maximum`, `quantity.minimum`, `status`, `tax_mode`, `unit_price.amount`, and `unit_price.currency_code`.
	*/
	OrderBy *string `in:"query=order_by;omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `50`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page;omitempty" json:"-"`
	// ProductID is a query parameter.
	// Return entities related to the specified product. Use a comma-separated list to specify multiple product IDs.
	ProductID []string `in:"query=product_id;omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status;omitempty" json:"-"`
	// Recurring is a query parameter.
	// Determine whether returned entities are for recurring prices (`true`) or one-time prices (`false`).
	Recurring *bool `in:"query=recurring;omitempty" json:"-"`
	// Type is a query parameter.
	// Return items that match the specified type.
	Type *string `in:"query=type;omitempty" json:"-"`

	// IncludeProduct allows requesting the product sub-resource as part of this request.
	// If set to true, will be included on the response.
	IncludeProduct bool `in:"paddle_include=product" json:"-"`
}

// ListPrices performs the GET operation on a Prices resource.
func (c *PricesClient) ListPrices(ctx context.Context, req *ListPricesRequest) (res *Collection[*PriceIncludes], err error) {
	if err := c.doer.Do(ctx, "GET", "/prices", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreatePriceRequest is given as an input to CreatePrice.
type CreatePriceRequest struct {
	// Description: Internal description for this price, not shown to customers. Typically notes for your team.
	Description string `json:"description,omitempty"`
	// ProductID: Paddle ID for the product that this price is for, prefixed with `pro_`.
	ProductID string `json:"product_id,omitempty"`
	// UnitPrice: Base price. This price applies to all customers, except for customers located in countries where you have `unit_price_overrides`.
	UnitPrice Money `json:"unit_price,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type *string `json:"type,omitempty"`
	// Name: Name of this price, shown to customers at checkout and on invoices. Typically describes how often the related product bills.
	Name *string `json:"name,omitempty"`
	// BillingCycle: How often this price should be charged. `null` if price is non-recurring (one-time).
	BillingCycle *Duration `json:"billing_cycle,omitempty"`
	// TrialPeriod: Trial period for the product related to this price. The billing cycle begins once the trial period is over. `null` for no trial period. Requires `billing_cycle`.
	TrialPeriod *Duration `json:"trial_period,omitempty"`
	// TaxMode: How tax is calculated for this price.
	TaxMode *string `json:"tax_mode,omitempty"`
	// UnitPriceOverrides: List of unit price overrides. Use to override the base price with a custom price and currency for a country or group of countries.
	UnitPriceOverrides []UnitPriceOverride `json:"unit_price_overrides,omitempty"`
	// Quantity: Limits on how many times the related product can be purchased at this price. Useful for discount campaigns. If omitted, defaults to 1-100.
	Quantity *PriceQuantity `json:"quantity,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
}

// CreatePrice performs the POST operation on a Prices resource.
func (c *PricesClient) CreatePrice(ctx context.Context, req *CreatePriceRequest) (res *Price, err error) {
	if err := c.doer.Do(ctx, "POST", "/prices", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPriceRequest is given as an input to GetPrice.
type GetPriceRequest struct {
	// URL path parameters.
	PriceID string `in:"path=price_id" json:"-"`

	// IncludeProduct allows requesting the product sub-resource as part of this request.
	// If set to true, will be included on the response.
	IncludeProduct bool `in:"paddle_include=product" json:"-"`
}

// GetPrice performs the GET operation on a Prices resource.
func (c *PricesClient) GetPrice(ctx context.Context, req *GetPriceRequest) (res *PriceIncludes, err error) {
	if err := c.doer.Do(ctx, "GET", "/prices/{price_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdatePriceRequest is given as an input to UpdatePrice.
type UpdatePriceRequest struct {
	// URL path parameters.
	PriceID string `in:"path=price_id" json:"-"`

	// Description: Internal description for this price, not shown to customers. Typically notes for your team.
	Description *PatchField[string] `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type *PatchField[string] `json:"type,omitempty"`
	// Name: Name of this price, shown to customers at checkout and on invoices. Typically describes how often the related product bills.
	Name *PatchField[*string] `json:"name,omitempty"`
	// BillingCycle: How often this price should be charged. `null` if price is non-recurring (one-time).
	BillingCycle *PatchField[*Duration] `json:"billing_cycle,omitempty"`
	// TrialPeriod: Trial period for the product related to this price. The billing cycle begins once the trial period is over. `null` for no trial period. Requires `billing_cycle`.
	TrialPeriod *PatchField[*Duration] `json:"trial_period,omitempty"`
	// TaxMode: How tax is calculated for this price.
	TaxMode *PatchField[string] `json:"tax_mode,omitempty"`
	// UnitPrice: Base price. This price applies to all customers, except for customers located in countries where you have `unit_price_overrides`.
	UnitPrice *PatchField[Money] `json:"unit_price,omitempty"`
	// UnitPriceOverrides: List of unit price overrides. Use to override the base price with a custom price and currency for a country or group of countries.
	UnitPriceOverrides *PatchField[[]UnitPriceOverride] `json:"unit_price_overrides,omitempty"`
	// Quantity: Limits on how many times the related product can be purchased at this price. Useful for discount campaigns.
	Quantity *PatchField[PriceQuantity] `json:"quantity,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status *PatchField[string] `json:"status,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData *PatchField[CustomData] `json:"custom_data,omitempty"`
}

// UpdatePrice performs the PATCH operation on a Prices resource.
func (c *PricesClient) UpdatePrice(ctx context.Context, req *UpdatePriceRequest) (res *Price, err error) {
	if err := c.doer.Do(ctx, "PATCH", "/prices/{price_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
