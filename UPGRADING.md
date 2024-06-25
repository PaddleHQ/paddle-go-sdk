# Upgrading

All breaking changes prior to v1 will be documented in this file to assist with upgrading.

## v0.4.0

This version includes a breaking change to shared types between Transactions and Subscriptions resources which have now been split into their own respective types.   

To upgrade to `0.4.0` from any earlier version you will need to refactor your use of these types as noted below depending on the resource:

| Original Type                                     | Transaction Type                      | Subscription Type                     | 
|---------------------------------------------------|---------------------------------------|---------------------------------------|
| `TransactionSubscriptionPriceCreateWithProduct`   | `TransactionPriceCreateWithProduct`   | `SubscriptionChargeCreateWithProduct` |
| `TransactionSubscriptionPriceCreateWithProductID` | `TransactionPriceCreateWithProductID` | `SubscriptionChargeCreateWithPrice`   |


