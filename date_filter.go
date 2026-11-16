package paddle

import "iter"

// DateFilter filters a list by a date. Set Equals for an exact match, or set
// one or more operators, such as GTE and LT for a range. Each date is an
// RFC 3339 datetime string, including a timezone.
type DateFilter struct {
	// Equals matches the date exactly, sent as `field=date`.
	Equals *string
	// LT matches dates before this one, sent as `field[LT]=date`.
	LT *string
	// LTE matches dates on or before this one, sent as `field[LTE]=date`.
	LTE *string
	// GT matches dates after this one, sent as `field[GT]=date`.
	GT *string
	// GTE matches dates on or after this one, sent as `field[GTE]=date`.
	GTE *string
}

// QueryOperators yields each operator and date to encode on the query string.
func (f DateFilter) QueryOperators() iter.Seq2[string, string] {
	return queryOperators(
		queryOperator{"", f.Equals},
		queryOperator{"LT", f.LT},
		queryOperator{"LTE", f.LTE},
		queryOperator{"GT", f.GT},
		queryOperator{"GTE", f.GTE},
	)
}

// DateRange filters a list to dates within inclusive bounds. Set one bound or
// both. Each date is an RFC 3339 datetime string, including a timezone.
type DateRange struct {
	// GTE matches dates on or after this one, sent as `field[GTE]=date`.
	GTE *string
	// LTE matches dates on or before this one, sent as `field[LTE]=date`.
	LTE *string
}

// QueryOperators yields each operator and date to encode on the query string.
func (r DateRange) QueryOperators() iter.Seq2[string, string] {
	return queryOperators(
		queryOperator{"GTE", r.GTE},
		queryOperator{"LTE", r.LTE},
	)
}

type queryOperator struct {
	operator string
	value    *string
}

// queryOperators yields the set operators in order, skipping unset or empty values.
func queryOperators(operators ...queryOperator) iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, o := range operators {
			if o.value == nil || *o.value == "" {
				continue
			}

			if !yield(o.operator, *o.value) {
				return
			}
		}
	}
}
