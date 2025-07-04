// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"
	"encoding/json"
	paddleerr "github.com/PaddleHQ/paddle-go-sdk/v4/pkg/paddleerr"
)

// ErrReportCreationLimitExceeded represents a `report_creation_limit_exceeded` error.
// See https://developer.paddle.com/errors/reports/report_creation_limit_exceeded for more information.
var ErrReportCreationLimitExceeded = &paddleerr.Error{
	Code: "report_creation_limit_exceeded",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrConcurrentReportGenerationNotAllowed represents a `concurrent_report_generation_not_allowed` error.
// See https://developer.paddle.com/errors/reports/concurrent_report_generation_not_allowed for more information.
var ErrConcurrentReportGenerationNotAllowed = &paddleerr.Error{
	Code: "concurrent_report_generation_not_allowed",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrReportNotReady represents a `report_not_ready` error.
// See https://developer.paddle.com/errors/reports/report_not_ready for more information.
var ErrReportNotReady = &paddleerr.Error{
	Code: "report_not_ready",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrReportExpired represents a `report_expired` error.
// See https://developer.paddle.com/errors/reports/report_expired for more information.
var ErrReportExpired = &paddleerr.Error{
	Code: "report_expired",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrReportFailed represents a `report_failed` error.
// See https://developer.paddle.com/errors/reports/report_failed for more information.
var ErrReportFailed = &paddleerr.Error{
	Code: "report_failed",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrReportEmpty represents a `report_empty` error.
// See https://developer.paddle.com/errors/reports/report_empty for more information.
var ErrReportEmpty = &paddleerr.Error{
	Code: "report_empty",
	Type: paddleerr.ErrorTypeRequestError,
}

/*
ReportStatus: Status of this report. Set automatically by Paddle.

Reports are created as `pending` initially, then move to `ready` when they're available to download..
*/
type ReportStatus string

const (
	ReportStatusPending ReportStatus = "pending"
	ReportStatusReady   ReportStatus = "ready"
	ReportStatusFailed  ReportStatus = "failed"
	ReportStatusExpired ReportStatus = "expired"
)

// ReportType: Type of report to create..
type ReportType string

const (
	ReportTypeAdjustments          ReportType = "adjustments"
	ReportTypeAdjustmentLineItems  ReportType = "adjustment_line_items"
	ReportTypeTransactions         ReportType = "transactions"
	ReportTypeTransactionLineItems ReportType = "transaction_line_items"
	ReportTypeProductsPrices       ReportType = "products_prices"
	ReportTypeDiscounts            ReportType = "discounts"
	ReportTypeBalance              ReportType = "balance"
)

// ReportFiltersName: Field name to filter by..
type ReportFiltersName string

const (
	ReportFiltersNameAction           ReportFiltersName = "action"
	ReportFiltersNameCurrencyCode     ReportFiltersName = "currency_code"
	ReportFiltersNameStatus           ReportFiltersName = "status"
	ReportFiltersNameUpdatedAt        ReportFiltersName = "updated_at"
	ReportFiltersNameCollectionMode   ReportFiltersName = "collection_mode"
	ReportFiltersNameOrigin           ReportFiltersName = "origin"
	ReportFiltersNameProductStatus    ReportFiltersName = "product_status"
	ReportFiltersNamePriceStatus      ReportFiltersName = "price_status"
	ReportFiltersNameProductType      ReportFiltersName = "product_type"
	ReportFiltersNamePriceType        ReportFiltersName = "price_type"
	ReportFiltersNameProductUpdatedAt ReportFiltersName = "product_updated_at"
	ReportFiltersNamePriceUpdatedAt   ReportFiltersName = "price_updated_at"
	ReportFiltersNameType             ReportFiltersName = "type"
)

// FilterOperator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise..
type FilterOperator string

const (
	FilterOperatorLt  FilterOperator = "lt"
	FilterOperatorGte FilterOperator = "gte"
)

// ReportFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type ReportFilters struct {
	// Name: Field name to filter by.
	Name ReportFiltersName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *FilterOperator `json:"operator,omitempty"`
	// Value: Value to filter by. Check the allowed values descriptions for the `name` field to see valid values for a field.
	Value any `json:"value,omitempty"`
}

// Report: Represents a report entity.
type Report struct {
	// ID: Unique Paddle ID for this report, prefixed with `rep_`
	ID string `json:"id,omitempty"`
	/*
	   Status: Status of this report. Set automatically by Paddle.

	   Reports are created as `pending` initially, then move to `ready` when they're available to download.
	*/
	Status ReportStatus `json:"status,omitempty"`
	// Rows: Number of records in this report. `null` if the report is `pending`.
	Rows *int `json:"rows,omitempty"`
	// ExpiresAt: RFC 3339 datetime string of when this report expires. The report is no longer available to download after this date.
	ExpiresAt *string `json:"expires_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this report was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this report was created.
	CreatedAt string `json:"created_at,omitempty"`
	// Type: Type of report to create.
	Type ReportType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []ReportFilters `json:"filters,omitempty"`
}

// AdjustmentsReportType: Type of report to create..
type AdjustmentsReportType string

const (
	AdjustmentsReportTypeAdjustments         AdjustmentsReportType = "adjustments"
	AdjustmentsReportTypeAdjustmentLineItems AdjustmentsReportType = "adjustment_line_items"
)

// AdjustmentsReportFilterName: Field name to filter by..
type AdjustmentsReportFilterName string

const (
	AdjustmentsReportFilterNameAction       AdjustmentsReportFilterName = "action"
	AdjustmentsReportFilterNameCurrencyCode AdjustmentsReportFilterName = "currency_code"
	AdjustmentsReportFilterNameStatus       AdjustmentsReportFilterName = "status"
	AdjustmentsReportFilterNameUpdatedAt    AdjustmentsReportFilterName = "updated_at"
)

// AdjustmentsReportFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type AdjustmentsReportFilters struct {
	// Name: Field name to filter by.
	Name AdjustmentsReportFilterName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *FilterOperator `json:"operator,omitempty"`
	// Value: Value to filter by. Check the allowed values descriptions for the `name` field to see valid values for a field.
	Value any `json:"value,omitempty"`
}

// AdjustmentsReports: Entity when working with reports for adjustments or adjustment line items.
type AdjustmentsReports struct {
	// Type: Type of report to create.
	Type AdjustmentsReportType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []AdjustmentsReportFilters `json:"filters,omitempty"`
}

// TransactionsReportType: Type of report to create..
type TransactionsReportType string

const (
	TransactionsReportTypeTransactions         TransactionsReportType = "transactions"
	TransactionsReportTypeTransactionLineItems TransactionsReportType = "transaction_line_items"
)

// TransactionsReportFilterName: Field name to filter by..
type TransactionsReportFilterName string

const (
	TransactionsReportFilterNameCollectionMode TransactionsReportFilterName = "collection_mode"
	TransactionsReportFilterNameCurrencyCode   TransactionsReportFilterName = "currency_code"
	TransactionsReportFilterNameOrigin         TransactionsReportFilterName = "origin"
	TransactionsReportFilterNameStatus         TransactionsReportFilterName = "status"
	TransactionsReportFilterNameUpdatedAt      TransactionsReportFilterName = "updated_at"
)

// TransactionsReportFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type TransactionsReportFilters struct {
	// Name: Field name to filter by.
	Name TransactionsReportFilterName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *FilterOperator `json:"operator,omitempty"`
	// Value: Value to filter by. Check the allowed values descriptions for the `name` field to see valid values for a field.
	Value any `json:"value,omitempty"`
}

// TransactionsReports: Entity when working with reports for transaction or transaction line items.
type TransactionsReports struct {
	// Type: Type of report to create.
	Type TransactionsReportType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []TransactionsReportFilters `json:"filters,omitempty"`
}

// ProductsPricesReportType: Type of report to create..
type ProductsPricesReportType string

const ProductsPricesReportTypeProductsPrices ProductsPricesReportType = "products_prices"

// ProductPricesReportFilterName: Field name to filter by..
type ProductPricesReportFilterName string

const (
	ProductPricesReportFilterNameProductStatus    ProductPricesReportFilterName = "product_status"
	ProductPricesReportFilterNamePriceStatus      ProductPricesReportFilterName = "price_status"
	ProductPricesReportFilterNameProductType      ProductPricesReportFilterName = "product_type"
	ProductPricesReportFilterNamePriceType        ProductPricesReportFilterName = "price_type"
	ProductPricesReportFilterNameProductUpdatedAt ProductPricesReportFilterName = "product_updated_at"
	ProductPricesReportFilterNamePriceUpdatedAt   ProductPricesReportFilterName = "price_updated_at"
)

// ProductPricesReportFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `product_updated_at` and `price_updated_at` are greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type ProductPricesReportFilters struct {
	// Name: Field name to filter by.
	Name ProductPricesReportFilterName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *FilterOperator `json:"operator,omitempty"`
	// Value: Value to filter by.
	Value any `json:"value,omitempty"`
}

// ProductsAndPricesReport: Entity when working with a products and prices report.
type ProductsAndPricesReport struct {
	// Type: Type of report to create.
	Type ProductsPricesReportType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `product_updated_at` and `price_updated_at` are greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []ProductPricesReportFilters `json:"filters,omitempty"`
}

// DiscountsReportType: Type of report to create..
type DiscountsReportType string

const DiscountsReportTypeDiscounts DiscountsReportType = "discounts"

// DiscountsReportFilterName: Field name to filter by..
type DiscountsReportFilterName string

const (
	DiscountsReportFilterNameType      DiscountsReportFilterName = "type"
	DiscountsReportFilterNameStatus    DiscountsReportFilterName = "status"
	DiscountsReportFilterNameUpdatedAt DiscountsReportFilterName = "updated_at"
)

// DiscountsReportFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type DiscountsReportFilters struct {
	// Name: Field name to filter by.
	Name DiscountsReportFilterName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *FilterOperator `json:"operator,omitempty"`
	// Value: Value to filter by. Check the allowed values descriptions for the `name` field to see valid values for a field.
	Value any `json:"value,omitempty"`
}

// DiscountsReport: Entity when working with a discounts report.
type DiscountsReport struct {
	// Type: Type of report to create.
	Type DiscountsReportType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []DiscountsReportFilters `json:"filters,omitempty"`
}

// BalanceReportType: Type of report to create..
type BalanceReportType string

const BalanceReportTypeBalance BalanceReportType = "balance"

// BalanceReportFilterName: Field name to filter by..
type BalanceReportFilterName string

const BalanceReportFilterNameUpdatedAt BalanceReportFilterName = "updated_at"

// BalanceReportFilters: Filter criteria for this report. If omitted when creating, reports are filtered to include Paddle Billing data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
type BalanceReportFilters struct {
	// Name: Field name to filter by.
	Name BalanceReportFilterName `json:"name,omitempty"`
	// Operator: Operator to use when filtering. Valid when filtering by `updated_at`, `null` otherwise.
	Operator *FilterOperator `json:"operator,omitempty"`
	// Value: Value to filter by. Check the allowed values descriptions for the `name` field to see valid values for a field.
	Value any `json:"value,omitempty"`
}

// BalanceReport: Entity when working with a balance report.
type BalanceReport struct {
	// Type: Type of report to create.
	Type BalanceReportType `json:"type,omitempty"`
	// Filters: Filter criteria for this report. If omitted when creating, reports are filtered to include Paddle Billing data updated in the last 30 days. This means `updated_at` is greater than or equal to (`gte`) the date 30 days ago from the time the report was generated.
	Filters []BalanceReportFilters `json:"filters,omitempty"`
}

type ReportCSV struct {
	// URL: URL of the requested resource.
	URL string `json:"url,omitempty"`
}

// ReportsClient is a client for the Reports resource.
type ReportsClient struct {
	doer Doer
}

// ListReportsRequest is given as an input to ListReports.
type ListReportsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `id`.
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

// ListReports performs the GET operation on a Reports resource.
func (c *ReportsClient) ListReports(ctx context.Context, req *ListReportsRequest) (res *Collection[*Report], err error) {
	if err := c.doer.Do(ctx, "GET", "/reports", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// NewCreateReportRequestAdjustmentsReports takes a AdjustmentsReports type
// and creates a CreateReportRequest for use in a request.
func NewCreateReportRequestAdjustmentsReports(r *AdjustmentsReports) *CreateReportRequest {
	return &CreateReportRequest{AdjustmentsReports: r}
}

// NewCreateReportRequestTransactionsReports takes a TransactionsReports type
// and creates a CreateReportRequest for use in a request.
func NewCreateReportRequestTransactionsReports(r *TransactionsReports) *CreateReportRequest {
	return &CreateReportRequest{TransactionsReports: r}
}

// NewCreateReportRequestProductsAndPricesReport takes a ProductsAndPricesReport type
// and creates a CreateReportRequest for use in a request.
func NewCreateReportRequestProductsAndPricesReport(r *ProductsAndPricesReport) *CreateReportRequest {
	return &CreateReportRequest{ProductsAndPricesReport: r}
}

// NewCreateReportRequestDiscountsReport takes a DiscountsReport type
// and creates a CreateReportRequest for use in a request.
func NewCreateReportRequestDiscountsReport(r *DiscountsReport) *CreateReportRequest {
	return &CreateReportRequest{DiscountsReport: r}
}

// NewCreateReportRequestBalanceReport takes a BalanceReport type
// and creates a CreateReportRequest for use in a request.
func NewCreateReportRequestBalanceReport(r *BalanceReport) *CreateReportRequest {
	return &CreateReportRequest{BalanceReport: r}
}

// CreateReportRequest represents a union request type of the following types:
//   - `AdjustmentsReports`
//   - `TransactionsReports`
//   - `ProductsAndPricesReport`
//   - `DiscountsReport`
//   - `BalanceReport`
//
// The following constructor functions can be used to create a new instance of this type.
//   - `NewCreateReportRequestAdjustmentsReports()`
//   - `NewCreateReportRequestTransactionsReports()`
//   - `NewCreateReportRequestProductsAndPricesReport()`
//   - `NewCreateReportRequestDiscountsReport()`
//   - `NewCreateReportRequestBalanceReport()`
//
// Only one of the values can be set at a time, the first non-nil value will be used in the request.
type CreateReportRequest struct {
	*AdjustmentsReports
	*TransactionsReports
	*ProductsAndPricesReport
	*DiscountsReport
	*BalanceReport
}

// CreateReport performs the POST operation on a Reports resource.
func (c *ReportsClient) CreateReport(ctx context.Context, req *CreateReportRequest) (res *Report, err error) {
	if err := c.doer.Do(ctx, "POST", "/reports", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (u CreateReportRequest) MarshalJSON() ([]byte, error) {
	if u.AdjustmentsReports != nil {
		return json.Marshal(u.AdjustmentsReports)
	}

	if u.TransactionsReports != nil {
		return json.Marshal(u.TransactionsReports)
	}

	if u.ProductsAndPricesReport != nil {
		return json.Marshal(u.ProductsAndPricesReport)
	}

	if u.DiscountsReport != nil {
		return json.Marshal(u.DiscountsReport)
	}

	if u.BalanceReport != nil {
		return json.Marshal(u.BalanceReport)
	}

	return nil, nil
}

// GetReportCSVRequest is given as an input to GetReportCSV.
type GetReportCSVRequest struct {
	// URL path parameters.
	ReportID string `in:"path=report_id" json:"-"`
}

// GetReportCSV performs the GET operation on a Reports resource.
func (c *ReportsClient) GetReportCSV(ctx context.Context, req *GetReportCSVRequest) (res *ReportCSV, err error) {
	if err := c.doer.Do(ctx, "GET", "/reports/{report_id}/download-url", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetReportRequest is given as an input to GetReport.
type GetReportRequest struct {
	// URL path parameters.
	ReportID string `in:"path=report_id" json:"-"`
}

// GetReport performs the GET operation on a Reports resource.
func (c *ReportsClient) GetReport(ctx context.Context, req *GetReportRequest) (res *Report, err error) {
	if err := c.doer.Do(ctx, "GET", "/reports/{report_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
