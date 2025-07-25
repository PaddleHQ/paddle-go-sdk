// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import "context"

// PricePreviewItem: List of items to preview price calculations for.
type PricePreviewItem struct {
	// PriceID: Paddle ID for the price to add to this transaction, prefixed with `pri_`.
	PriceID string `json:"price_id,omitempty"`
	// Quantity: Quantity of the item to preview.
	Quantity int `json:"quantity,omitempty"`
}

// PricePreviewDiscounts: Array of discounts applied to this preview line item. Empty if no discounts applied.
type PricePreviewDiscounts struct {
	// Discount: Related discount entity for this preview line item.
	Discount Discount `json:"discount,omitempty"`
	// Total: Total amount discounted as a result of this discount.
	Total string `json:"total,omitempty"`
	// FormattedTotal: Total amount discounted as a result of this discount in the format of a given currency. '
	FormattedTotal string `json:"formatted_total,omitempty"`
}

// PricePreviewLineItem: Information about line items for this preview. Includes totals calculated by Paddle. Considered the source of truth for line item totals.
type PricePreviewLineItem struct {
	// Price: Related price entity for this preview line item.
	Price Price `json:"price,omitempty"`
	// Quantity: Quantity of this preview line item.
	Quantity int `json:"quantity,omitempty"`
	// TaxRate: Rate used to calculate tax for this preview line item.
	TaxRate string `json:"tax_rate,omitempty"`
	// UnitTotals: Breakdown of the charge for one unit in the lowest denomination of a currency (e.g. cents for USD).
	UnitTotals Totals `json:"unit_totals,omitempty"`
	// FormattedUnitTotals: Breakdown of the charge for one unit in the format of a given currency.
	FormattedUnitTotals Totals `json:"formatted_unit_totals,omitempty"`
	// Totals: Breakdown of a charge in the lowest denomination of a currency (e.g. cents for USD).
	Totals Totals `json:"totals,omitempty"`
	// FormattedTotals: The financial breakdown of a charge in the format of a given currency.
	FormattedTotals Totals `json:"formatted_totals,omitempty"`
	// Product: Related product entity for this preview line item price.
	Product Product `json:"product,omitempty"`
	// Discounts: Array of discounts applied to this preview line item. Empty if no discounts applied.
	Discounts []PricePreviewDiscounts `json:"discounts,omitempty"`
}

// PricePreviewDetails: Calculated totals for a price preview, including discounts, tax, and currency conversion.
type PricePreviewDetails struct {
	// LineItems: Information about line items for this preview. Includes totals calculated by Paddle. Considered the source of truth for line item totals.
	LineItems []PricePreviewLineItem `json:"line_items,omitempty"`
}

type PricePreview struct {
	// CustomerID: Paddle ID of the customer that this preview is for, prefixed with `ctm_`.
	CustomerID *string `json:"customer_id,omitempty"`
	// AddressID: Paddle ID of the address that this preview is for, prefixed with `add_`. Send one of `address_id`, `customer_ip_address`, or the `address` object when previewing.
	AddressID *string `json:"address_id,omitempty"`
	// BusinessID: Paddle ID of the business that this preview is for, prefixed with `biz_`.
	BusinessID *string `json:"business_id,omitempty"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code.
	CurrencyCode *CurrencyCode `json:"currency_code,omitempty"`
	// DiscountID: Paddle ID of the discount applied to this preview, prefixed with `dsc_`.
	DiscountID *string `json:"discount_id,omitempty"`
	// Address: Address for this preview. Send one of `address_id`, `customer_ip_address`, or the `address` object when previewing.
	Address *AddressPreview `json:"address,omitempty"`
	// CustomerIPAddress: IP address for this transaction preview. Send one of `address_id`, `customer_ip_address`, or the `address` object when previewing.
	CustomerIPAddress *string `json:"customer_ip_address,omitempty"`
	// Details: Calculated totals for a price preview, including discounts, tax, and currency conversion.
	Details PricePreviewDetails `json:"details,omitempty"`
	// AvailablePaymentMethods: List of available payment methods for Paddle Checkout given the price and location information passed.
	AvailablePaymentMethods []PaymentMethodType `json:"available_payment_methods,omitempty"`
}

// PricingPreviewClient is a client for the Pricing preview resource.
type PricingPreviewClient struct {
	doer Doer
}

// PreviewPricesRequest is given as an input to PreviewPrices.
type PreviewPricesRequest struct {
	// Items: List of items to preview price calculations for.
	Items []PricePreviewItem `json:"items,omitempty"`
	// CustomerID: Paddle ID of the customer that this preview is for, prefixed with `ctm_`.
	CustomerID *string `json:"customer_id,omitempty"`
	// AddressID: Paddle ID of the address that this preview is for, prefixed with `add_`. Send one of `address_id`, `customer_ip_address`, or the `address` object when previewing.
	AddressID *string `json:"address_id,omitempty"`
	// BusinessID: Paddle ID of the business that this preview is for, prefixed with `biz_`.
	BusinessID *string `json:"business_id,omitempty"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code.
	CurrencyCode *CurrencyCode `json:"currency_code,omitempty"`
	// DiscountID: Paddle ID of the discount applied to this preview, prefixed with `dsc_`.
	DiscountID *string `json:"discount_id,omitempty"`
	// Address: Address for this preview. Send one of `address_id`, `customer_ip_address`, or the `address` object when previewing.
	Address *AddressPreview `json:"address,omitempty"`
	// CustomerIPAddress: IP address for this transaction preview. Send one of `address_id`, `customer_ip_address`, or the `address` object when previewing.
	CustomerIPAddress *string `json:"customer_ip_address,omitempty"`
}

// PreviewPrices performs the POST operation on a Pricing preview resource.
func (c *PricingPreviewClient) PreviewPrices(ctx context.Context, req *PreviewPricesRequest) (res *PricePreview, err error) {
	if err := c.doer.Do(ctx, "POST", "/pricing-preview", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
