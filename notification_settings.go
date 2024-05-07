// Code generated by the Paddle SDK Generator; DO NOT EDIT.
package paddle

import "context"

// NotificationSettingType: Where notifications should be sent for this destination..
type NotificationSettingType string

const (
	NotificationSettingTypeEmail = "email"
	NotificationSettingTypeURL   = "url"
)

// NotificationSetting: Represents a notification destination.
type NotificationSetting struct {
	// ID: Unique Paddle ID for this notification setting, prefixed with `ntfset_`.
	ID string `json:"id,omitempty"`
	// Description: Short description for this notification destination. Shown in the Paddle web app.
	Description string `json:"description,omitempty"`
	// Type: Where notifications should be sent for this destination.
	Type string `json:"type,omitempty"`
	// Destination: Webhook endpoint URL or email address.
	Destination string `json:"destination,omitempty"`
	// Active: Whether Paddle should try to deliver events to this notification destination.
	Active bool `json:"active,omitempty"`
	// APIVersion: API version that returned objects for events should conform to. Must be a valid version of the Paddle API. Cannot be a version older than your account default. Defaults to your account default if not included.
	APIVersion int `json:"api_version,omitempty"`
	// IncludeSensitiveFields: Whether potentially sensitive fields should be sent to this notification destination.
	IncludeSensitiveFields bool `json:"include_sensitive_fields,omitempty"`
	// SubscribedEvents: Represents an event type.
	SubscribedEvents []EventType `json:"subscribed_events,omitempty"`
	// EndpointSecretKey: Webhook destination secret key, prefixed with `pdl_ntfset_`. Used for signature verification.
	EndpointSecretKey string `json:"endpoint_secret_key,omitempty"`
}

// NotificationSettingsClient is a client for the Notification settings resource.
type NotificationSettingsClient struct {
	doer Doer
}

// ListNotificationSettingsRequest is given as an input to ListNotificationSettings.
type ListNotificationSettingsRequest struct {
	// After is a query parameter.
	// Return entities after the specified Paddle ID when working with paginated endpoints. Used in the `meta.pagination.next` URL in responses for list operations.
	After *string `in:"query=after,omitempty" json:"-"`
	// PerPage is a query parameter.
	/*
	   Set how many entities are returned per page. Paddle returns the maximum number of results if a number greater than the maximum is requested. Check `meta.pagination.per_page` in the response to see how many were returned.

	   Default: `200`; Maximum: `200`.
	*/
	PerPage *int `in:"query=per_page,omitempty" json:"-"`
	// OrderBy is a query parameter.
	/*
	   Order returned entities by the specified field and direction (`[ASC]` or `[DESC]`). For example, `?order_by=id[ASC]`.

	   Valid fields for ordering: `id`.
	*/
	OrderBy *string `in:"query=order_by,omitempty" json:"-"`
}

// ListNotificationSettings performs the GET operation on a Notification settings resource.
func (c *NotificationSettingsClient) ListNotificationSettings(ctx context.Context, req *ListNotificationSettingsRequest) (res *Collection[*NotificationSetting], err error) {
	if err := c.doer.Do(ctx, "GET", "/notification-settings", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateNotificationSettingRequest is given as an input to CreateNotificationSetting.
type CreateNotificationSettingRequest struct {
	// Description: Short description for this notification destination. Shown in the Paddle Dashboard.
	Description string `json:"description,omitempty"`
	// Destination: Webhook endpoint URL or email address.
	Destination string `json:"destination,omitempty"`
	// SubscribedEvents: Subscribed events for this notification destination. When creating or updating a notification destination, pass an array of event type names only. Paddle returns the complete event type object.
	SubscribedEvents []Event `json:"subscribed_events,omitempty"`
	// Type: Where notifications should be sent for this destination.
	Type string `json:"type,omitempty"`
	// APIVersion: API version that returned objects for events should conform to. Must be a valid version of the Paddle API. Cannot be a version older than your account default. Defaults to your account default if not included.
	APIVersion *int `json:"api_version,omitempty"`
	// IncludeSensitiveFields: Whether potentially sensitive fields should be sent to this notification destination.
	IncludeSensitiveFields *bool `json:"include_sensitive_fields,omitempty"`
}

// CreateNotificationSetting performs the POST operation on a Notification settings resource.
func (c *NotificationSettingsClient) CreateNotificationSetting(ctx context.Context, req *CreateNotificationSettingRequest) (res *NotificationSetting, err error) {
	if err := c.doer.Do(ctx, "POST", "/notification-settings", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetNotificationSettingRequest is given as an input to GetNotificationSetting.
type GetNotificationSettingRequest struct {
	// URL path parameters.
	NotificationSettingID string `in:"path=notification_setting_id" json:"-"`
}

// GetNotificationSetting performs the GET operation on a Notification settings resource.
func (c *NotificationSettingsClient) GetNotificationSetting(ctx context.Context, req *GetNotificationSettingRequest) (res *NotificationSetting, err error) {
	if err := c.doer.Do(ctx, "GET", "/notification-settings/{notification_setting_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// UpdateNotificationSettingRequest is given as an input to UpdateNotificationSetting.
type UpdateNotificationSettingRequest struct {
	// URL path parameters.
	NotificationSettingID string `in:"path=notification_setting_id" json:"-"`

	// Description: Short description for this notification destination. Shown in the Paddle Dashboard.
	Description *PatchField[string] `json:"description,omitempty"`
	// Destination: Webhook endpoint URL or email address.
	Destination *PatchField[string] `json:"destination,omitempty"`
	// Active: Whether Paddle should try to deliver events to this notification destination.
	Active *PatchField[bool] `json:"active,omitempty"`
	// APIVersion: API version that returned objects for events should conform to. Must be a valid version of the Paddle API. Cannot be a version older than your account default. Defaults to your account default if omitted.
	APIVersion *PatchField[int] `json:"api_version,omitempty"`
	// IncludeSensitiveFields: Whether potentially sensitive fields should be sent to this notification destination.
	IncludeSensitiveFields *PatchField[bool] `json:"include_sensitive_fields,omitempty"`
	// SubscribedEvents: Subscribed events for this notification destination. When creating or updating a notification destination, pass an array of event type names only. Paddle returns the complete event type object.
	SubscribedEvents *PatchField[[]Event] `json:"subscribed_events,omitempty"`
}

// UpdateNotificationSetting performs the PATCH operation on a Notification settings resource.
func (c *NotificationSettingsClient) UpdateNotificationSetting(ctx context.Context, req *UpdateNotificationSettingRequest) (res *NotificationSetting, err error) {
	if err := c.doer.Do(ctx, "PATCH", "/notification-settings/{notification_setting_id}", req, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// DeleteNotificationSettingRequest is given as an input to DeleteNotificationSetting.
type DeleteNotificationSettingRequest struct {
	// URL path parameters.
	NotificationSettingID string `in:"path=notification_setting_id" json:"-"`
}

// DeleteNotificationSetting performs the DELETE operation on a Notification settings resource.
func (c *NotificationSettingsClient) DeleteNotificationSetting(ctx context.Context, req *DeleteNotificationSettingRequest) (err error) {
	if err := c.doer.Do(ctx, "DELETE", "/notification-settings/{notification_setting_id}", req, nil); err != nil {
		return err
	}

	return nil
}
