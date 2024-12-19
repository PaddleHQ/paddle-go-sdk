// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"

	paddleerr "github.com/PaddleHQ/paddle-go-sdk/v2/pkg/paddleerr"
)

// ErrAdjustmentTransactionMissingCustomerID represents a `adjustment_transaction_missing_customer_id` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_missing_customer_id for more information.
var ErrAdjustmentTransactionMissingCustomerID = &paddleerr.Error{
	Code: "adjustment_transaction_missing_customer_id",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionCustomerMismatch represents a `adjustment_transaction_customer_mismatch` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_customer_mismatch for more information.
var ErrAdjustmentTransactionCustomerMismatch = &paddleerr.Error{
	Code: "adjustment_transaction_customer_mismatch",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionSubscriptionMismatch represents a `adjustment_transaction_subscription_mismatch` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_subscription_mismatch for more information.
var ErrAdjustmentTransactionSubscriptionMismatch = &paddleerr.Error{
	Code: "adjustment_transaction_subscription_mismatch",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionInvalidStatusForCredit represents a `adjustment_transaction_invalid_status_for_credit` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_invalid_status_for_credit for more information.
var ErrAdjustmentTransactionInvalidStatusForCredit = &paddleerr.Error{
	Code: "adjustment_transaction_invalid_status_for_credit",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionInvalidStatusForRefund represents a `adjustment_transaction_invalid_status_for_refund` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_invalid_status_for_refund for more information.
var ErrAdjustmentTransactionInvalidStatusForRefund = &paddleerr.Error{
	Code: "adjustment_transaction_invalid_status_for_refund",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentInvalidCreditAction represents a `adjustment_invalid_credit_action` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_invalid_credit_action for more information.
var ErrAdjustmentInvalidCreditAction = &paddleerr.Error{
	Code: "adjustment_invalid_credit_action",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentInvalidCombinationOfTypes represents a `adjustment_invalid_combination_of_types` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_invalid_combination_of_types for more information.
var ErrAdjustmentInvalidCombinationOfTypes = &paddleerr.Error{
	Code: "adjustment_invalid_combination_of_types",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentAmountAboveRemainingAllowed represents a `adjustment_amount_above_remaining_allowed` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_amount_above_remaining_allowed for more information.
var ErrAdjustmentAmountAboveRemainingAllowed = &paddleerr.Error{
	Code: "adjustment_amount_above_remaining_allowed",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTotalAmountAboveRemainingAllowed represents a `adjustment_total_amount_above_remaining_allowed` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_total_amount_above_remaining_allowed for more information.
var ErrAdjustmentTotalAmountAboveRemainingAllowed = &paddleerr.Error{
	Code: "adjustment_total_amount_above_remaining_allowed",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionItemHasAlreadyBeenFullyAdjusted represents a `adjustment_transaction_item_has_already_been_fully_adjusted` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_item_has_already_been_fully_adjusted for more information.
var ErrAdjustmentTransactionItemHasAlreadyBeenFullyAdjusted = &paddleerr.Error{
	Code: "adjustment_transaction_item_has_already_been_fully_adjusted",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentNoTaxAvailableToAdjust represents a `adjustment_no_tax_available_to_adjust` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_no_tax_available_to_adjust for more information.
var ErrAdjustmentNoTaxAvailableToAdjust = &paddleerr.Error{
	Code: "adjustment_no_tax_available_to_adjust",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentAmountCannotBeZero represents a `adjustment_amount_cannot_be_zero` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_amount_cannot_be_zero for more information.
var ErrAdjustmentAmountCannotBeZero = &paddleerr.Error{
	Code: "adjustment_amount_cannot_be_zero",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentPendingRefundRequest represents a `adjustment_pending_refund_request` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_pending_refund_request for more information.
var ErrAdjustmentPendingRefundRequest = &paddleerr.Error{
	Code: "adjustment_pending_refund_request",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionItemOverAdjustment represents a `adjustment_transaction_item_over_adjustment` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_item_over_adjustment for more information.
var ErrAdjustmentTransactionItemOverAdjustment = &paddleerr.Error{
	Code: "adjustment_transaction_item_over_adjustment",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentTransactionItemInvalid represents a `adjustment_transaction_item_invalid` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_transaction_item_invalid for more information.
var ErrAdjustmentTransactionItemInvalid = &paddleerr.Error{
	Code: "adjustment_transaction_item_invalid",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrAdjustmentCannotAdjustImportedTransaction represents a `adjustment_cannot_adjust_imported_transaction` error.
// See https://developer.paddle.com/errors/adjustments/adjustment_cannot_adjust_imported_transaction for more information.
var ErrAdjustmentCannotAdjustImportedTransaction = &paddleerr.Error{
	Code: "adjustment_cannot_adjust_imported_transaction",
	Type: paddleerr.ErrorTypeRequestError,
}

// AdjustmentActionQuery: Return entities for the specified action..
type AdjustmentActionQuery string

const (
	AdjustmentActionQueryChargeback        AdjustmentActionQuery = "chargeback"
	AdjustmentActionQueryChargebackReverse AdjustmentActionQuery = "chargeback_reverse"
	AdjustmentActionQueryChargebackWarning AdjustmentActionQuery = "chargeback_warning"
	AdjustmentActionQueryCredit            AdjustmentActionQuery = "credit"
	AdjustmentActionQueryCreditReverse     AdjustmentActionQuery = "credit_reverse"
	AdjustmentActionQueryRefund            AdjustmentActionQuery = "refund"
)

type AdjustmentCreditNotePDF struct {
	// URL: URL of the requested resource.
	URL string `json:"url,omitempty"`
}

// CustomerBalance: Totals for this credit balance. Where a customer has more than one subscription in this currency with a credit balance, includes totals for all subscriptions.
type CustomerBalance struct {
	// Available: Total amount of credit available to use.
	Available string `json:"available,omitempty"`
	// Reserved: Total amount of credit temporarily reserved for `billed` transactions.
	Reserved string `json:"reserved,omitempty"`
	// Used: Total amount of credit used.
	Used string `json:"used,omitempty"`
}

// CreditBalance: Represents a credit balance for a customer.
type CreditBalance struct {
	// CustomerID: Paddle ID of the customer that this credit balance is for, prefixed with `ctm_`.
	CustomerID string `json:"customer_id,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code for this credit balance.
	CurrencyCode CurrencyCode `json:"currency_code,omitempty"`
	// Balance: Totals for this credit balance. Where a customer has more than one subscription in this currency with a credit balance, includes totals for all subscriptions.
	Balance CustomerBalance `json:"balance,omitempty"`
}

// AdjustmentsClient is a client for the Adjustments resource.
type AdjustmentsClient struct {
	doer Doer
}

// ListAdjustmentsRequest is given as an input to ListAdjustments.
type ListAdjustmentsRequest struct {
	// Action is a query parameter.
	// Return entities for the specified action.
	Action *string `in:"query=action;omitempty" json:"-"`
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// CustomerID is a query parameter.
	// Return entities related to the specified customer. Use a comma-separated list to specify multiple customer IDs.
	CustomerID []string `in:"query=customer_id;omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `id`.
	*/
	OrderBy *string `in:"query=order_by;omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `10`; Maximum: `50`.
	*/
	PerPage *int `in:"query=per_page;omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status;omitempty" json:"-"`
	// SubscriptionID is a query parameter.
	// Return entities related to the specified subscription. Use a comma-separated list to specify multiple subscription IDs.
	SubscriptionID []string `in:"query=subscription_id;omitempty" json:"-"`
	// TransactionID is a query parameter.
	// Return entities related to the specified transaction. Use a comma-separated list to specify multiple transaction IDs.
	TransactionID []string `in:"query=transaction_id;omitempty" json:"-"`
	// ID is a query parameter.
	// Return only the IDs specified. Use a comma-separated list to get multiple entities.
	ID []string `in:"query=id;omitempty" json:"-"`
}

// ListAdjustments performs the GET operation on a Adjustments resource.
func (c *AdjustmentsClient) ListAdjustments(ctx context.Context, req *ListAdjustmentsRequest) (res *Collection[*Adjustment], err error) {
	if err := c.doer.Do(ctx, "GET", "/adjustments", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateAdjustmentRequest is given as an input to CreateAdjustment.
type CreateAdjustmentRequest struct {
	// Action: How this adjustment impacts the related transaction.
	Action AdjustmentAction `json:"action,omitempty"`
	// Reason: Why this adjustment was created. Appears in the Paddle dashboard. Retained for record-keeping purposes.
	Reason string `json:"reason,omitempty"`
	/*
	   TransactionID: Paddle ID of the transaction that this adjustment is for, prefixed with `txn_`.

	   Automatically-collected transactions must be `completed`; manually-collected transactions must have a status of `billed` or `past_due`

	   You can't create an adjustment for a transaction that has a refund that's pending approval.
	*/
	TransactionID string `json:"transaction_id,omitempty"`
	// Type: Type of adjustment. Use `full` to adjust the grand total for the related transaction. Include an `items` array when creating a `partial` adjustment. If omitted, defaults to `partial`.
	Type *AdjustmentType `json:"type,omitempty"`
	// Items: List of transaction items to adjust. Required if `type` is not populated or set to `partial`.
	Items []AdjustmentItem `json:"items,omitempty"`
}

// CreateAdjustment performs the POST operation on a Adjustments resource.
func (c *AdjustmentsClient) CreateAdjustment(ctx context.Context, req *CreateAdjustmentRequest) (res *Adjustment, err error) {
	if err := c.doer.Do(ctx, "POST", "/adjustments", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAdjustmentCreditNoteRequest is given as an input to GetAdjustmentCreditNote.
type GetAdjustmentCreditNoteRequest struct {
	// URL path parameters.
	AdjustmentID string `in:"path=adjustment_id" json:"-"`

	// Disposition is a query parameter.
	/*
	   Determine whether the generated URL should download the PDF as an attachment saved locally, or open it inline in the browser.

	   Default: `attachment`.
	*/
	Disposition *string `in:"query=disposition;omitempty" json:"-"`
}

// GetAdjustmentCreditNote performs the GET operation on a Adjustments resource.
func (c *AdjustmentsClient) GetAdjustmentCreditNote(ctx context.Context, req *GetAdjustmentCreditNoteRequest) (res *AdjustmentCreditNotePDF, err error) {
	if err := c.doer.Do(ctx, "GET", "/adjustments/{adjustment_id}/credit-note", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// ListCreditBalancesRequest is given as an input to ListCreditBalances.
type ListCreditBalancesRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`

	// CurrencyCode is a query parameter.
	// Return entities that match the currency code. Use a comma-separated list to specify multiple currency codes.
	CurrencyCode []string `in:"query=currency_code;omitempty" json:"-"`
}

// ListCreditBalances performs the GET operation on a Adjustments resource.
func (c *AdjustmentsClient) ListCreditBalances(ctx context.Context, req *ListCreditBalancesRequest) (res *Collection[*CreditBalance], err error) {
	if err := c.doer.Do(ctx, "GET", "/customers/{customer_id}/credit-balances", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
