// Code generated by the Paddle SDK Generator; DO NOT EDIT.
package paddle

import (
	"context"
	paddleerr "github.com/PaddleHQ/paddle-go-sdk/pkg/paddleerr"
)

// ErrProductTaxCategoryNotApproved represents a `product_tax_category_not_approved` error.
// See https://developer.paddle.com/errors/products/product_tax_category_not_approved for more information.
var ErrProductTaxCategoryNotApproved = &paddleerr.Error{
	Code: "product_tax_category_not_approved",
	Type: paddleerr.ErrorTypeRequestError,
}

// ProductWithIncludes: Represents a product entity with included entities.
type ProductWithIncludes struct {
	// ID: Unique Paddle ID for this product, prefixed with `pro_`.
	ID string `json:"id,omitempty"`
	// Name: Name of this product.
	Name string `json:"name,omitempty"`
	// Description: Short description for this product.
	Description *string `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type string `json:"type,omitempty"`
	// TaxCategory: Tax category for this product. Used for charging the correct rate of tax. Selected tax category must be enabled on your Paddle account.
	TaxCategory string `json:"tax_category,omitempty"`
	// ImageURL: Image for this product. Included in the checkout and on some customer documents.
	ImageURL *string `json:"image_url,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status string `json:"status,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Prices: Represents a price entity.
	Prices []Price `json:"prices,omitempty"`
}

// ProductsClient is a client for the Products resource.
type ProductsClient struct {
	doer Doer
}

// ListProductsRequest is given as an input to ListProducts.
type ListProductsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after,omitempty" json:"-"`
	// ID is a query parameter.
	// Return only the IDs specified. Use a comma-separated list to get multiple entities.
	ID []string `in:"query=id,omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `created_at`, `custom_data`, `description`, `id`, `image_url`, `name`, `status`, `tax_category`, and `updated_at`.
	*/
	OrderBy *string `in:"query=order_by,omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `50`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page,omitempty" json:"-"`
	// Status is a query parameter.
	// Return entities that match the specified status. Use a comma-separated list to specify multiple status values.
	Status []string `in:"query=status,omitempty" json:"-"`
	// TaxCategory is a query parameter.
	// Return entities that match the specified tax category. Use a comma-separated list to specify multiple tax categories.
	TaxCategory []string `in:"query=tax_category,omitempty" json:"-"`
	// Type is a query parameter.
	// Return items that match the specified type.
	Type *string `in:"query=type,omitempty" json:"-"`

	// IncludePrices allows requesting the prices sub-resource as part of this request.
	// If set to true, will be included on the response.
	IncludePrices bool `in:"paddle_include=prices" json:"-"`
}

// ListProducts performs the GET operation on a Products resource.
func (c *ProductsClient) ListProducts(ctx context.Context, req *ListProductsRequest) (res *Collection[*ProductWithIncludes], err error) {
	if err := c.doer.Do(ctx, "GET", "/products", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateProductRequest is given as an input to CreateProduct.
type CreateProductRequest struct {
	// Name: Name of this product.
	Name string `json:"name,omitempty"`
	// TaxCategory: Tax category for this product. Used for charging the correct rate of tax. Selected tax category must be enabled on your Paddle account.
	TaxCategory string `json:"tax_category,omitempty"`
	// Description: Short description for this product.
	Description *string `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type *string `json:"type,omitempty"`
	// ImageURL: Image for this product. Included in the checkout and on some customer documents.
	ImageURL *string `json:"image_url,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
}

// CreateProduct performs the POST operation on a Products resource.
func (c *ProductsClient) CreateProduct(ctx context.Context, req *CreateProductRequest) (res *Product, err error) {
	if err := c.doer.Do(ctx, "POST", "/products", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetProductRequest is given as an input to GetProduct.
type GetProductRequest struct {
	// URL path parameters.
	ProductID string `in:"path=product_id" json:"-"`

	// IncludePrices allows requesting the prices sub-resource as part of this request.
	// If set to true, will be included on the response.
	IncludePrices bool `in:"paddle_include=prices" json:"-"`
}

// GetProduct performs the GET operation on a Products resource.
func (c *ProductsClient) GetProduct(ctx context.Context, req *GetProductRequest) (res *ProductWithIncludes, err error) {
	if err := c.doer.Do(ctx, "GET", "/products/{product_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateProductRequest is given as an input to UpdateProduct.
type UpdateProductRequest struct {
	// URL path parameters.
	ProductID string `in:"path=product_id" json:"-"`

	// Name: Name of this product.
	Name *PatchField[string] `json:"name,omitempty"`
	// Description: Short description for this product.
	Description *PatchField[*string] `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type *PatchField[string] `json:"type,omitempty"`
	// TaxCategory: Tax category for this product. Used for charging the correct rate of tax. Selected tax category must be enabled on your Paddle account.
	TaxCategory *PatchField[string] `json:"tax_category,omitempty"`
	// ImageURL: Image for this product. Included in the checkout and on some customer documents.
	ImageURL *PatchField[*string] `json:"image_url,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData *PatchField[CustomData] `json:"custom_data,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status *PatchField[string] `json:"status,omitempty"`
}

// UpdateProduct performs the PATCH operation on a Products resource.
func (c *ProductsClient) UpdateProduct(ctx context.Context, req *UpdateProductRequest) (res *Product, err error) {
	if err := c.doer.Do(ctx, "PATCH", "/products/{product_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}
