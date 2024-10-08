// Code generated by the Paddle SDK Generator; DO NOT EDIT.

package paddlenotification

// EventTypeName: Type of event sent by Paddle, in the format `entity.event_type`..
type EventTypeName string

const (
	EventTypeNameAddressCreated           EventTypeName = "address.created"
	EventTypeNameAddressImported          EventTypeName = "address.imported"
	EventTypeNameAddressUpdated           EventTypeName = "address.updated"
	EventTypeNameAdjustmentCreated        EventTypeName = "adjustment.created"
	EventTypeNameAdjustmentUpdated        EventTypeName = "adjustment.updated"
	EventTypeNameBusinessCreated          EventTypeName = "business.created"
	EventTypeNameBusinessImported         EventTypeName = "business.imported"
	EventTypeNameBusinessUpdated          EventTypeName = "business.updated"
	EventTypeNameCustomerCreated          EventTypeName = "customer.created"
	EventTypeNameCustomerImported         EventTypeName = "customer.imported"
	EventTypeNameCustomerUpdated          EventTypeName = "customer.updated"
	EventTypeNameDiscountCreated          EventTypeName = "discount.created"
	EventTypeNameDiscountImported         EventTypeName = "discount.imported"
	EventTypeNameDiscountUpdated          EventTypeName = "discount.updated"
	EventTypeNamePayoutCreated            EventTypeName = "payout.created"
	EventTypeNamePayoutPaid               EventTypeName = "payout.paid"
	EventTypeNamePriceCreated             EventTypeName = "price.created"
	EventTypeNamePriceImported            EventTypeName = "price.imported"
	EventTypeNamePriceUpdated             EventTypeName = "price.updated"
	EventTypeNameProductCreated           EventTypeName = "product.created"
	EventTypeNameProductImported          EventTypeName = "product.imported"
	EventTypeNameProductUpdated           EventTypeName = "product.updated"
	EventTypeNameReportCreated            EventTypeName = "report.created"
	EventTypeNameReportUpdated            EventTypeName = "report.updated"
	EventTypeNameSubscriptionActivated    EventTypeName = "subscription.activated"
	EventTypeNameSubscriptionCanceled     EventTypeName = "subscription.canceled"
	EventTypeNameSubscriptionCreated      EventTypeName = "subscription.created"
	EventTypeNameSubscriptionImported     EventTypeName = "subscription.imported"
	EventTypeNameSubscriptionPastDue      EventTypeName = "subscription.past_due"
	EventTypeNameSubscriptionPaused       EventTypeName = "subscription.paused"
	EventTypeNameSubscriptionResumed      EventTypeName = "subscription.resumed"
	EventTypeNameSubscriptionTrialing     EventTypeName = "subscription.trialing"
	EventTypeNameSubscriptionUpdated      EventTypeName = "subscription.updated"
	EventTypeNameTransactionBilled        EventTypeName = "transaction.billed"
	EventTypeNameTransactionCanceled      EventTypeName = "transaction.canceled"
	EventTypeNameTransactionCompleted     EventTypeName = "transaction.completed"
	EventTypeNameTransactionCreated       EventTypeName = "transaction.created"
	EventTypeNameTransactionPaid          EventTypeName = "transaction.paid"
	EventTypeNameTransactionPastDue       EventTypeName = "transaction.past_due"
	EventTypeNameTransactionPaymentFailed EventTypeName = "transaction.payment_failed"
	EventTypeNameTransactionReady         EventTypeName = "transaction.ready"
	EventTypeNameTransactionUpdated       EventTypeName = "transaction.updated"
)

// CountryCode: Supported two-letter ISO 3166-1 alpha-2 country code for this address..
type CountryCode string

