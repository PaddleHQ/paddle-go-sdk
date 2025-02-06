// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddle

// SDK contains all sub-clients for the Paddle API.
type SDK struct {
	*ProductsClient
	*PricesClient
	*TransactionsClient
	*PricingPreviewClient
	*AdjustmentsClient
	*CustomersClient
	*AddressesClient
	*BusinessesClient
	*PaymentMethodsClient
	*CustomerPortalSessionsClient
	*NotificationSettingsClient
	*EventTypesClient
	*EventsClient
	*NotificationsClient
	*NotificationLogsClient
	*SimulationTypesClient
	*SimulationsClient
	*SimulationRunsClient
	*SimulationRunEventsClient
	*IPAddressesClient
	*DiscountsClient
	*SubscriptionsClient
	*ReportsClient
}

// newSDK creates a new SDK instance. This is auto-generated, modifications should be done in the generator.
func newSDK(d Doer) *SDK {
	return &SDK{
		AddressesClient:              &AddressesClient{doer: d},
		AdjustmentsClient:            &AdjustmentsClient{doer: d},
		BusinessesClient:             &BusinessesClient{doer: d},
		CustomerPortalSessionsClient: &CustomerPortalSessionsClient{doer: d},
		CustomersClient:              &CustomersClient{doer: d},
		DiscountsClient:              &DiscountsClient{doer: d},
		EventTypesClient:             &EventTypesClient{doer: d},
		EventsClient:                 &EventsClient{doer: d},
		IPAddressesClient:            &IPAddressesClient{doer: d},
		NotificationLogsClient:       &NotificationLogsClient{doer: d},
		NotificationSettingsClient:   &NotificationSettingsClient{doer: d},
		NotificationsClient:          &NotificationsClient{doer: d},
		PaymentMethodsClient:         &PaymentMethodsClient{doer: d},
		PricesClient:                 &PricesClient{doer: d},
		PricingPreviewClient:         &PricingPreviewClient{doer: d},
		ProductsClient:               &ProductsClient{doer: d},
		ReportsClient:                &ReportsClient{doer: d},
		SimulationRunEventsClient:    &SimulationRunEventsClient{doer: d},
		SimulationRunsClient:         &SimulationRunsClient{doer: d},
		SimulationTypesClient:        &SimulationTypesClient{doer: d},
		SimulationsClient:            &SimulationsClient{doer: d},
		SubscriptionsClient:          &SubscriptionsClient{doer: d},
		TransactionsClient:           &TransactionsClient{doer: d},
	}
}
