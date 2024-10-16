// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// TransactionBilled represents the transaction.billed event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionBilled struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionCanceled represents the transaction.canceled event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionCanceled struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionCompleted represents the transaction.completed event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionCompleted struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionCreated represents the transaction.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionCreated struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionPaid represents the transaction.paid event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionPaid struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionPastDue represents the transaction.past_due event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionPastDue struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionPaymentFailed represents the transaction.payment_failed event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionPaymentFailed struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionReady represents the transaction.ready event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionReady struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionUpdated represents the transaction.updated event.
// See https://developer.paddle.com/webhooks/overview for more information.
type TransactionUpdated struct {
	GenericNotificationEvent
	Data TransactionNotification `json:"data"`
}

// TransactionStatus: Status of this transaction. You may set a transaction to `billed` or `canceled`, other statuses are set automatically by Paddle. Automatically-collected transactions may return `completed` if payment is captured successfully, or `past_due` if payment failed..
type TransactionStatus string

const (
	TransactionStatusDraft     TransactionStatus = "draft"
	TransactionStatusReady     TransactionStatus = "ready"
	TransactionStatusBilled    TransactionStatus = "billed"
	TransactionStatusPaid      TransactionStatus = "paid"
	TransactionStatusCompleted TransactionStatus = "completed"
	TransactionStatusCanceled  TransactionStatus = "canceled"
	TransactionStatusPastDue   TransactionStatus = "past_due"
)

// TransactionOrigin: Describes how this transaction was created..
type TransactionOrigin string

const (
	TransactionOriginAPI                             TransactionOrigin = "api"
	TransactionOriginSubscriptionCharge              TransactionOrigin = "subscription_charge"
	TransactionOriginSubscriptionPaymentMethodChange TransactionOrigin = "subscription_payment_method_change"
	TransactionOriginSubscriptionRecurring           TransactionOrigin = "subscription_recurring"
	TransactionOriginSubscriptionUpdate              TransactionOrigin = "subscription_update"
	TransactionOriginWeb                             TransactionOrigin = "web"
)

// TransactionItem: List of items on this transaction. For calculated totals, use `details.line_items`.
type TransactionItem struct {
	// PriceID: Paddle ID for the price to add to this transaction, prefixed with `pri_`.
	PriceID string `json:"price_id,omitempty"`
	// Price: Represents a price entity.
	Price Price `json:"price,omitempty"`
	// Quantity: Quantity of this item on the transaction.
	Quantity int `json:"quantity,omitempty"`
	// Proration: How proration was calculated for this item. Populated when a transaction is created from a subscription change, where `proration_billing_mode` was `prorated_immediately` or `prorated_next_billing_period`. Set automatically by Paddle.
	Proration *Proration `json:"proration,omitempty"`
}