const (
	CountryCodeAD CountryCode = "AD"
	CountryCodeAE CountryCode = "AE"
	CountryCodeAG CountryCode = "AG"
	CountryCodeAI CountryCode = "AI"
	CountryCodeAL CountryCode = "AL"
	CountryCodeAM CountryCode = "AM"
	CountryCodeAO CountryCode = "AO"
	CountryCodeAR CountryCode = "AR"
	CountryCodeAS CountryCode = "AS"
	CountryCodeAT CountryCode = "AT"
	CountryCodeAU CountryCode = "AU"
	CountryCodeAW CountryCode = "AW"
	CountryCodeAX CountryCode = "AX"
	CountryCodeAZ CountryCode = "AZ"
	CountryCodeBA CountryCode = "BA"
	CountryCodeBB CountryCode = "BB"
	CountryCodeBD CountryCode = "BD"
	CountryCodeBE CountryCode = "BE"
	CountryCodeBF CountryCode = "BF"
	CountryCodeBG CountryCode = "BG"
	CountryCodeBH CountryCode = "BH"
	CountryCodeBI CountryCode = "BI"
	CountryCodeBJ CountryCode = "BJ"
	CountryCodeBL CountryCode = "BL"
	CountryCodeBM CountryCode = "BM"
	CountryCodeBN CountryCode = "BN"
	CountryCodeBO CountryCode = "BO"
	CountryCodeBQ CountryCode = "BQ"
	CountryCodeBR CountryCode = "BR"
	CountryCodeBS CountryCode = "BS"
	CountryCodeBT CountryCode = "BT"
	CountryCodeBV CountryCode = "BV"
	CountryCodeBW CountryCode = "BW"
	CountryCodeBZ CountryCode = "BZ"
	CountryCodeCA CountryCode = "CA"
	CountryCodeCC CountryCode = "CC"
	CountryCodeCG CountryCode = "CG"
	CountryCodeCH CountryCode = "CH"
	CountryCodeCI CountryCode = "CI"
	CountryCodeCK CountryCode = "CK"
	CountryCodeCL CountryCode = "CL"
	CountryCodeCM CountryCode = "CM"
	CountryCodeCN CountryCode = "CN"
	CountryCodeCO CountryCode = "CO"
	CountryCodeCR CountryCode = "CR"
	CountryCodeCV CountryCode = "CV"
	CountryCodeCW CountryCode = "CW"
	CountryCodeCX CountryCode = "CX"
	CountryCodeCY CountryCode = "CY"
	CountryCodeCZ CountryCode = "CZ"
	CountryCodeDE CountryCode = "DE"
	CountryCodeDJ CountryCode = "DJ"
	CountryCodeDK CountryCode = "DK"
	CountryCodeDM CountryCode = "DM"
	CountryCodeDO CountryCode = "DO"
	CountryCodeDZ CountryCode = "DZ"
	CountryCodeEC CountryCode = "EC"
	CountryCodeEE CountryCode = "EE"
	CountryCodeEG CountryCode = "EG"
	CountryCodeEH CountryCode = "EH"
	CountryCodeER CountryCode = "ER"
	CountryCodeES CountryCode = "ES"
	CountryCodeET CountryCode = "ET"
	CountryCodeFI CountryCode = "FI"
	CountryCodeFJ CountryCode = "FJ"
	CountryCodeFK CountryCode = "FK"
	CountryCodeFM CountryCode = "FM"
	CountryCodeFO CountryCode = "FO"
	CountryCodeFR CountryCode = "FR"
	CountryCodeGA CountryCode = "GA"
	CountryCodeGB CountryCode = "GB"
	CountryCodeGD CountryCode = "GD"
	CountryCodeGE CountryCode = "GE"
	CountryCodeGF CountryCode = "GF"
	CountryCodeGG CountryCode = "GG"
	CountryCodeGH CountryCode = "GH"
	CountryCodeGI CountryCode = "GI"
	CountryCodeGL CountryCode = "GL"
	CountryCodeGM CountryCode = "GM"
	CountryCodeGN CountryCode = "GN"
	CountryCodeGP CountryCode = "GP"
	CountryCodeGQ CountryCode = "GQ"
	CountryCodeGR CountryCode = "GR"
	CountryCodeGS CountryCode = "GS"
	CountryCodeGT CountryCode = "GT"
	CountryCodeGU CountryCode = "GU"
	CountryCodeGW CountryCode = "GW"
	CountryCodeGY CountryCode = "GY"
	CountryCodeHK CountryCode = "HK"
	CountryCodeHM CountryCode = "HM"
	CountryCodeHN CountryCode = "HN"
	CountryCodeHR CountryCode = "HR"
	CountryCodeHU CountryCode = "HU"
	CountryCodeID CountryCode = "ID"
	CountryCodeIE CountryCode = "IE"
	CountryCodeIL CountryCode = "IL"
	CountryCodeIM CountryCode = "IM"
	CountryCodeIN CountryCode = "IN"
	CountryCodeIO CountryCode = "IO"
	CountryCodeIQ CountryCode = "IQ"
	CountryCodeIS CountryCode = "IS"
	CountryCodeIT CountryCode = "IT"
	CountryCodeJE CountryCode = "JE"
	CountryCodeJM CountryCode = "JM"
	CountryCodeJO CountryCode = "JO"
	CountryCodeJP CountryCode = "JP"
	CountryCodeKE CountryCode = "KE"
	CountryCodeKG CountryCode = "KG"
	CountryCodeKH CountryCode = "KH"
	CountryCodeKI CountryCode = "KI"
	CountryCodeKM CountryCode = "KM"
	CountryCodeKN CountryCode = "KN"
	CountryCodeKR CountryCode = "KR"
	CountryCodeKW CountryCode = "KW"
	CountryCodeKY CountryCode = "KY"
	CountryCodeKZ CountryCode = "KZ"
	CountryCodeLA CountryCode = "LA"
	CountryCodeLB CountryCode = "LB"
	CountryCodeLC CountryCode = "LC"
	CountryCodeLI CountryCode = "LI"
	CountryCodeLK CountryCode = "LK"
	CountryCodeLR CountryCode = "LR"
	CountryCodeLS CountryCode = "LS"
	CountryCodeLT CountryCode = "LT"
	CountryCodeLU CountryCode = "LU"
	CountryCodeLV CountryCode = "LV"
	CountryCodeMA CountryCode = "MA"
	CountryCodeMC CountryCode = "MC"
	CountryCodeMD CountryCode = "MD"
	CountryCodeME CountryCode = "ME"
	CountryCodeMF CountryCode = "MF"
	CountryCodeMG CountryCode = "MG"
	CountryCodeMH CountryCode = "MH"
	CountryCodeMK CountryCode = "MK"
	CountryCodeMN CountryCode = "MN"
	CountryCodeMO CountryCode = "MO"
	CountryCodeMP CountryCode = "MP"
	CountryCodeMQ CountryCode = "MQ"
	CountryCodeMR CountryCode = "MR"
	CountryCodeMS CountryCode = "MS"
	CountryCodeMT CountryCode = "MT"
	CountryCodeMU CountryCode = "MU"
	CountryCodeMV CountryCode = "MV"
	CountryCodeMW CountryCode = "MW"
	CountryCodeMX CountryCode = "MX"
	CountryCodeMY CountryCode = "MY"
	CountryCodeMZ CountryCode = "MZ"
	CountryCodeNA CountryCode = "NA"
	CountryCodeNC CountryCode = "NC"
	CountryCodeNE CountryCode = "NE"
	CountryCodeNF CountryCode = "NF"
	CountryCodeNG CountryCode = "NG"
	CountryCodeNL CountryCode = "NL"
	CountryCodeNO CountryCode = "NO"
	CountryCodeNP CountryCode = "NP"
	CountryCodeNR CountryCode = "NR"
	CountryCodeNU CountryCode = "NU"
	CountryCodeNZ CountryCode = "NZ"
	CountryCodeOM CountryCode = "OM"
	CountryCodePA CountryCode = "PA"
	CountryCodePE CountryCode = "PE"
	CountryCodePF CountryCode = "PF"
	CountryCodePG CountryCode = "PG"
	CountryCodePH CountryCode = "PH"
	CountryCodePK CountryCode = "PK"
	CountryCodePL CountryCode = "PL"
	CountryCodePM CountryCode = "PM"
	CountryCodePN CountryCode = "PN"
	CountryCodePR CountryCode = "PR"
	CountryCodePS CountryCode = "PS"
	CountryCodePT CountryCode = "PT"
	CountryCodePW CountryCode = "PW"
	CountryCodePY CountryCode = "PY"
	CountryCodeQA CountryCode = "QA"
	CountryCodeRE CountryCode = "RE"
	CountryCodeRO CountryCode = "RO"
	CountryCodeRS CountryCode = "RS"
	CountryCodeRW CountryCode = "RW"
	CountryCodeSA CountryCode = "SA"
	CountryCodeSB CountryCode = "SB"
	CountryCodeSC CountryCode = "SC"
	CountryCodeSE CountryCode = "SE"
	CountryCodeSG CountryCode = "SG"
	CountryCodeSH CountryCode = "SH"
	CountryCodeSI CountryCode = "SI"
	CountryCodeSJ CountryCode = "SJ"
	CountryCodeSK CountryCode = "SK"
	CountryCodeSL CountryCode = "SL"
	CountryCodeSM CountryCode = "SM"
	CountryCodeSN CountryCode = "SN"
	CountryCodeSR CountryCode = "SR"
	CountryCodeST CountryCode = "ST"
	CountryCodeSV CountryCode = "SV"
	CountryCodeSX CountryCode = "SX"
	CountryCodeSZ CountryCode = "SZ"
	CountryCodeTC CountryCode = "TC"
	CountryCodeTD CountryCode = "TD"
	CountryCodeTF CountryCode = "TF"
	CountryCodeTG CountryCode = "TG"
	CountryCodeTH CountryCode = "TH"
	CountryCodeTJ CountryCode = "TJ"
	CountryCodeTK CountryCode = "TK"
	CountryCodeTL CountryCode = "TL"
	CountryCodeTM CountryCode = "TM"
	CountryCodeTN CountryCode = "TN"
	CountryCodeTO CountryCode = "TO"
	CountryCodeTR CountryCode = "TR"
	CountryCodeTT CountryCode = "TT"
	CountryCodeTV CountryCode = "TV"
	CountryCodeTW CountryCode = "TW"
	CountryCodeTZ CountryCode = "TZ"
	CountryCodeUA CountryCode = "UA"
	CountryCodeUG CountryCode = "UG"
	CountryCodeUM CountryCode = "UM"
	CountryCodeUS CountryCode = "US"
	CountryCodeUY CountryCode = "UY"
	CountryCodeUZ CountryCode = "UZ"
	CountryCodeVA CountryCode = "VA"
	CountryCodeVC CountryCode = "VC"
	CountryCodeVG CountryCode = "VG"
	CountryCodeVI CountryCode = "VI"
	CountryCodeVN CountryCode = "VN"
	CountryCodeVU CountryCode = "VU"
	CountryCodeWF CountryCode = "WF"
	CountryCodeWS CountryCode = "WS"
	CountryCodeXK CountryCode = "XK"
	CountryCodeYT CountryCode = "YT"
	CountryCodeZA CountryCode = "ZA"
	CountryCodeZM CountryCode = "ZM"
)

