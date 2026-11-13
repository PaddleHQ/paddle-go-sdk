package client

import (
	"errors"
	"reflect"

	"github.com/ggicci/httpin/core"
)

// Operand is implemented by query values that can carry an operator, such as
// paddle.DateComparison.
type Operand interface {
	QueryOperator() (operator, value string)
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

// Encode sets the query parameter `name[OPERATOR]=value`, or `name=value`
// when the value has no operator. A nil or empty value sets nothing, like omitempty.
// e.g.
//
//	type Request struct {
//	 BilledAt *paddle.DateComparison `in:"paddle_operator=billed_at"`
//	}
//
//	Req{BilledAt: &paddle.DateComparison{Date: "2024-01-01T00:00:00Z", Operator: "LT"}}
//	// ?billed_at[LT]=2024-01-01T00:00:00Z
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

	operator, value := operand.QueryOperator()
	if value == "" {
		return nil
	}

	key := rtm.Directive.Argv[0]
	if operator != "" {
		key += "[" + operator + "]"
	}

	rtm.GetRequestBuilder().SetQuery(key, []string{value})

	return nil
}
