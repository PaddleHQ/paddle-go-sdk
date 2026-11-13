package client

import (
	"errors"
	"iter"
	"reflect"

	"github.com/ggicci/httpin/core"
)

// Operand is implemented by query values that carry operators, such as
// paddle.DateFilter. QueryOperators yields each operator and value in the
// order they're encoded; an empty operator is an exact match.
type Operand interface {
	QueryOperators() iter.Seq2[string, string]
}

// DirectiveOperator is a directive that adds a query parameter keyed by the
// directive's argument plus the value's operator.
type DirectiveOperator struct{}

// Decode is not implemented. This is because we only deal with encoding
// concerns in the client.
func (*DirectiveOperator) Decode(_ *core.DirectiveRuntime) error {
	//nolint:err113 // typed error not needed here.
	return errors.New("not implemented")
}

// ErrDirectiveRequiresOperand is returned when the directive's struct field
// does not implement Operand.
var ErrDirectiveRequiresOperand = errors.New("operator directive requires an Operand value")

// Encode sets a query parameter `name[OPERATOR]=value` for each operator the
// value yields, or `name=value` for an empty operator. A nil value sets
// nothing, like omitempty.
// e.g.
//
//	type Request struct {
//	 BilledAt *paddle.DateFilter `in:"paddle_operator=billed_at"`
//	}
//
//	Req{BilledAt: &paddle.DateFilter{GTE: paddle.PtrTo("2024-01-01T00:00:00Z"), LT: paddle.PtrTo("2024-02-01T00:00:00Z")}}
//	// ?billed_at[GTE]=2024-01-01T00:00:00Z&billed_at[LT]=2024-02-01T00:00:00Z
func (*DirectiveOperator) Encode(rtm *core.DirectiveRuntime) error {
	if len(rtm.Directive.Argv) == 0 {
		return nil
	}

	if rtm.Value.Kind() == reflect.Pointer && rtm.Value.IsNil() {
		return nil
	}

	operand, ok := rtm.Value.Interface().(Operand)
	if !ok {
		return ErrDirectiveRequiresOperand
	}

	name := rtm.Directive.Argv[0]
	for operator, value := range operand.QueryOperators() {
		key := name
		if operator != "" {
			key += "[" + operator + "]"
		}

		rtm.GetRequestBuilder().SetQuery(key, []string{value})
	}

	return nil
}
