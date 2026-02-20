# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Check our main [developer changelog](https://developer.paddle.com/?utm_source=dx&utm_medium=paddle-go-sdk) for information about changes to the Paddle Billing platform, the Paddle API, and other developer tools.

## 5.0.1 - 2026-02-20

### Fixed

- Resolved module path version mismatch which prevented the installation of v5

## 5.0.0 - 2026-02-11

### Added

- Parse `Retry-After` header on API errors for rate limiting and retry scenarios. Exposes `RetryAfter` on `paddleerr.Error` with `TotalDelay()`, `WaitTime()`, and `IsExpired()` methods
- Filter subscriptions by `next_billed_at` when listing (e.g. to identify cardless trials), see [changelog](https://developer.paddle.com/changelog/2025/cardless-trials-developer-preview?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for payout reconciliation reports deprecating balance reports, see [changelog](https://developer.paddle.com/changelog/2025/payout-reconciliation-report?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for `remittance_reference` on payout notifications, see [changelog](https://developer.paddle.com/changelog/2025/payout-reconciliation-report?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for new `location` tax mode to automatically present prices as inclusive or exclusive of tax based on customer location, see [changelog](https://developer.paddle.com/changelog/2025/automatic-tax-inclusive-exclusive-prices?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for `grand_total_tax` on transaction totals to see the tax amount charged after credits are applied, see [changelog](https://developer.paddle.com/changelog/2026/grand-total-tax-field?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for `requires_payment_method` on price trial period for cardless trials, see [changelog](https://developer.paddle.com/changelog/2025/cardless-trials-developer-preview?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for new Korean payment method types (KakaoPay, NaverPay, Samsung Pay, Payco, Korean local cards), see [changelog](https://developer.paddle.com/changelog/2025/improved-korean-payment-methods?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for API key exposures and `api_key_exposure.created` notifications, see [changelog](https://developer.paddle.com/changelog/2025/secret-scanning?utm_source=dx&utm_medium=paddle-go-sdk)

### Changed

- Minimum supported Go version is now 1.25 see [UPGRADING](./UPGRADING.md) for details
- `TrialPeriod` property for Price entities has changed type from `Duration` to `TrialPeriod` to support `RequiresPaymentMethod`, see [UPGRADING](./UPGRADING.md) for details

### Fixed

- Added missing `discount_group.updated` event type name constant, see [changelog](https://developer.paddle.com/changelog/2025/discount-groups-new-api-operations?utm_source=dx&utm_medium=paddle-go-sdk)
- Transaction update example in documentation

## 4.2.0 - 2025-09-23

### Added

- Error support for adjustment_payout_earnings_cannot_be_zero, see [docs](https://developer.paddle.com/errors/adjustments/adjustment_payout_earnings_cannot_be_zero?utm_source=dx&utm_medium=paddle-go-sdk)
- Handling new enum value of PaymentMethodOrigin for explicit consent to saved payments, see [changelog](https://developer.paddle.com/changelog/2025/spm-consent-subscriptions?utm_source=dx&utm_medium=paddle-go-sdk)
- New payment methods, BLIK, MB Way, Pix and UPI, see [changelog](https://developer.paddle.com/changelog/2025/blik-mbway-payment-methods?utm_source=dx&utm_medium=paddle-go-sdk)
- Non-catalog discounts on Transactions, see [changelog](https://developer.paddle.com/changelog/2025/custom-discounts?utm_source=dx&utm_medium=paddle-go-sdk)
- Support `retained_fee` field on totals objects to show the fees retained by Paddle for the adjustment

## 4.1.0 - 2025-08-01

### Added

- Add timestamp tolerance checks to `WebhookVerifier` to be safe against replay attacks
- Support for ClientTokens API ([changelog](https://developer.paddle.com/changelog/2025/client-side-token-api?utm_source=dx&utm_medium=paddle-go-sdk))
- Add event type filtering on `ListEvents` ([changelog](https://developer.paddle.com/changelog/2025/filter-events-by-type?utm_source=dx&utm_medium=paddle-go-sdk))
- Updates to DiscountGroups API ([changelog](https://developer.paddle.com/changelog/2025/discount-groups-updates?utm_source=dx&utm_medium=paddle-go-sdk))
  - `GetDiscountGroup` and `UpdateDiscountGroup` API support
  - `discount_group.updated` event and notification
- Support for `exchange_rate` and `fee_rate` fields on payout totals
- Support for new API errors

## 4.0.0 - 2025-07-02

### Added

- Support for Balance Reports ([changelog](https://developer.paddle.com/changelog/2025/balance-reports?utm_source=dx&utm_medium=paddle-go-sdk))
- Support for Simulation Scenario configuration ([changelog](https://developer.paddle.com/changelog/2025/webhook-simulator-scenario-configuration?utm_source=dx&utm_medium=paddle-go-sdk))
- Support for Adjustment `tax_mode` ([changelog](https://developer.paddle.com/changelog/2025/tax-exclusive-refunds?utm_source=dx&utm_medium=paddle-go-sdk))
- Support for Discount Groups ([changelog](https://developer.paddle.com/changelog/2025/discount-groups?utm_source=dx&utm_medium=paddle-go-sdk))
- Support for local Korea payment methods ([changelog](https://developer.paddle.com/changelog/2024/korean-payment-methods?utm_source=dx&utm_medium=paddle-go-sdk))
- Support for Discount mode for checkout recovery ([changelog](https://developer.paddle.com/changelog/2025/checkout-recovery?utm_source=dx&utm_medium=paddle-go-sdk#change-fields))
- Support for AdjustmentAction `chargeback_warning_reverse`
- Support for API Key Events and Notifications

### Changed

- Minimum supported Go version is now 1.23
- Consolidated `ReportFiltersOperator` and `FilterOperator` types
- Consolidated `SubscriptionTransactionDetailsPreview` and `TransactionDetailsPreview` into a single type
- Consolidated `SubscriptionTransactionLineItemPreview` and `TransactionLineItemPreview` into a single type
- `CurrencyCode` is now a pointer type for Transaction Previews
- Improved property visibility: structs no longer include properties not contextually correct
- Client organisation and naming cleanup:
  - `EventTypesClient` naming fixed
  - `NotificationSettingReplaysClient` merged into `NotificationsClient`
  - `ListCreditBalances` moved to `CustomerClient` (breaking for direct client usage)
- Various types moved into `shared` package, unused errors removed, new errors added

### Fixed

- `SimulationSingleEventUpdate` and `SimulationScenarioUpdate` now correctly use `PatchFields` for updates
- Added missing event type name constants: `payment_method.saved`, `payment_method.deleted`, `transaction.revised`
- Simulation unmarshal of payment_method events
- Miscellaneous formatting and spec description updates for consistency


## 3.1.1 - 2025-05-12

### Fixed

- Resolved bug preventing custom client being used via `paddle.WithClient` method

## 3.1.0 - 2025-02-07

### Added

- New API errors
- Support for calling Transaction Preview API without Address information
- Support for the Transaction Revise API, see related [changelog](https://developer.paddle.com/changelog/2024/revise-transaction-customer-information?utm_source=dx&utm_medium=paddle-go-sdk)

## 3.0.0 - 2024-12-20

### Fixed

- Limited bytes that can be read via the webhook verification middleware
- Handle nullable `SubscriptionDiscountTimePeriod` property `StartsAt` making it a pointer

### Added

- Support for new API errors
- Vietnamese Dong (VND) currency support, see related [changelog](https://developer.paddle.com/changelog/2024/vietnamese-dong-vnd-supported-currency?utm_source=dx&utm_medium=paddle-go-sdk)
- Simpler creation of full adjustments and a new `AdjustmentType` property, see related [changelog](https://developer.paddle.com/changelog/2024/refund-credit-full-total?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for controlling how Paddle bills when resuming subscriptions with `SubscriptionOnResume`, see related [chagnelog](https://developer.paddle.com/changelog/2024/resume-subscription-billing-period-options?utm_source=dx&utm_medium=paddle-go-sdk)
- Support for non catalog items on Transaction Previews with pointer ID properties

### Changed

- Removed write only property `items` from PricePreview entity
- Removed unused DiscountStatus values `expired` and `used`
- Corrected URL acronym casing for `SubscriptionManagementURLs`

## 2.1.3 - 2024-11-06

### Fixed

- Corrected internal package version number
- Added missing changelog updates

## 2.1.2 - 2024-11-06

### Added

- An example for webhook unmarshalling was added to the package.

## 2.1.1 - 2024-10-18

### Fixed

- Create v2 module for the latest major version

## 2.1.0 - 2024-10-17

### Added

- Support for the Simulation API - see related [changelog entry](https://developer.paddle.com/changelog/2024/webhook-simulator?utm_source=dx&utm_medium=paddle-go-sdk)
- Filtering of customer credit balances by `currency_code` in `AdjustmentsClient.ListCreditBalances`
- New API error mappings

### Fixed

- Subscription resume union funcs now accept the required path param
  - `NewResumeSubscriptionRequestResumeOnASpecificDate`
  - `NewResumeSubscriptionRequestResumeImmediately`
- Response unmarshal for GetNotifications to the Notification struct

## 2.0.0 - 2024-09-18

### Changed

- `PricePreview` has moved to `PreviewPrices` to match API documentation
- Refactored non catalog item types for consistency through the SDK, see [UPGRADING](./UPGRADING.md) for details

### Added

- Support for new Paddle API errors
- Adjustment credit notes (`GetAdjustmentCreditNote`)
- `disposition` query parameter for `GetTransactionInvoice` and `GetAdjustmentCreditNote`
- `active` query parameter filter support for NotificationSettings 
- Subscription update and preview update support non catalog items

### Fixed

- `Type` for Product and Price notifications are nullable `*CatalogType`
- `PaymentMethodID` on Transaction payments is nullable `*PaymentMethodID` for both notifications and API calls

## 1.0.0 - 2024-08-15

No documented changes.

## 0.7.0 - 2024-08-09

### Changed

- Consistency to naming where some occurrences of pluralised prefixes where the pattern is to be singular see [UPGRADING](./UPGRADING.md) for details
- Post consistency fixes to some type names that conflicted, see [UPGRADING](./UPGRADING.md) for details

## 0.6.0 - 2024-08-07

### Changed

- Reports architecture has changed to a oneOf design

### Fixed

- Report filter values incorrectly typed as a `string` when it requires either `string` or `[]string` depending on the filter name. 

## 0.5.0 - 2024-07-26

### Changed

- Paddle API enum types are now correctly enforced where applicable
- Notifications and Events `EventType` property now use `EventTypeName` instead of `string` type
- `Totals` is now a unified type from the previous 4 types `Totals`, `TotalsFormatted`, `UnitTotals` and `UnitTotalsFormatted`
- Various types have been renamed to be less generic and lower the risk of a future breaking change, see [UPGRADING](./UPGRADING.md) for details

### Added

- `Product` has been added as a property of `SubscriptionItem`
- `TaxRatesUsed` has been added as a property of `Adjustments` and `AdjustmentNotification`
- `Notification` property `Type` narrows to `EventTypeName` from `string`
- All notifications and events now have `EventType` narrowed to `EventTypeName` from `string`

### Fixed

- Handling collections for API endpoints without pagination
- `UnitPriceOverride` property `CountryCodes` is now typed as a slice
- `Transaction` property `AvailablePaymentMethods` is now typed as a slice
- `PricePreview` property `AvailablePaymentMethods` is now typed as a slice
- `CreateNotificationSettingRequest` property `SubscribedEvents` is now correctly typed
- `UpdateNotificationSettingRequest` property `SubscribedEvents` is now correctly typed

## 0.4.0 - 2024-06-26

### Changed

- Shared types `TransactionSubscriptionPriceCreateWithProduct` and `TransactionSubscriptionPriceCreateWithProductID` have diverged into 2 new types for Transactions and Subscriptions
- Include entity types are unified with their non-include variant, e.g. `ProductWithIncludes` and `Product` are now just `Product`

### Added

- `AUD` and `CAD` chargeback currencies are available

### Fixed

- Subscription resume properly accepts the subscription id as a path parameter on `ResumeSubscriptionRequest`
- Correctly make a `PATCH` call for `PreviewSubscription` instead of `POST` and use `PatchField` types in the `PreviewSubscriptionRequest` request

## 0.3.0 - 2024-05-27

### Changed

- `Event` and `NotificationsEvent` are now an interfaces.
- `NotificationsEvent` has moved package to `paddlenotification`.

### Added

- New `paddlenotification` package.
- Support for `/events` and `/notifications` endpoints.


## 0.1.0 - 2024-05-07

### Added

- Initial early access release. Added support for the most frequently used Paddle Billing entities and API operations. Check the [README](./README.md) for more information.