// Totals: Calculated totals for the tax applied to this transaction.
type Totals struct {
	// Subtotal: Subtotal before discount, tax, and deductions. If an item, unit price multiplied by quantity.
	Subtotal string `json:"subtotal,omitempty"`
	/*
	   Discount: Total discount as a result of any discounts applied.

	   Except for percentage discounts, Paddle applies tax to discounts based on the line item `price.tax_mode`. If `price.tax_mode` for a line item is `internal`, Paddle removes tax from the discount applied.
	*/
	Discount string `json:"discount,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after discount and tax.
	Total string `json:"total,omitempty"`
}

// TaxRatesUsed: List of tax rates applied for this transaction.
type TaxRatesUsed struct {
	// TaxRate: Rate used to calculate tax for this transaction.
	TaxRate string `json:"tax_rate,omitempty"`
	// Totals: Calculated totals for the tax applied to this transaction.
	Totals Totals `json:"totals,omitempty"`
}

// TransactionTotals: Breakdown of the total for a transaction. These numbers can become negative when dealing with subscription updates that result in credit.
type TransactionTotals struct {
	// Subtotal: Subtotal before discount, tax, and deductions. If an item, unit price multiplied by quantity.
	Subtotal string `json:"subtotal,omitempty"`
	/*
	   Discount: Total discount as a result of any discounts applied.

	   Except for percentage discounts, Paddle applies tax to discounts based on the line item `price.tax_mode`. If `price.tax_mode` for a line item is `internal`, Paddle removes tax from the discount applied.
	*/
	Discount string `json:"discount,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after discount and tax.
	Total string `json:"total,omitempty"`
	// Credit: Total credit applied to this transaction. This includes credits applied using a customer's credit balance and adjustments to a `billed` transaction.
	Credit string `json:"credit,omitempty"`
	// CreditToBalance: Additional credit generated from negative `details.line_items`. This credit is added to the customer balance.
	CreditToBalance string `json:"credit_to_balance,omitempty"`
	// Balance: Total due on a transaction after credits and any payments.
	Balance string `json:"balance,omitempty"`
	// GrandTotal: Total due on a transaction after credits but before any payments.
	GrandTotal string `json:"grand_total,omitempty"`
	// Fee: Total fee taken by Paddle for this transaction. `null` until the transaction is `completed` and the fee is processed.
	Fee *string `json:"fee,omitempty"`
	// Earnings: Total earnings for this transaction. This is the total minus the Paddle fee. `null` until the transaction is `completed` and the fee is processed.
	Earnings *string `json:"earnings,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code of the currency used for this transaction.
	CurrencyCode CurrencyCode `json:"currency_code,omitempty"`
}

// TransactionTotalsAdjusted: Breakdown of the payout totals for a transaction after adjustments. `null` until the transaction is `completed`.
type TransactionTotalsAdjusted struct {
	// Subtotal: Subtotal before discount, tax, and deductions. If an item, unit price multiplied by quantity.
	Subtotal string `json:"subtotal,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after tax.
	Total string `json:"total,omitempty"`
	// GrandTotal: Total due after credits but before any payments.
	GrandTotal string `json:"grand_total,omitempty"`
	// Fee: Total fee taken by Paddle for this transaction. `null` until the transaction is `completed` and the fee is processed.
	Fee *string `json:"fee,omitempty"`
	/*
	   Earnings: Total earnings for this transaction. This is the total minus the Paddle fee.
	   `null` until the transaction is `completed` and the fee is processed.
	*/
	Earnings *string `json:"earnings,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code of the currency used for this transaction.
	CurrencyCode CurrencyCode `json:"currency_code,omitempty"`
}

// TransactionPayoutTotals: Breakdown of the payout total for a transaction. `null` until the transaction is `completed`. Returned in your payout currency.
type TransactionPayoutTotals struct {
	// Subtotal: Total before tax and fees.
	Subtotal string `json:"subtotal,omitempty"`
	/*
	   Discount: Total discount as a result of any discounts applied.
	   Except for percentage discounts, Paddle applies tax to discounts based on the line item `price.tax_mode`. If `price.tax_mode` for a line item is `internal`, Paddle removes tax from the discount applied.
	*/
	Discount string `json:"discount,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after tax.
	Total string `json:"total,omitempty"`
	// Credit: Total credit applied to this transaction. This includes credits applied using a customer's credit balance and adjustments to a `billed` transaction.
	Credit string `json:"credit,omitempty"`
	// CreditToBalance: Additional credit generated from negative `details.line_items`. This credit is added to the customer balance.
	CreditToBalance string `json:"credit_to_balance,omitempty"`
	// Balance: Total due on a transaction after credits and any payments.
	Balance string `json:"balance,omitempty"`
	// GrandTotal: Total due on a transaction after credits but before any payments.
	GrandTotal string `json:"grand_total,omitempty"`
	// Fee: Total fee taken by Paddle for this payout.
	Fee string `json:"fee,omitempty"`
	// Earnings: Total earnings for this payout. This is the subtotal minus the Paddle fee.
	Earnings string `json:"earnings,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code used for the payout for this transaction. If your primary currency has changed, this reflects the primary currency at the time the transaction was billed.
	CurrencyCode CurrencyCodePayouts `json:"currency_code,omitempty"`
}

// TransactionPayoutTotalsAdjusted: Breakdown of the payout total for a transaction after adjustments. `null` until the transaction is `completed`.
type TransactionPayoutTotalsAdjusted struct {
	// Subtotal: Total before tax and fees.
	Subtotal string `json:"subtotal,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after tax.
	Total string `json:"total,omitempty"`
	// Fee: Total fee taken by Paddle for this payout.
	Fee string `json:"fee,omitempty"`
	// ChargebackFee: Details of any chargeback fees incurred for this transaction.
	ChargebackFee ChargebackFee `json:"chargeback_fee,omitempty"`
	// Earnings: Total earnings for this payout. This is the subtotal minus the Paddle fee, excluding chargeback fees.
	Earnings string `json:"earnings,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code used for the payout for this transaction. If your primary currency has changed, this reflects the primary currency at the time the transaction was billed.
	CurrencyCode CurrencyCodePayouts `json:"currency_code,omitempty"`
}

// TransactionLineItem: Information about line items for this transaction. Different from transaction `items` as they include totals calculated by Paddle. Considered the source of truth for line item totals.
type TransactionLineItem struct {
	// ID: Unique Paddle ID for this transaction item, prefixed with `txnitm_`. Used when working with [adjustments](https://developer.paddle.com/build/transactions/create-transaction-adjustments).
	ID string `json:"id,omitempty"`
	// PriceID: Paddle ID for the price related to this transaction line item, prefixed with `pri_`.
	PriceID string `json:"price_id,omitempty"`
	// Quantity: Quantity of this transaction line item.
	Quantity int `json:"quantity,omitempty"`
	// Proration: How proration was calculated for this item. Populated when a transaction is created from a subscription change, where `proration_billing_mode` was `prorated_immediately` or `prorated_next_billing_period`. Set automatically by Paddle.
	Proration *Proration `json:"proration,omitempty"`
	// TaxRate: Rate used to calculate tax for this transaction line item.
	TaxRate string `json:"tax_rate,omitempty"`
	// UnitTotals: Breakdown of the charge for one unit in the lowest denomination of a currency (e.g. cents for USD).
	UnitTotals Totals `json:"unit_totals,omitempty"`
	// Totals: Breakdown of a charge in the lowest denomination of a currency (e.g. cents for USD).
	Totals Totals `json:"totals,omitempty"`
	// Product: Related product entity for this transaction line item price.
	Product Product `json:"product,omitempty"`
}

// TransactionDetails: Calculated totals for a transaction, including proration, discounts, tax, and currency conversion. Considered the source of truth for totals on a transaction.
type TransactionDetails struct {
	// TaxRatesUsed: List of tax rates applied for this transaction.
	TaxRatesUsed []TaxRatesUsed `json:"tax_rates_used,omitempty"`
	// Totals: Breakdown of the total for a transaction. These numbers can become negative when dealing with subscription updates that result in credit.
	Totals TransactionTotals `json:"totals,omitempty"`
	// AdjustedTotals: Breakdown of the payout totals for a transaction after adjustments. `null` until the transaction is `completed`.
	AdjustedTotals TransactionTotalsAdjusted `json:"adjusted_totals,omitempty"`
	// PayoutTotals: Breakdown of the payout total for a transaction. `null` until the transaction is `completed`. Returned in your payout currency.
	PayoutTotals *TransactionPayoutTotals `json:"payout_totals,omitempty"`
	// AdjustedPayoutTotals: Breakdown of the payout total for a transaction after adjustments. `null` until the transaction is `completed`.
	AdjustedPayoutTotals *TransactionPayoutTotalsAdjusted `json:"adjusted_payout_totals,omitempty"`
	// LineItems: Information about line items for this transaction. Different from transaction `items` as they include totals calculated by Paddle. Considered the source of truth for line item totals.
	LineItems []TransactionLineItem `json:"line_items,omitempty"`
}

// PaymentAttemptStatus: Status of this payment attempt..
type PaymentAttemptStatus string

const (
	PaymentAttemptStatusAuthorized              PaymentAttemptStatus = "authorized"
	PaymentAttemptStatusAuthorizedFlagged       PaymentAttemptStatus = "authorized_flagged"
	PaymentAttemptStatusCanceled                PaymentAttemptStatus = "canceled"
	PaymentAttemptStatusCaptured                PaymentAttemptStatus = "captured"
	PaymentAttemptStatusError                   PaymentAttemptStatus = "error"
	PaymentAttemptStatusActionRequired          PaymentAttemptStatus = "action_required"
	PaymentAttemptStatusPendingNoActionRequired PaymentAttemptStatus = "pending_no_action_required"
	PaymentAttemptStatusCreated                 PaymentAttemptStatus = "created"
	PaymentAttemptStatusUnknown                 PaymentAttemptStatus = "unknown"
	PaymentAttemptStatusDropped                 PaymentAttemptStatus = "dropped"
)

// ErrorCode: Reason why a payment attempt failed. Returns `null` if payment captured successfully..
type ErrorCode string

const (
	ErrorCodeAlreadyCanceled         ErrorCode = "already_canceled"
	ErrorCodeAlreadyRefunded         ErrorCode = "already_refunded"
	ErrorCodeAuthenticationFailed    ErrorCode = "authentication_failed"
	ErrorCodeBlockedCard             ErrorCode = "blocked_card"
	ErrorCodeCanceled                ErrorCode = "canceled"
	ErrorCodeDeclined                ErrorCode = "declined"
	ErrorCodeDeclinedNotRetryable    ErrorCode = "declined_not_retryable"
	ErrorCodeExpiredCard             ErrorCode = "expired_card"
	ErrorCodeFraud                   ErrorCode = "fraud"
	ErrorCodeInvalidAmount           ErrorCode = "invalid_amount"
	ErrorCodeInvalidPaymentDetails   ErrorCode = "invalid_payment_details"
	ErrorCodeIssuerUnavailable       ErrorCode = "issuer_unavailable"
	ErrorCodeNotEnoughBalance        ErrorCode = "not_enough_balance"
	ErrorCodePspError                ErrorCode = "psp_error"
	ErrorCodeRedactedPaymentMethod   ErrorCode = "redacted_payment_method"
	ErrorCodeSystemError             ErrorCode = "system_error"
	ErrorCodeTransactionNotPermitted ErrorCode = "transaction_not_permitted"
	ErrorCodeUnknown                 ErrorCode = "unknown"
)

// PaymentMethodType: Type of payment method used for this payment attempt..
type PaymentMethodType string

const (
	PaymentMethodTypeAlipay       PaymentMethodType = "alipay"
	PaymentMethodTypeApplePay     PaymentMethodType = "apple_pay"
	PaymentMethodTypeBancontact   PaymentMethodType = "bancontact"
	PaymentMethodTypeCard         PaymentMethodType = "card"
	PaymentMethodTypeGooglePay    PaymentMethodType = "google_pay"
	PaymentMethodTypeIdeal        PaymentMethodType = "ideal"
	PaymentMethodTypeOffline      PaymentMethodType = "offline"
	PaymentMethodTypePaypal       PaymentMethodType = "paypal"
	PaymentMethodTypeUnknown      PaymentMethodType = "unknown"
	PaymentMethodTypeWireTransfer PaymentMethodType = "wire_transfer"
)

// CardType: Type of credit or debit card used to pay..
type CardType string

const (
	CardTypeAmericanExpress CardType = "american_express"
	CardTypeDinersClub      CardType = "diners_club"
	CardTypeDiscover        CardType = "discover"
	CardTypeJcb             CardType = "jcb"
	CardTypeMada            CardType = "mada"
	CardTypeMaestro         CardType = "maestro"
	CardTypeMastercard      CardType = "mastercard"
	CardTypeUnionPay        CardType = "union_pay"
	CardTypeUnknown         CardType = "unknown"
	CardTypeVisa            CardType = "visa"
)

// Card: Information about the credit or debit card used to pay. `null` unless `type` is `card`.
type Card struct {
	// Type: Type of credit or debit card used to pay.
	Type CardType `json:"type,omitempty"`
	// Last4: Last four digits of the card used to pay.
	Last4 string `json:"last4,omitempty"`
	// ExpiryMonth: Month of the expiry date of the card used to pay.
	ExpiryMonth int `json:"expiry_month,omitempty"`
	// ExpiryYear: Year of the expiry date of the card used to pay.
	ExpiryYear int `json:"expiry_year,omitempty"`
	// CardholderName: The name on the card used to pay.
	CardholderName string `json:"cardholder_name,omitempty"`
}

// MethodDetails: Information about the payment method used for a payment attempt.
type MethodDetails struct {
	// Type: Type of payment method used for this payment attempt.
	Type PaymentMethodType `json:"type,omitempty"`
	// Card: Information about the credit or debit card used to pay. `null` unless `type` is `card`.
	Card *Card `json:"card,omitempty"`
}

// TransactionPaymentAttempt: List of payment attempts for this transaction, including successful payments. Sorted by `created_at` in descending order, so most recent attempts are returned first.
type TransactionPaymentAttempt struct {
	// PaymentAttemptID: UUID for this payment attempt.
	PaymentAttemptID string `json:"payment_attempt_id,omitempty"`
	// StoredPaymentMethodID: UUID for the stored payment method used for this payment attempt. Deprecated - use `payment_method_id` instead.
	StoredPaymentMethodID string `json:"stored_payment_method_id,omitempty"`
	// PaymentMethodID: Paddle ID of the payment method used for this payment attempt, prefixed with `paymtd_`.
	PaymentMethodID *string `json:"payment_method_id,omitempty"`
	// Amount: Amount for collection in the lowest denomination of a currency (e.g. cents for USD).
	Amount string `json:"amount,omitempty"`
	// Status: Status of this payment attempt.
	Status PaymentAttemptStatus `json:"status,omitempty"`
	// ErrorCode: Reason why a payment attempt failed. Returns `null` if payment captured successfully.
	ErrorCode *ErrorCode `json:"error_code,omitempty"`
	// MethodDetails: Information about the payment method used for a payment attempt.
	MethodDetails MethodDetails `json:"method_details,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// CapturedAt: RFC 3339 datetime string of when this payment was captured. `null` if `status` is not `captured`.
	CapturedAt *string `json:"captured_at,omitempty"`
}

