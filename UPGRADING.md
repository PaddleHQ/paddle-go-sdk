# Upgrading

All breaking changes prior to v1 will be documented in this file to assist with upgrading.

## v0.5.0

1. This update fixes correctly resolves generated enum types. Prior to this enum types in our spec were using there respective basic type `string`

To upgrade to `0.5.0` from any earlier version you will need to ensure you're correctly utilising the provided constants for these types, e.g.

```go
// Instead of this
priceUpdate := &paddle.UpdatePriceRequest{Status: paddle.NewPatchField("archived")}

// paddle.Status can be passed using paddle.StatusArchived
priceUpdate := &paddle.UpdatePriceRequest{Status: paddle.NewPatchField(paddle.StatusArchived)}
```

2. Notifications and Events `EventType` property now use `EventTypeName` instead of `string` type. 

Any functions this property is passed into will need to be updated accordingly.

3. `Totals` is now a unified type from the previous 4 types `Totals`, `TotalsFormatted`, `UnitTotals` and `UnitTotalsFormatted`

The type shape is the same however receiving functions will need to be updated accordingly.

Any usage of the following types may cause breakages.

| Package              | Type                                      |
|----------------------|-------------------------------------------|
| `paddlenotification` | `TransactionLineItem`                     |
| `paddle`             | `PricePreviewLineItem`                    |
| `paddle`             | `TransactionLineItem`                     |
| `paddle`             | `SubscriptionsTransactionLineItemPreview` |
| `paddle`             | `TransactionLineItemPreview`              |

4. Various types have been renamed to be less generic and lower the risk of a future breaking change

| Previous Type                      | New Type                            |
|------------------------------------|-------------------------------------|
| `Action`                           | `AdjustmentAction`                  |
| `IPAddressesData`                  | `IPAddress`                         | 
| `Origin`                           | `NotificationOrigin`                |
| `SubscriptionsCatalogItem`         | `SubscriptionsUpdateCatalogItem`    | 
| `ResultAction`                     | `UpdateSummaryResultAction`         |
| `Result`                           | `UpdateSummaryResult`               |
| `PreviewSubscriptionUpdateSummary` | `SubscriptionPreviewUpdateSummary`  |
| `Breakdown`                        | `AdjustmentsTotalsBreakdown`        |
| `PreviewSubscriptionRequest`       | `PreviewSubscriptionUpdateRequest`  |
| `PreviewTransactionRequest`        | `PreviewTransactionCreateRequest`   |
| `CountryAndZipPostalCode`          | `TransactionPreviewByAddress`       |
| `CountryAndZipPostalCodeItems`     | `TransactionPreviewByAddressItems`  |
| `IPAddress`                        | `TransactionPreviewByIP`            |
| `IPAddressItems`                   | `TransactionPreviewByIPItems`       |
| `ExistingCustomerPaddleIDs`        | `TransactionPreviewByCustomer`      |
| `ExistingCustomerPaddleIDsItems`   | `TransactionPreviewByCustomerItems` |

Alongside these type changes some functions have been respectively updated and any usages will need to be changed in your codebase.

| Previous Function Name                                                             | New Function Name                                                                     |
|------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| `NewCountryAndZipPostalCodeItemsTransactionsCatalogItem`                           | `NewTransactionPreviewByAddressItemsTransactionsCatalogItem`                          | 
| `NewCountryAndZipPostalCodeItemsTransactionsNonCatalogPriceForAnExistingProduct`   | `NewTransactionPreviewByAddressItemsTransactionsNonCatalogPriceForAnExistingProduct`  |
| `NewCountryAndZipPostalCodeItemsTransactionsNonCatalogPriceAndProduct`             | `NewTransactionPreviewByAddressItemsTransactionsNonCatalogPriceAndProduct`            |
| `NewIPAddressItemsTransactionsCatalogItem`                                         | `NewTransactionPreviewByIPItemsTransactionsCatalogItem`                               |
| `NewIPAddressItemsTransactionsNonCatalogPriceForAnExistingProduct`                 | `NewTransactionPreviewByIPItemsTransactionsNonCatalogPriceForAnExistingProduct`       |
| `NewIPAddressItemsTransactionsNonCatalogPriceAndProduct`                           | `NewTransactionPreviewByIPItemsTransactionsNonCatalogPriceAndProduct`                 |
| `NewExistingCustomerPaddleIDsItemsTransactionsCatalogItem`                         | `NewTransactionPreviewByCustomerItemsTransactionsCatalogItem`                         |
| `NewExistingCustomerPaddleIDsItemsTransactionsNonCatalogPriceForAnExistingProduct` | `NewTransactionPreviewByCustomerItemsTransactionsNonCatalogPriceForAnExistingProduct` |
| `NewExistingCustomerPaddleIDsItemsTransactionsNonCatalogPriceAndProduct`           | `NewTransactionPreviewByCustomerItemsTransactionsNonCatalogPriceAndProduct`           |
| `NewCreateSubscriptionChargeItemsSubscriptionsSubscriptionsCatalogItem`            | `NewCreateSubscriptionChargeItemsSubscriptionsCatalogItem`                            |
| `NewPreviewSubscriptionChargeItemsSubscriptionsSubscriptionsCatalogItem`           | `NewPreviewSubscriptionChargeItemsSubscriptionsCatalogItem`                           |
| `PreviewSubscription`                                                              | `PreviewSubscriptionUpdate`                                                           |
| `PreviewTransaction`                                                               | `PreviewTransactionCreate`                                                            |


## v0.4.0

1. This version includes a breaking change to shared types between Transactions and Subscriptions resources which have now been split into their own respective types.   

To upgrade to `0.4.0` from any earlier version you will need to refactor your use of these types as noted below depending on the resource:

| Original Type                                     | Transaction Type                      | Subscription Type                     | 
|---------------------------------------------------|---------------------------------------|---------------------------------------|
| `TransactionSubscriptionPriceCreateWithProduct`   | `TransactionPriceCreateWithProduct`   | `SubscriptionChargeCreateWithProduct` |
| `TransactionSubscriptionPriceCreateWithProductID` | `TransactionPriceCreateWithProductID` | `SubscriptionChargeCreateWithPrice`   |

2. We have unified types that have includes with their non-include variant, e.g. `ProductWithIncludes` and `Product` are now just `Product`

To upgrade to `0.4.0` any references to the types below need to be refactored accordingly

| Previous Type          | New Type            |
|------------------------|---------------------|
| `ProductWithIncludes`  | `Product`           |
| `CustomerIncludes`     | `Customer`          | 
| `TransactionIncludes`  | `Transaction`       |
| `SubscriptionIncludes` | `Subscription`      | 
| `PriceIncludes`        | `Price`             |
