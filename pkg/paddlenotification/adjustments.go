// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// AdjustmentCreated represents the adjustment.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type AdjustmentCreated struct {
	GenericNotificationEvent
	Data AdjustmentNotification `json:"data"`
}

// AdjustmentUpdated represents the adjustment.updated event.
// See https://developer.paddle.com/webhooks/overview for more information.
type AdjustmentUpdated struct {
	GenericNotificationEvent
	Data AdjustmentNotification `json:"data"`
}

// AdjustmentAction: How this adjustment impacts the related transaction..
type AdjustmentAction string

const (
	AdjustmentActionCredit            AdjustmentAction = "credit"
	AdjustmentActionRefund            AdjustmentAction = "refund"
	AdjustmentActionChargeback        AdjustmentAction = "chargeback"
	AdjustmentActionChargebackReverse AdjustmentAction = "chargeback_reverse"
	AdjustmentActionChargebackWarning AdjustmentAction = "chargeback_warning"
	AdjustmentActionCreditReverse     AdjustmentAction = "credit_reverse"
)

/*
AdjustmentStatus: Status of this adjustment. Set automatically by Paddle.

Most refunds for live accounts are created with the status of `pending_approval` until reviewed by Paddle, but some are automatically approved. For sandbox accounts, Paddle automatically approves refunds every ten minutes.

Credit adjustments don't require approval from Paddle, so they're created as `approved`..
*/
type AdjustmentStatus string

const (
	AdjustmentStatusPendingApproval AdjustmentStatus = "pending_approval"
	AdjustmentStatusApproved        AdjustmentStatus = "approved"
	AdjustmentStatusRejected        AdjustmentStatus = "rejected"
	AdjustmentStatusReversed        AdjustmentStatus = "reversed"
)

/*
AdjustmentType: Type of adjustment for this transaction item. `tax` adjustments are automatically created by Paddle.
Include `amount` when creating a `partial` adjustment..
*/
type AdjustmentType string

const (
	AdjustmentTypeFull      AdjustmentType = "full"
	AdjustmentTypePartial   AdjustmentType = "partial"
	AdjustmentTypeTax       AdjustmentType = "tax"
	AdjustmentTypeProration AdjustmentType = "proration"
)

// AdjustmentItemTotals: Breakdown of the total for an adjustment item.
type AdjustmentItemTotals struct {
	// Subtotal: Amount multiplied by quantity.
	Subtotal string `json:"subtotal,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after tax.
	Total string `json:"total,omitempty"`
}

// AdjustmentItem: List of items on this adjustment.
type AdjustmentItem struct {
	// ID: Unique Paddle ID for this adjustment item, prefixed with `adjitm_`.
	ID string `json:"id,omitempty"`
	// ItemID: Paddle ID for the transaction item that this adjustment item relates to, prefixed with `txnitm_`.
	ItemID string `json:"item_id,omitempty"`
	/*
	   Type: Type of adjustment for this transaction item. `tax` adjustments are automatically created by Paddle.
	   Include `amount` when creating a `partial` adjustment.
	*/
	Type AdjustmentType `json:"type,omitempty"`
	// Amount: Amount adjusted for this transaction item. Required when adjustment type is `partial`.
	Amount *string `json:"amount,omitempty"`
	// Proration: How proration was calculated for this adjustment item.
	Proration *Proration `json:"proration,omitempty"`
	// Totals: Breakdown of the total for an adjustment item.
	Totals AdjustmentItemTotals `json:"totals,omitempty"`
}

// AdjustmentTotals: Breakdown of the total for an adjustment.
type AdjustmentTotals struct {
	// Subtotal: Total before tax. For tax adjustments, the value is 0.
	Subtotal string `json:"subtotal,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after tax.
	Total string `json:"total,omitempty"`
	// Fee: Total fee taken by Paddle for this adjustment.
	Fee string `json:"fee,omitempty"`
	/*
	   Earnings: Total earnings. This is the subtotal minus the Paddle fee.
	   For tax adjustments, this value is negative, which means a positive effect in the transaction earnings.
	   This is because the fee is originally calculated from the transaction total, so if a tax adjustment is made,
	   then the fee portion of it is returned.
	*/
	Earnings string `json:"earnings,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code used for this adjustment.
	CurrencyCode CurrencyCode `json:"currency_code,omitempty"`
}

// PayoutTotalsAdjustment: Breakdown of how this adjustment affects your payout balance.
type PayoutTotalsAdjustment struct {
	// Subtotal: Adjustment total before tax and fees.
	Subtotal string `json:"subtotal,omitempty"`
	// Tax: Total tax on the adjustment subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Adjustment total after tax.
	Total string `json:"total,omitempty"`
	// Fee: Adjusted Paddle fee.
	Fee string `json:"fee,omitempty"`
	// ChargebackFee: Chargeback fees incurred for this adjustment. Only returned when the adjustment `action` is `chargeback` or `chargeback_warning`.
	ChargebackFee ChargebackFee `json:"chargeback_fee,omitempty"`
	// Earnings: Adjusted payout earnings. This is the adjustment total plus adjusted Paddle fees, excluding chargeback fees.
	Earnings string `json:"earnings,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code used for the payout for this transaction. If your primary currency has changed, this reflects the primary currency at the time the transaction was billed.
	CurrencyCode CurrencyCodePayouts `json:"currency_code,omitempty"`
}

