# Upgrading

All breaking changes will be documented in this file to assist with upgrading.

## v5.0.0

This release introduces breaking changes. Please review the following to upgrade:

1. Minimum Go Version  
   - The minimum supported Go version is now 1.25.  
   - Update your environment to Go 1.25 or later.

2. TrialPeriod Type Change  
   - The `TrialPeriod` property for Price entities has changed type from `Duration` to `TrialPeriod` to support the new `RequiresPaymentMethod` field for cardless trials, see [changelog](https://developer.paddle.com/changelog/2025/cardless-trials-developer-preview?utm_source=dx&utm_medium=paddle-go-sdk).  
   - Update any code that constructs trial periods for prices. Use `TrialPeriod` as a drop-in replacement, set `RequiresPaymentMethod` to take advantage of the new functionality. 

   ```go
   // Before (v4.x)
   TrialPeriod: &paddle.Duration{
       Interval:  paddle.IntervalMonth,
       Frequency: 1,
   }

   // After (v5.0)
   TrialPeriod: &paddle.TrialPeriod{
       Interval:               paddle.IntervalMonth,
       Frequency:              1,
       RequiresPaymentMethod: true, // defaults to true if omitted or set to true to be explicit
   }
   ```

   - The new `TrialPeriod` struct has the same `Interval` and `Frequency` fields as `Duration`, plus the new `RequiresPaymentMethod` boolean.  
   - Affected types: `Price`, `PriceNotification` (paddlenotification), `CreatePriceRequest`, `UpdatePriceRequest`, `TransactionPriceCreateWithProduct`, `TransactionPriceCreateWithProductID`, and related transaction price types.

## v4.0.0

This release introduces several breaking changes and improvements. Please review the following to upgrade:

1. Minimum Go Version  
   - The minimum supported Go version is now 1.23.  
   - Update your environment to Go 1.23 or later.

2. Type Consolidation  
   - `ReportFiltersOperator` and `FilterOperator` types are now consolidated into a single type.  
   - `SubscriptionTransactionDetailsPreview` and `TransactionDetailsPreview` are now a single type.  
   - `SubscriptionTransactionLineItemPreview` and `TransactionLineItemPreview` are now a single type.  
   - Update any references to the removed types to use the new consolidated types.

3. Struct Property Visibility  
   - Structs no longer include properties that are not contextually correct (e.g., writeOnly properties in response entities).  
   - Remove any code that referenced these properties, as they are no longer present.

4. PatchFields for Simulation Updates  
   - `SimulationSingleEventUpdate` and `SimulationScenarioUpdate` now use `PatchFields` for update operations.  
   - Update any code constructing these structs to use patch fields.

5. Transaction Preview Currency Code  
   - `CurrencyCode` is now a pointer type for transaction previews to allow omitting in requests.  
   - Update your code to use pointer values for `CurrencyCode` e.g. within `PreviewTransactionCreate` operations.

6. Client Organisation and Naming  
   - _Applicable if you're using clients directly_
   - `EventTypesClient` naming has been fixed.  
   - `NotificationSettingReplaysClient` has been merged into `NotificationsClient`.  
   - `ListCreditBalances` has moved to `CustomerClient`.  
   - Update any direct usage of these clients to use the new structure and names.

7. Errors  
   - Unused errors have been removed, and new errors have been added.
   - Review handling code to ensure it handles any new errors or doesn't reference removed errors.

## v3.0.0

This release contains 4 commits introducing breaking changes. To summarise: 

1. Rename `DiscountNotification` `Type` property to `DiscountType` to reduce the risk of conflicting type names as `type` is a common property name within the API.

Any existing references to the `Type` property on `DiscountNotification` will need to be updated to reference the property `DiscountType`

2. Support for non catalog items on Transaction Previews with pointer ID properties

- Non catalog items when in a transaction preview can produce a nullable ID, to support these the type are updated to be a pointer ID where applicable.
- We've introduced `ProductPreview` and `TransactionPricePreview` types to support these changes without impacting the core `Product` and `Price` types.
- Any type references for these will have to be updated accordingly and access of the ID properties will need to be dereferenced. 

3. Fix acronym casing for `SubscriptionManagementUrLs`
 
Any existing references to the type `SubscriptionManagementUrLs` will need to be updated to `SubscriptionManagementURLs`.

4. Support nullable discount `starts_at` for subscriptions

- It's possibly for Paddle to return `starts_at` as `null` for Subscription discounts. With this being the case our types have been updated to pointers.
- `SubscriptionDiscountTimePeriod` within `paddle` and `paddlenotification` packages are effected
- The property `StartsAt` is now a pointer and will now need to be dereferenced

## v2.0.0

This release brings 2 breaking changes and fixes that may require some changes in your code to upgrade.

1. `PricePreview` has moved to `PreviewPrices` to match API documentation

Any usage of the method `PricePreview` will need to me refactored to `PreviewPrices` which also includes the request type `PricePreviewRequest` to `PreviewPricesRequest`.

2. Refactored non catalog item types for consistency through the SDK

This release added support for non catalog items on Subscription updates and preview of Subscription updates. With this introduction naming conflicts occurred and were standardised throughout.

Subscription updates now accepts a `[]UpdateSubscriptionItems`instead of `[]SubscriptionUpdateCatalogItem` use `NewUpdateSubscriptionItemsSubscriptionUpdateItem*` functions to create this.

Preview subscription updates now accepts a `[]PreviewSubscriptionUpdateItems` instead of a `[]SubscriptionUpdateCatalogItem` use `NewPreviewSubscriptionUpdateItemsSubscriptionUpdateItem*` functions to create this.

To support these changes you may have to refactor some of your type usage, see the below table for reference:

| Previous Type                                   | New Type                                |
|-------------------------------------------------|-----------------------------------------|
| CatalogItem                                     | TransactionItemFromCatalog              | 
| NonCatalogPriceForAnExistingProduct             | TransactionItemCreateWithPrice          |
| NonCatalogPriceAndProduct                       | TransactionItemCreateWithProduct        |
| TransactionCatalogItem                          | TransactionPreviewItemFromCatalog       |
| TransactionNonCatalogPriceForAnExistingProduct  | TransactionPreviewItemCreateWithPrice   |
| TransactionNonCatalogPriceAndProduct            | TransactionPreviewItemCreateWithProduct |
| SubscriptionUpdateCatalogItem                   | SubscriptionUpdateItemFromCatalog       |
| SubscriptionCatalogItem                         | SubscriptionChargeItemFromCatalog       |
| SubscriptionNonCatalogPriceForAnExistingProduct | SubscriptionChargeItemCreateWithPrice   |
| SubscriptionNonCatalogPriceAndProduct           | SubscriptionChargeItemCreateWithProduct |

Functions names have also changed as part of this standardisation, see the following table and update your code accordingly: 

| Previous Function                                                                  | New Function                                                                |
|------------------------------------------------------------------------------------|-----------------------------------------------------------------------------|
| NewCreateTransactionItemsCatalogItem                                               | NewCreateTransactionItemsTransactionItemFromCatalog                         |
| NewCreateTransactionItemsNonCatalogPriceForAnExistingProduct                       | NewCreateTransactionItemsTransactionItemCreateWithPrice                     |
| NewCreateTransactionItemsNonCatalogPriceAndProduct                                 | NewCreateTransactionItemsTransactionItemCreateWithProduct                   |
| NewUpdateTransactionItemsCatalogItem                                               | NewUpdateTransactionItemsTransactionItemFromCatalog                         |
| NewUpdateTransactionItemsNonCatalogPriceForAnExistingProduct                       | NewUpdateTransactionItemsTransactionItemCreateWithPrice                     |
| NewUpdateTransactionItemsNonCatalogPriceAndProduct                                 | NewUpdateTransactionItemsTransactionItemCreateWithProduct                   |
| NewTransactionPreviewByAddressItemsTransactionCatalogItem                          | NewTransactionPreviewByAddressItemsTransactionPreviewItemFromCatalog        |
| NewTransactionPreviewByAddressItemsTransactionNonCatalogPriceForAnExistingProduct  | NewTransactionPreviewByAddressItemsTransactionPreviewItemCreateWithPrice    |
| NewTransactionPreviewByAddressItemsTransactionNonCatalogPriceAndProduct            | NewTransactionPreviewByAddressItemsTransactionPreviewItemCreateWithProduct  |
| NewTransactionPreviewByIPItemsTransactionCatalogItem                               | NewTransactionPreviewByIPItemsTransactionPreviewItemFromCatalog             |
| NewTransactionPreviewByIPItemsTransactionNonCatalogPriceForAnExistingProduct       | NewTransactionPreviewByIPItemsTransactionPreviewItemCreateWithPrice         |
| NewTransactionPreviewByIPItemsTransactionNonCatalogPriceAndProduct                 | NewTransactionPreviewByIPItemsTransactionPreviewItemCreateWithProduct       |
| NewTransactionPreviewByCustomerItemsTransactionCatalogItem                         | NewTransactionPreviewByCustomerItemsTransactionPreviewItemFromCatalog       |
| NewTransactionPreviewByCustomerItemsTransactionNonCatalogPriceForAnExistingProduct | NewTransactionPreviewByCustomerItemsTransactionPreviewItemCreateWithPrice   |
| NewTransactionPreviewByCustomerItemsTransactionNonCatalogPriceAndProduct           | NewTransactionPreviewByCustomerItemsTransactionPreviewItemCreateWithProduct |

3. Some fields have been moved to pointers to correctly facilitate them being nullable.

- `Type` for Product and Price notifications are nullable `*CatalogType`
- `PaymentMethodID` on Transaction payments is nullable `*PaymentMethodID` for both notifications and API calls

Any usages of `Type` on `PriceNotification` or `ProductNotification` types will need to be changed to handle the `*CatalogType` type.
Any usages of `PaymentMethodID` on `TransactionPaymentAttempt` will need to be changed to handle a `*string` type.

## v1.0.0

- No documented change

## v0.7.0

1. This release brings consistency to naming. We had some occurrences of pluralised prefixes where the pattern is to be singular. 

As these are changing type definitions you will need to refactor the usage of these types accordingly:

| Previous Type                                      | New Type                                          |
|----------------------------------------------------|---------------------------------------------------|
| `BusinessesContacts`                               | `BusinessContacts`                                |
| `NotificationsEvent`                               | `NotificationEvent`                               |
| `GenericNotificationsEvent`                        | `GenericNotificationEvent`                        |
| `ReportsStatus`                                    | `ReportStatus`                                    |
| `ReportsType`                                      | `ReportType`                                      |
| `ReportsFiltersName`                               | `ReportFiltersName`                               |
| `ReportsFiltersOperator`                           | `ReportFiltersOperator`                           |
| `ReportsFilters`                                   | `ReportFilters`                                   |
| `ReportTypeAdjustments`                            | `AdjustmentsReportType`                           |
| `FilterNameAdjustments`                            | `AdjustmentsReportFilterName`                     |
| `ReportFiltersAdjustments`                         | `AdjustmentsReportFilters`                        |
| `ReportTypeTransactions`                           | `TransactionsReportType`                          |
| `FilterNameTransactions`                           | `TransactionsReportFilterName`                    |
| `ReportFiltersTransactions`                        | `TransactionsReportFilters`                       |
| `ReportTypeProductsPrices`                         | `ProductsPricesReportType`                        |
| `FilterNameProductPrices`                          | `ProductPricesReportFilterName`                   |
| `ReportFiltersProductPrices`                       | `ProductPricesReportFilters`                      |
| `ReportTypeDiscounts`                              | `DiscountsReportType`                             |
| `FilterNameDiscounts`                              | `DiscountsReportFilterName`                       |
| `ReportFiltersDiscounts`                           | `DiscountsReportFilters`                          |
| `SubscriptionsTransactionLineItemPreview`          | `SubscriptionTransactionLineItemPreview`          |
| `SubscriptionsTaxRatesUsed`                        | `TaxRatesUsed`                                    |
| `SubscriptionsAdjustmentItem`                      | `SubscriptionsAdjustmentItem`                     |
| `SubscriptionsUpdateCatalogItem`                   | `SubscriptionUpdateCatalogItem`                   |
| `SubscriptionsCatalogItem`                         | `SubscriptionCatalogItem`                         |
| `SubscriptionsNonCatalogPriceForAnExistingProduct` | `SubscriptionNonCatalogPriceForAnExistingProduct` |
| `SubscriptionsNonCatalogPriceAndProduct`           | `SubscriptionNonCatalogPriceAndProduct`           |
| `AdjustmentsTotalsBreakdown`                       | `AdjustmentTotalsBreakdown`                       |
| `AdjustmentsTotals`                                | `TransactionAdjustmentTotals`                     |
| `TransactionsCatalogItem`                          | `TransactionCatalogItem`                          |
| `TransactionsNonCatalogPriceForAnExistingProduct`  | `TransactionNonCatalogPriceForAnExistingProduct`  |
| `TransactionsNonCatalogPriceAndProduct`            | `TransactionNonCatalogPriceAndProduct`            |
| `TransactionsTransactionsCheckout`                 | `TransactionCheckout`                             |


Some types were enumerations using constants and therefore the following constants are renamed too, below is a note of the prefix changes:

| Previous Constant Prefix   | New Constant Prefix             |
|----------------------------|---------------------------------|
| `ReportsStatus`            | `ReportStatus`                  |
| `ReportsType`              | `ReportType`                    |
| `ReportsFiltersName`       | `ReportFiltersName`             |
| `ReportsFiltersOperator`   | `ReportFiltersOperator`         |
| `ReportTypeAdjustments`    | `AdjustmentsReportType`         |
| `FilterNameAdjustments`    | `AdjustmentsReportFilterName`   |
| `ReportTypeTransactions`   | `TransactionsReportType`        |
| `FilterNameTransactions`   | `TransactionsReportFilterName`  |
| `ReportTypeProductsPrices` | `ProductsPricesReportType`      |
| `FilterNameProductPrices`  | `ProductPricesReportFilterName` |
| `ReportTypeDiscounts`      | `DiscountsReportType`           |
| `FilterNameDiscounts`      | `DiscountsReportFilterName`     |


Some types were used as part of a union and therefore the following functions are renamed too:

| Previous Function                                                                     | New Function                                                                         |
|---------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------|
| `NewCreateSubscriptionChargeItemsSubscriptionsCatalogItem`                            | `NewCreateSubscriptionChargeItemsSubscriptionCatalogItem`                            |                              
| `NewCreateSubscriptionChargeItemsSubscriptionsNonCatalogPriceForAnExistingProduct`    | `NewCreateSubscriptionChargeItemsSubscriptionNonCatalogPriceForAnExistingProduct`    |      
| `NewCreateSubscriptionChargeItemsSubscriptionsNonCatalogPriceAndProduct`              | `NewCreateSubscriptionChargeItemsSubscriptionNonCatalogPriceAndProduct`              |                
| `NewPreviewSubscriptionChargeItemsSubscriptionsCatalogItem`                           | `NewPreviewSubscriptionChargeItemsSubscriptionCatalogItem`                           |                             
| `NewPreviewSubscriptionChargeItemsSubscriptionsNonCatalogPriceForAnExistingProduct`   | `NewPreviewSubscriptionChargeItemsSubscriptionNonCatalogPriceForAnExistingProduct`   |     
| `NewPreviewSubscriptionChargeItemsSubscriptionsNonCatalogPriceAndProduct`             | `NewPreviewSubscriptionChargeItemsSubscriptionNonCatalogPriceAndProduct`             |               
| `NewTransactionPreviewByAddressItemsTransactionsCatalogItem`                          | `NewTransactionPreviewByAddressItemsTransactionCatalogItem`                          |                            
| `NewTransactionPreviewByAddressItemsTransactionsNonCatalogPriceForAnExistingProduct`  | `NewTransactionPreviewByAddressItemsTransactionNonCatalogPriceForAnExistingProduct`  |    
| `NewTransactionPreviewByAddressItemsTransactionsNonCatalogPriceAndProduct`            | `NewTransactionPreviewByAddressItemsTransactionNonCatalogPriceAndProduct`            |              
| `NewTransactionPreviewByIPItemsTransactionsCatalogItem`                               | `NewTransactionPreviewByIPItemsTransactionCatalogItem`                               |                                 
| `NewTransactionPreviewByIPItemsTransactionsNonCatalogPriceForAnExistingProduct`       | `NewTransactionPreviewByIPItemsTransactionNonCatalogPriceForAnExistingProduct`       |         
| `NewTransactionPreviewByIPItemsTransactionsNonCatalogPriceAndProduct`                 | `NewTransactionPreviewByIPItemsTransactionNonCatalogPriceAndProduct`                 |                   
| `NewTransactionPreviewByCustomerItemsTransactionsCatalogItem`                         | `NewTransactionPreviewByCustomerItemsTransactionCatalogItem`                         |                           
| `NewTransactionPreviewByCustomerItemsTransactionsNonCatalogPriceForAnExistingProduct` | `NewTransactionPreviewByCustomerItemsTransactionNonCatalogPriceForAnExistingProduct` |
| `NewTransactionPreviewByCustomerItemsTransactionsNonCatalogPriceAndProduct`           | `NewTransactionPreviewByCustomerItemsTransactionNonCatalogPriceAndProduct`           |             


2.The consistency change of pluralised prefixes highlighted conflicting type names for `SubscriptionDiscount` and `Checkout`.

This has resulted in some type changes

| Previous Type           | New Type                            |
|-------------------------|-------------------------------------|
| `SubscriptionDiscount`  | `SubscriptionDiscountTimePeriod`    |
| `SubscriptionsDiscount` | `SubscriptionDiscountEffectiveFrom` |
| `Checkout`              | `TransactionCheckout`               |
| `TransactionsCheckout`  | `TransactionCheckout`               |


## v0.6.0

1. This update makes a significant change to the way the SDK works with the [Reports API](https://developer.paddle.com/api-reference/reports/overview?utm_source=dx&utm_medium=paddle-go-sdk)

When updating to this version if you're integrating with Reports you will need to refactor the code accordingly as it now closely matches the oneOf design of the API. 

An example generating a product/price report: 

```go
res, err := client.CreateReport(ctx,
	paddle.NewCreateReportRequestProductsAndPricesReport(&paddle.ProductsAndPricesReport{
		Type: paddle.ReportTypeProductsPricesProductsPrices,
		Filters: []paddle.ReportFiltersProductPrices{
			paddle.ReportFiltersProductPrices{
				Name:  paddle.FilterNameProductPricesPriceStatus,
				Value: []string{"archived"},
			},
			paddle.ReportFiltersProductPrices{
				Name:  paddle.FilterNameProductPricesProductStatus,
				Value: []string{"archived"},
			},
		},
	}),
)
```

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
