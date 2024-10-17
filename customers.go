// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"

	paddleerr "github.com/PaddleHQ/paddle-go-sdk/v2/pkg/paddleerr"
)

// ErrCustomerAlreadyExists represents a `customer_already_exists` error.
// See https://developer.paddle.com/errors/customers/customer_already_exists for more information.
var ErrCustomerAlreadyExists = &paddleerr.Error{
	Code: "customer_already_exists",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrCustomerEmailDomainNotAllowed represents a `customer_email_domain_not_allowed` error.
// See https://developer.paddle.com/errors/customers/customer_email_domain_not_allowed for more information.
var ErrCustomerEmailDomainNotAllowed = &paddleerr.Error{
	Code: "customer_email_domain_not_allowed",
	Type: paddleerr.ErrorTypeRequestError,
}

// ErrCustomerEmailInvalid represents a `customer_email_invalid` error.
// See https://developer.paddle.com/errors/customers/customer_email_invalid for more information.
var ErrCustomerEmailInvalid = &paddleerr.Error{
	Code: "customer_email_invalid",
	Type: paddleerr.ErrorTypeRequestError,
}

// Customer: Represents a customer entity with included entities.
type Customer struct {
	// ID: Unique Paddle ID for this customer entity, prefixed with `ctm_`.
	ID string `json:"id,omitempty"`
	// Name: Full name of this customer. Required when creating transactions where `collection_mode` is `manual` (invoices).
	Name *string `json:"name,omitempty"`
	// Email: Email address for this customer.
	Email string `json:"email,omitempty"`
	/*
	   MarketingConsent: Whether this customer opted into marketing from you. `false` unless customers check the marketing consent box
	   when using Paddle Checkout. Set automatically by Paddle.
	*/
	MarketingConsent bool `json:"marketing_consent,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status Status `json:"status,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// Locale: Valid IETF BCP 47 short form locale tag. If omitted, defaults to `en`.
	Locale string `json:"locale,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
}

// CustomersClient is a client for the Customers resource.
type CustomersClient struct {
	doer Doer
}

// ListCustomersRequest is given as an input to ListCustomers.
type ListCustomersRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
	// Email is a query parameter.
	// Return entities that exactly match the specified email address. Use a comma-separated list to specify multiple email addresses. Recommended for precise matching of email addresses.
	Email []string `in:"query=email;omitempty" json:"-"`
	// ID is a query parameter.
	// Return only the IDs specified. Use a comma-separated list to get multiple entities.
	ID []string `in:"query=id;omitempty" json:"-"`
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
	// Search is a query parameter.
	// Return entities that match a search query. Searches `id`, `name`, and `email` fields. Use the `email` query parameter for precise matching of email addresses.
	Search *string `in:"query=search;omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status;omitempty" json:"-"`
}

// ListCustomers performs the GET operation on a Customers resource.
func (c *CustomersClient) ListCustomers(ctx context.Context, req *ListCustomersRequest) (res *Collection[*Customer], err error) {
	if err := c.doer.Do(ctx, "GET", "/customers", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateCustomerRequest is given as an input to CreateCustomer.
type CreateCustomerRequest struct {
	// Email: Email address for this customer.
	Email string `json:"email,omitempty"`
	// Name: Full name of this customer. Required when creating transactions where `collection_mode` is `manual` (invoices).
	Name *string `json:"name,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// Locale: Valid IETF BCP 47 short form locale tag. If omitted, defaults to `en`.
	Locale *string `json:"locale,omitempty"`
}

// CreateCustomer performs the POST operation on a Customers resource.
func (c *CustomersClient) CreateCustomer(ctx context.Context, req *CreateCustomerRequest) (res *Customer, err error) {
	if err := c.doer.Do(ctx, "POST", "/customers", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetCustomerRequest is given as an input to GetCustomer.
type GetCustomerRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`
}

// GetCustomer performs the GET operation on a Customers resource.
func (c *CustomersClient) GetCustomer(ctx context.Context, req *GetCustomerRequest) (res *Customer, err error) {
	if err := c.doer.Do(ctx, "GET", "/customers/{customer_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateCustomerRequest is given as an input to UpdateCustomer.
type UpdateCustomerRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`

	// Name: Full name of this customer. Required when creating transactions where `collection_mode` is `manual` (invoices).
	Name *PatchField[*string] `json:"name,omitempty"`
	// Email: Email address for this customer.
	Email *PatchField[string] `json:"email,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status *PatchField[Status] `json:"status,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData *PatchField[CustomData] `json:"custom_data,omitempty"`
	// Locale: Valid IETF BCP 47 short form locale tag.
	Locale *PatchField[string] `json:"locale,omitempty"`
}

// UpdateCustomer performs the PATCH operation on a Customers resource.
func (c *CustomersClient) UpdateCustomer(ctx context.Context, req *UpdateCustomerRequest) (res *Customer, err error) {
	if err := c.doer.Do(ctx, "PATCH", "/customers/{customer_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
