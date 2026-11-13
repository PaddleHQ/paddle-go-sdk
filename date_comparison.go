package paddle

// ComparisonOperator compares a date filter against its date.
type ComparisonOperator string

// Comparison operators accepted by date filters. Each filter documents which it supports.
const (
	ComparisonOperatorLT  ComparisonOperator = "LT"
	ComparisonOperatorLTE ComparisonOperator = "LTE"
	ComparisonOperatorGT  ComparisonOperator = "GT"
	ComparisonOperatorGTE ComparisonOperator = "GTE"
)

// DateComparison filters a list by a date, sent as `field[OPERATOR]=date`.
// Leave Operator empty to match the date exactly, where the filter allows it.
type DateComparison struct {
	// Date is an RFC 3339 datetime string, including a timezone.
	Date string
	// Operator compares results against Date. Empty means an exact match.
	Operator ComparisonOperator
}

// QueryOperator returns the operator and date to encode on the query string.
func (d DateComparison) QueryOperator() (operator, value string) {
	return string(d.Operator), d.Date
}