// CustomData: Your own structured key-value data.
type CustomData map[string]any

// Status: Whether this entity can be used in Paddle..
type Status string

const (
	StatusActive   Status = "active"
	StatusArchived Status = "archived"
)

// ImportMeta: Import information for this entity. `null` if this entity is not imported.
type ImportMeta struct {
	// ExternalID: Reference or identifier for this entity from the solution where it was imported from.
	ExternalID *string `json:"external_id,omitempty"`
	// ImportedFrom: Name of the platform where this entity was imported from.
	ImportedFrom string `json:"imported_from,omitempty"`
}

// CurrencyCode: Three-letter ISO 4217 currency code for this adjustment. Set automatically by Paddle based on the `currency_code` of the related transaction..
type CurrencyCode string

const (
	CurrencyCodeUSD CurrencyCode = "USD"
	CurrencyCodeEUR CurrencyCode = "EUR"
	CurrencyCodeGBP CurrencyCode = "GBP"
	CurrencyCodeJPY CurrencyCode = "JPY"
	CurrencyCodeAUD CurrencyCode = "AUD"
	CurrencyCodeCAD CurrencyCode = "CAD"
	CurrencyCodeCHF CurrencyCode = "CHF"
	CurrencyCodeHKD CurrencyCode = "HKD"
	CurrencyCodeSGD CurrencyCode = "SGD"
	CurrencyCodeSEK CurrencyCode = "SEK"
	CurrencyCodeARS CurrencyCode = "ARS"
	CurrencyCodeBRL CurrencyCode = "BRL"
	CurrencyCodeCNY CurrencyCode = "CNY"
	CurrencyCodeCOP CurrencyCode = "COP"
	CurrencyCodeCZK CurrencyCode = "CZK"
	CurrencyCodeDKK CurrencyCode = "DKK"
	CurrencyCodeHUF CurrencyCode = "HUF"
	CurrencyCodeILS CurrencyCode = "ILS"
	CurrencyCodeINR CurrencyCode = "INR"
	CurrencyCodeKRW CurrencyCode = "KRW"
	CurrencyCodeMXN CurrencyCode = "MXN"
	CurrencyCodeNOK CurrencyCode = "NOK"
	CurrencyCodeNZD CurrencyCode = "NZD"
	CurrencyCodePLN CurrencyCode = "PLN"
	CurrencyCodeRUB CurrencyCode = "RUB"
	CurrencyCodeTHB CurrencyCode = "THB"
	CurrencyCodeTRY CurrencyCode = "TRY"
	CurrencyCodeTWD CurrencyCode = "TWD"
	CurrencyCodeUAH CurrencyCode = "UAH"
	CurrencyCodeZAR CurrencyCode = "ZAR"
)

