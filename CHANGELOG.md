# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Check our main [developer changelog](https://developer.paddle.com/?utm_source=dx&utm_medium=paddle-go-sdk) for information about changes to the Paddle Billing platform, the Paddle API, and other developer tools.

## 0.5.0 - 2024-07-26

### Breaking changes

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

### Breaking changes

- Shared types `TransactionSubscriptionPriceCreateWithProduct` and `TransactionSubscriptionPriceCreateWithProductID` have diverged into 2 new types for Transactions and Subscriptions
- Include entity types are unified with their non-include variant, e.g. `ProductWithIncludes` and `Product` are now just `Product`

### Added

- `AUD` and `CAD` chargeback currencies are available

### Fixed

- Subscription resume properly accepts the subscription id as a path parameter on `ResumeSubscriptionRequest`
- Correctly make a `PATCH` call for `PreviewSubscription` instead of `POST` and use `PatchField` types in the `PreviewSubscriptionRequest` request

## 0.3.0 - 2024-05-27

### Breaking changes

- `Event` and `NotificationsEvent` are now an interfaces.
- `NotificationsEvent` has moved package to `paddlenotification`.

### Added

- New `paddlenotification` package.
- Support for `/events` and `/notifications` endpoints.


## 0.1.0 - 2024-05-07

### Added

- Initial early access release. Added support for the most frequently used Paddle Billing entities and API operations. Check the [README](./README.md) for more information.
