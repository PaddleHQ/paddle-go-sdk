// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

import (
	"context"
	paddleerr "github.com/PaddleHQ/paddle-go-sdk/pkg/paddleerr"
)

// ErrAddressLocationNotAllowed represents a `address_location_not_allowed` error.
// See https://developer.paddle.com/errors/addresses/address_location_not_allowed for more information.
var ErrAddressLocationNotAllowed = &paddleerr.Error{
	Code: "address_location_not_allowed",
	Type: paddleerr.ErrorTypeRequestError,
}

// AddressesClient is a client for the Addresses resource.
type AddressesClient struct {
	doer Doer
}

// ListAddressesRequest is given as an input to ListAddresses.
type ListAddressesRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`

	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after;omitempty" json:"-"`
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
	// Return entities that match a search query. Searches all fields except `status`, `created_at`, and `updated_at`.
	Search *string `in:"query=search;omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status;omitempty" json:"-"`
}

// ListAddresses performs the GET operation on a Addresses resource.
func (c *AddressesClient) ListAddresses(ctx context.Context, req *ListAddressesRequest) (res *Collection[*Address], err error) {
	if err := c.doer.Do(ctx, "GET", "/customers/{customer_id}/addresses", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateAddressRequest is given as an input to CreateAddress.
type CreateAddressRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`

	// CountryCode: Supported two-letter ISO 3166-1 alpha-2 country code for this address.
	CountryCode string `json:"country_code,omitempty"`
	// Description: Memorable description for this address.
	Description *string `json:"description,omitempty"`
	// FirstLine: First line of this address.
	FirstLine *string `json:"first_line,omitempty"`
	// SecondLine: Second line of this address.
	SecondLine *string `json:"second_line,omitempty"`
	// City: City of this address.
	City *string `json:"city,omitempty"`
	// PostalCode: ZIP or postal code of this address. Required for some countries.
	PostalCode *string `json:"postal_code,omitempty"`
	// Region: State, county, or region of this address.
	Region *string `json:"region,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
}

// CreateAddress performs the POST operation on a Addresses resource.
func (c *AddressesClient) CreateAddress(ctx context.Context, req *CreateAddressRequest) (res *Address, err error) {
	if err := c.doer.Do(ctx, "POST", "/customers/{customer_id}/addresses", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetAddressRequest is given as an input to GetAddress.
type GetAddressRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`
	AddressID  string `in:"path=address_id" json:"-"`
}

// GetAddress performs the GET operation on a Addresses resource.
func (c *AddressesClient) GetAddress(ctx context.Context, req *GetAddressRequest) (res *Address, err error) {
	if err := c.doer.Do(ctx, "GET", "/customers/{customer_id}/addresses/{address_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateAddressRequest is given as an input to UpdateAddress.
type UpdateAddressRequest struct {
	// URL path parameters.
	CustomerID string `in:"path=customer_id" json:"-"`
	AddressID  string `in:"path=address_id" json:"-"`

	// Description: Memorable description for this address.
	Description *PatchField[*string] `json:"description,omitempty"`
	// FirstLine: First line of this address.
	FirstLine *PatchField[*string] `json:"first_line,omitempty"`
	// SecondLine: Second line of this address.
	SecondLine *PatchField[*string] `json:"second_line,omitempty"`
	// City: City of this address.
	City *PatchField[*string] `json:"city,omitempty"`
	// PostalCode: ZIP or postal code of this address. Required for some countries.
	PostalCode *PatchField[*string] `json:"postal_code,omitempty"`
	// Region: State, county, or region of this address.
	Region *PatchField[*string] `json:"region,omitempty"`
	// CountryCode: Supported two-letter ISO 3166-1 alpha-2 country code for this address.
	CountryCode *PatchField[string] `json:"country_code,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData *PatchField[CustomData] `json:"custom_data,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status *PatchField[string] `json:"status,omitempty"`
}

// UpdateAddress performs the PATCH operation on a Addresses resource.
func (c *AddressesClient) UpdateAddress(ctx context.Context, req *UpdateAddressRequest) (res *Address, err error) {
	if err := c.doer.Do(ctx, "PATCH", "/customers/{customer_id}/addresses/{address_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