// Checkout: Paddle Checkout details for this transaction. Returned for automatically-collected transactions and where `billing_details.enable_checkout` is `true` for manually-collected transactions; `null` otherwise.
type Checkout struct {
	// URL: Paddle Checkout URL for this transaction, composed of the URL passed in the request or your default payment URL + `_?txn=` and the Paddle ID for this transaction.
	URL *string `json:"url,omitempty"`
}

// TransactionNotification: New or changed entity.
type TransactionNotification struct {
	// ID: Unique Paddle ID for this transaction entity, prefixed with `txn_`.
	ID string `json:"id"`
	// Status: Status of this transaction. You may set a transaction to `billed` or `canceled`, other statuses are set automatically by Paddle. Automatically-collected transactions may return `completed` if payment is captured successfully, or `past_due` if payment failed.
	Status TransactionStatus `json:"status"`
	// CustomerID: Paddle ID of the customer that this transaction is for, prefixed with `ctm_`.
	CustomerID *string `json:"customer_id"`
	// AddressID: Paddle ID of the address that this transaction is for, prefixed with `add_`.
	AddressID *string `json:"address_id"`
	// BusinessID: Paddle ID of the business that this transaction is for, prefixed with `biz_`.
	BusinessID *string `json:"business_id"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code. Must be `USD`, `EUR`, or `GBP` if `collection_mode` is `manual`.
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Origin: Describes how this transaction was created.
	Origin TransactionOrigin `json:"origin"`
	// SubscriptionID: Paddle ID of the subscription that this transaction is for, prefixed with `sub_`.
	SubscriptionID *string `json:"subscription_id"`
	// InvoiceID: Paddle ID of the invoice that this transaction is related to, prefixed with `inv_`. Used for compatibility with the Paddle Invoice API, which is now deprecated. This field is scheduled to be removed in the next version of the Paddle API.
	InvoiceID *string `json:"invoice_id"`
	// InvoiceNumber: Invoice number for this transaction. Automatically generated by Paddle when you mark a transaction as `billed` where `collection_mode` is `manual`.
	InvoiceNumber *string `json:"invoice_number"`
	// CollectionMode: How payment is collected for this transaction. `automatic` for checkout, `manual` for invoices.
	CollectionMode CollectionMode `json:"collection_mode"`
	// DiscountID: Paddle ID of the discount applied to this transaction, prefixed with `dsc_`.
	DiscountID *string `json:"discount_id"`
	// BillingDetails: Details for invoicing. Required if `collection_mode` is `manual`.
	BillingDetails *BillingDetails `json:"billing_details"`
	// BillingPeriod: Time period that this transaction is for. Set automatically by Paddle for subscription renewals to describe the period that charges are for.
	BillingPeriod *TimePeriod `json:"billing_period"`
	// Items: List of items on this transaction. For calculated totals, use `details.line_items`.
	Items []TransactionItem `json:"items"`
	// Details: Calculated totals for a transaction, including proration, discounts, tax, and currency conversion. Considered the source of truth for totals on a transaction.
	Details TransactionDetails `json:"details"`
	// Payments: List of payment attempts for this transaction, including successful payments. Sorted by `created_at` in descending order, so most recent attempts are returned first.
	Payments []TransactionPaymentAttempt `json:"payments"`
	// Checkout: Paddle Checkout details for this transaction. Returned for automatically-collected transactions and where `billing_details.enable_checkout` is `true` for manually-collected transactions; `null` otherwise.
	Checkout *Checkout `json:"checkout"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at"`
	// BilledAt: RFC 3339 datetime string of when this transaction was marked as `billed`. `null` for transactions that are not `billed` or `completed`. Set automatically by Paddle.
	BilledAt *string `json:"billed_at"`
}
