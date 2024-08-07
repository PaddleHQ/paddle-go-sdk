// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// ReportCreated represents the report.created event.
// See https://developer.paddle.com/webhooks/overview for more information.
type ReportCreated struct {
	GenericNotificationsEvent
	Data ReportNotification `json:"data"`
}

// ReportUpdated represents the report.updated event.
// See https://developer.paddle.com/webhooks/overview for more information.
type ReportUpdated struct {
	GenericNotificationsEvent
	Data ReportNotification `json:"data"`
}

/*
ReportsStatus: Status of this report. Set automatically by Paddle.

Reports are created as `pending` initially, then move to `ready` when they're available to download..
*/
type ReportsStatus string

const (
	ReportsStatusPending ReportsStatus = "pending"
	ReportsStatusReady   ReportsStatus = "ready"
	ReportsStatusFailed  ReportsStatus = "failed"
	ReportsStatusExpired ReportsStatus = "expired"
)

// ReportsType: Type of report to create..
type ReportsType string

const (
	ReportsTypeAdjustments          ReportsType = "adjustments"
	ReportsTypeAdjustmentLineItems  ReportsType = "adjustment_line_items"
	ReportsTypeTransactions         ReportsType = "transactions"
	ReportsTypeTransactionLineItems ReportsType = "transaction_line_items"
	ReportsTypeProductsPrices       ReportsType = "products_prices"
	ReportsTypeDiscounts            ReportsType = "discounts"
)

// ReportsFiltersName: Field name to filter by..
type ReportsFiltersName string

const (
	ReportsFiltersNameAction           ReportsFiltersName = "action"
	ReportsFiltersNameCurrencyCode     ReportsFiltersName = "currency_code"
	ReportsFiltersNameStatus           ReportsFiltersName = "status"
	ReportsFiltersNameUpdatedAt        ReportsFiltersName = "updated_at"
	ReportsFiltersNameCollectionMode   ReportsFiltersName = "collection_mode"
	ReportsFiltersNameOrigin           ReportsFiltersName = "origin"
	ReportsFiltersNameProductStatus    ReportsFiltersName = "product_status"
	ReportsFiltersNamePriceStatus      ReportsFiltersName = "price_status"
	ReportsFiltersNameProductType      ReportsFiltersName = "product_type"
	ReportsFiltersNamePriceType        ReportsFiltersName = "price_type"
	ReportsFiltersNameProductUpdatedAt ReportsFiltersName = "product_updated_at"
	ReportsFiltersNamePriceUpdatedAt   ReportsFiltersName = "price_updated_at"
	ReportsFiltersNameType             ReportsFiltersName = "type"
)

// ReportsFiltersOperator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise..
type ReportsFiltersOperator string

const (
	ReportsFiltersOperatorLt  ReportsFiltersOperator = "lt"
	ReportsFiltersOperatorGte ReportsFiltersOperator = "gte"
)

// ReportsFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type ReportsFilters struct {
	// Name: Field name to filter by.
	Name ReportsFiltersName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *ReportsFiltersOperator `json:"operator,omitempty"`
	// Value: Value to filter by. Check the allowed values descriptions for the `name` field to see valid values for a field.
	Value any `json:"value,omitempty"`
}

// ReportNotification: New or changed entity.
type ReportNotification struct {
	// ID: Unique Paddle ID for this report, prefixed with `rep_`
	ID string `json:"id,omitempty"`
	/*
	   Status: Status of this report. Set automatically by Paddle.

	   Reports are created as `pending` initially, then move to `ready` when they're available to download.
	*/
	Status ReportsStatus `json:"status,omitempty"`
	// Rows: Number of records in this report. `null` if the report is `pending`.
	Rows *int `json:"rows,omitempty"`
	// ExpiresAt: RFC 3339 datetime string of when this report expires. The report is no longer available to download after this date.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this report was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this report was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Type: Type of report to create.
	Type ReportsType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []ReportsFilters `json:"filters,omitempty"`
}