// TimePeriod: Billing period that proration is based on.
type TimePeriod struct {
	// StartsAt: RFC 3339 datetime string of when this period starts.
	StartsAt string `json:"starts_at,omitempty"`
	// EndsAt: RFC 3339 datetime string of when this period ends.
	EndsAt string `json:"ends_at,omitempty"`
}

// Proration: How proration was calculated for this adjustment item.
type Proration struct {
	// Rate: Rate used to calculate proration.
	Rate string `json:"rate,omitempty"`
	// BillingPeriod: Billing period that proration is based on.
	BillingPeriod TimePeriod `json:"billing_period,omitempty"`
}

// CurrencyCodeChargebacks: Three-letter ISO 4217 currency code for the original chargeback fee..
type CurrencyCodeChargebacks string

const (
	CurrencyCodeChargebacksAUD CurrencyCodeChargebacks = "AUD"
	CurrencyCodeChargebacksCAD CurrencyCodeChargebacks = "CAD"
	CurrencyCodeChargebacksEUR CurrencyCodeChargebacks = "EUR"
	CurrencyCodeChargebacksGBP CurrencyCodeChargebacks = "GBP"
	CurrencyCodeChargebacksUSD CurrencyCodeChargebacks = "USD"
)

// Original: Chargeback fee before conversion to the payout currency. `null` when the chargeback fee is the same as the payout currency.
type Original struct {
	// Amount: Fee amount for this chargeback in the original currency.
	Amount string `json:"amount,omitempty"`
	// CurrencyCode: Three-letter ISO 4217 currency code for the original chargeback fee.
	CurrencyCode CurrencyCodeChargebacks `json:"currency_code,omitempty"`
}