// AdjustmentTaxRateUsedTotals: Calculated totals for the tax applied to this adjustment.
type AdjustmentTaxRateUsedTotals struct {
	// Subtotal: Total before tax. For tax adjustments, the value is 0.
	Subtotal string `json:"subtotal,omitempty"`
	// Tax: Total tax on the subtotal.
	Tax string `json:"tax,omitempty"`
	// Total: Total after tax.
	Total string `json:"total,omitempty"`
}

// AdjustmentTaxRateUsed: List of tax rates applied for this adjustment.
type AdjustmentTaxRateUsed struct {
	// TaxRate: Rate used to calculate tax for this adjustment.
	TaxRate string `json:"tax_rate,omitempty"`
	// Totals: Calculated totals for the tax applied to this adjustment.
	Totals AdjustmentTaxRateUsedTotals `json:"totals,omitempty"`
}

// AdjustmentNotification: New or changed entity.
type AdjustmentNotification struct {
	// ID: Unique Paddle ID for this adjustment entity, prefixed with `adj_`.
	ID string `json:"id,omitempty"`
	// Action: How this adjustment impacts the related transaction.
	Action AdjustmentAction `json:"action,omitempty"`
	// TransactionID: Paddle ID of the transaction that this adjustment is for, prefixed with `txn_`.
	TransactionID string `json:"transaction_id,omitempty"`
	/*
	   SubscriptionID: Paddle ID for the subscription related to this adjustment, prefixed with `sub_`.
	   Set automatically by Paddle based on the `subscription_id` of the related transaction.
	*/
	SubscriptionID *string `json:"subscription_id,omitempty"`
	/*
	   CustomerID: Paddle ID for the customer related to this adjustment, prefixed with `ctm_`.
	   Set automatically by Paddle based on the `customer_id` of the related transaction.
	*/
	CustomerID string `json:"customer_id,omitempty"`
	// Reason: Why this adjustment was created. Appears in the Paddle dashboard. Retained for record-keeping purposes.
	Reason string `json:"reason,omitempty"`
	/*
	   CreditAppliedToBalance: Whether this adjustment was applied to the related customer's credit balance. Only returned for `credit` adjustments.

	   `false` where the related transaction is `billed`. The adjustment reduces the amount due on the transaction.

	   `true` where the related transaction is `completed`. The amount is added the customer's credit balance and used to pay future transactions.
	*/
	CreditAppliedToBalance *bool `json:"credit_applied_to_balance,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code for this adjustment. Set automatically by Paddle based on the `currency_code` of the related transaction.
	CurrencyCode CurrencyCode `json:"currency_code,omitempty"`
	/*
	   Status: Status of this adjustment. Set automatically by Paddle.

	   Most refunds for live accounts are created with the status of `pending_approval` until reviewed by Paddle, but some are automatically approved. For sandbox accounts, Paddle automatically approves refunds every ten minutes.

	   Credit adjustments don't require approval from Paddle, so they're created as `approved`.
	*/
	Status AdjustmentStatus `json:"status,omitempty"`
	// Items: List of items on this adjustment.
	Items []AdjustmentItem `json:"items,omitempty"`
	// Totals: Breakdown of the total for an adjustment.
	Totals AdjustmentTotals `json:"totals,omitempty"`
	// PayoutTotals: Breakdown of how this adjustment affects your payout balance.
	PayoutTotals *PayoutTotalsAdjustment `json:"payout_totals,omitempty"`
	// TaxRatesUsed: List of tax rates applied for this adjustment.
	TaxRatesUsed []AdjustmentTaxRateUsed `json:"tax_rates_used,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
}