// ChargebackFee: Chargeback fees incurred for this adjustment. Only returned when the adjustment `action` is `chargeback` or `chargeback_warning`.
type ChargebackFee struct {
	// Amount: Chargeback fee converted into the payout currency.
	Amount string `json:"amount,omitempty"`
	// Original: Chargeback fee before conversion to the payout currency. `null` when the chargeback fee is the same as the payout currency.
	Original *Original `json:"original,omitempty"`
}

// CurrencyCodePayouts: Three-letter ISO 4217 currency code used for the payout for this transaction. If your primary currency has changed, this reflects the primary currency at the time the transaction was billed..
type CurrencyCodePayouts string

const (
	CurrencyCodePayoutsAUD CurrencyCodePayouts = "AUD"
	CurrencyCodePayoutsCAD CurrencyCodePayouts = "CAD"
	CurrencyCodePayoutsCHF CurrencyCodePayouts = "CHF"
	CurrencyCodePayoutsCNY CurrencyCodePayouts = "CNY"
	CurrencyCodePayoutsCZK CurrencyCodePayouts = "CZK"
	CurrencyCodePayoutsDKK CurrencyCodePayouts = "DKK"
	CurrencyCodePayoutsEUR CurrencyCodePayouts = "EUR"
	CurrencyCodePayoutsGBP CurrencyCodePayouts = "GBP"
	CurrencyCodePayoutsHUF CurrencyCodePayouts = "HUF"
	CurrencyCodePayoutsPLN CurrencyCodePayouts = "PLN"
	CurrencyCodePayoutsSEK CurrencyCodePayouts = "SEK"
	CurrencyCodePayoutsUSD CurrencyCodePayouts = "USD"
	CurrencyCodePayoutsZAR CurrencyCodePayouts = "ZAR"
)

// CatalogType: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app..
type CatalogType string

const (
	CatalogTypeCustom   CatalogType = "custom"
	CatalogTypeStandard CatalogType = "standard"
)

// Interval: Unit of time..
type Interval string

const (
	IntervalDay   Interval = "day"
	IntervalWeek  Interval = "week"
	IntervalMonth Interval = "month"
	IntervalYear  Interval = "year"
)

// Duration: How often this price should be charged. `null` if price is non-recurring (one-time).
type Duration struct {
	// Interval: Unit of time.
	Interval Interval `json:"interval,omitempty"`
	// Frequency: Amount of time.
	Frequency int `json:"frequency,omitempty"`
}

// TaxMode: How tax is calculated for this price..
type TaxMode string

const (
	TaxModeAccountSetting TaxMode = "account_setting"
	TaxModeExternal       TaxMode = "external"
	TaxModeInternal       TaxMode = "internal"
)

// Money: Base price. This price applies to all customers, except for customers located in countries where you have `unit_price_overrides`.
type Money struct {
	// Amount: Amount in the lowest denomination for the currency, e.g. 10 USD = 1000 (cents).
	Amount string `json:"amount,omitempty"`
	// CurrencyCode: Supported three-letter ISO 4217 currency code.
	CurrencyCode CurrencyCode `json:"currency_code,omitempty"`
}

// UnitPriceOverride: List of unit price overrides. Use to override the base price with a custom price and currency for a country or group of countries.
type UnitPriceOverride struct {
	// CountryCodes: Supported two-letter ISO 3166-1 alpha-2 country code.
	CountryCodes []CountryCode `json:"country_codes,omitempty"`
	// UnitPrice: Override price. This price applies to customers located in the countries for this unit price override.
	UnitPrice Money `json:"unit_price,omitempty"`
}

// PriceQuantity: Limits on how many times the related product can be purchased at this price. Useful for discount campaigns.
type PriceQuantity struct {
	// Minimum: Minimum quantity of the product related to this price that can be bought. Required if `maximum` set.
	Minimum int `json:"minimum,omitempty"`
	// Maximum: Maximum quantity of the product related to this price that can be bought. Required if `minimum` set. Must be greater than or equal to the `minimum` value.
	Maximum int `json:"maximum,omitempty"`
}

// TaxCategory: Tax category for this product. Used for charging the correct rate of tax. Selected tax category must be enabled on your Paddle account..
type TaxCategory string

const (
	TaxCategoryDigitalGoods                TaxCategory = "digital-goods"
	TaxCategoryEbooks                      TaxCategory = "ebooks"
	TaxCategoryImplementationServices      TaxCategory = "implementation-services"
	TaxCategoryProfessionalServices        TaxCategory = "professional-services"
	TaxCategorySaas                        TaxCategory = "saas"
	TaxCategorySoftwareProgrammingServices TaxCategory = "software-programming-services"
	TaxCategoryStandard                    TaxCategory = "standard"
	TaxCategoryTrainingServices            TaxCategory = "training-services"
	TaxCategoryWebsiteHosting              TaxCategory = "website-hosting"
)

// CollectionMode: How payment is collected for transactions created for this subscription. `automatic` for checkout, `manual` for invoices..
type CollectionMode string

const (
	CollectionModeAutomatic CollectionMode = "automatic"
	CollectionModeManual    CollectionMode = "manual"
)

// BillingDetails: Details for invoicing. Required if `collection_mode` is `manual`.
type BillingDetails struct {
	// EnableCheckout: Whether the related transaction may be paid using a Paddle Checkout. If omitted when creating a transaction, defaults to `false`.
	EnableCheckout bool `json:"enable_checkout,omitempty"`
	// PurchaseOrderNumber: Customer purchase order number. Appears on invoice documents.
	PurchaseOrderNumber string `json:"purchase_order_number,omitempty"`
	// AdditionalInformation: Notes or other information to include on this invoice. Appears on invoice documents.
	AdditionalInformation *string `json:"additional_information,omitempty"`
	// PaymentTerms: How long a customer has to pay this invoice once issued.
	PaymentTerms Duration `json:"payment_terms,omitempty"`
}

// Price: Related price entity for this item. This reflects the price entity at the time it was added to the subscription.
type Price struct {
	// ID: Unique Paddle ID for this price, prefixed with `pri_`.
	ID string `json:"id,omitempty"`
	// ProductID: Paddle ID for the product that this price is for, prefixed with `pro_`.
	ProductID string `json:"product_id,omitempty"`
	// Description: Internal description for this price, not shown to customers. Typically notes for your team.
	Description string `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type CatalogType `json:"type,omitempty"`
	// Name: Name of this price, shown to customers at checkout and on invoices. Typically describes how often the related product bills.
	Name *string `json:"name,omitempty"`
	// BillingCycle: How often this price should be charged. `null` if price is non-recurring (one-time).
	BillingCycle *Duration `json:"billing_cycle,omitempty"`
	// TrialPeriod: Trial period for the product related to this price. The billing cycle begins once the trial period is over. `null` for no trial period. Requires `billing_cycle`.
	TrialPeriod *Duration `json:"trial_period,omitempty"`
	// TaxMode: How tax is calculated for this price.
	TaxMode TaxMode `json:"tax_mode,omitempty"`
	// UnitPrice: Base price. This price applies to all customers, except for customers located in countries where you have `unit_price_overrides`.
	UnitPrice Money `json:"unit_price,omitempty"`
	// UnitPriceOverrides: List of unit price overrides. Use to override the base price with a custom price and currency for a country or group of countries.
	UnitPriceOverrides []UnitPriceOverride `json:"unit_price_overrides,omitempty"`
	// Quantity: Limits on how many times the related product can be purchased at this price. Useful for discount campaigns.
	Quantity PriceQuantity `json:"quantity,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status Status `json:"status,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Product: Related product entity for this item. This reflects the product entity at the time it was added to the subscription.
type Product struct {
	// ID: Unique Paddle ID for this product, prefixed with `pro_`.
	ID string `json:"id,omitempty"`
	// Name: Name of this product.
	Name string `json:"name,omitempty"`
	// Description: Short description for this product.
	Description *string `json:"description,omitempty"`
	// Type: Type of item. Standard items are considered part of your catalog and are shown on the Paddle web app.
	Type CatalogType `json:"type,omitempty"`
	// TaxCategory: Tax category for this product. Used for charging the correct rate of tax. Selected tax category must be enabled on your Paddle account.
	TaxCategory TaxCategory `json:"tax_category,omitempty"`
	// ImageURL: Image for this product. Included in the checkout and on some customer documents.
	ImageURL *string `json:"image_url,omitempty"`
	// CustomData: Your own structured key-value data.
	CustomData CustomData `json:"custom_data,omitempty"`
	// Status: Whether this entity can be used in Paddle.
	Status Status `json:"status,omitempty"`
	// ImportMeta: Import information for this entity. `null` if this entity is not imported.
	ImportMeta *ImportMeta `json:"import_meta,omitempty"`
	// CreatedAt: RFC 3339 datetime string of when this entity was created. Set automatically by Paddle.
	CreatedAt string `json:"created_at,omitempty"`
	// UpdatedAt: RFC 3339 datetime string of when this entity was updated. Set automatically by Paddle.
	UpdatedAt string `json:"updated_at,omitempty"`
}
