package client

import (
	"errors"
	"reflect"

	"github.com/ggicci/httpin/core"
)

// DirectiveInclude is a directive that adds the custom 'include' query
// parameter to the request, based on the value of the directive's struct field.
type DirectiveInclude struct{}

// Decode is not implemented. This is because we only deal with encoding
// concerns in the client.
func (*DirectiveInclude) Decode(_ *core.DirectiveRuntime) error {
	//nolint:err113 // typed error not needed here.
	return errors.New("not implemented")
}

// ErrDirectiveRequiresBool is returned when the directive's struct field is not
// a boolean.
var ErrDirectiveRequiresBool = errors.New("include directive requires a boolean value")

// Encode adds the directive's value to the comma-separated 'include' query
// parameter.
// The value of the directive's struct field is expected to be a boolean.
// If the value is true, the directive's value query parameter 'include'.
// e.g.
//
//	type Request struct {
//	 IncludeCustomer bool `in:"paddle_include=customer"`
//	}
//
//	Req{IncludeCustomer: true} // ?paddle_include=customer
func (*DirectiveInclude) Encode(rtm *core.DirectiveRuntime) error {
	if len(rtm.Directive.Argv) == 0 {
		return nil
	}

	if rtm.Value.Kind() != reflect.Bool {
		return ErrDirectiveRequiresBool
	}

	if !rtm.Value.Bool() {
		return nil
	}

	// One comma-separated `include`, like every other list param, rather than
	// a repeated key some services may not read.
	encoder := &core.FormEncoder{
		Setter: func(_ string, _ []string) {
			query := rtm.GetRequestBuilder().Query
			if existing := query.Get("include"); existing != "" {
				query.Set("include", existing+","+rtm.Directive.Argv[0])
			} else {
				query.Set("include", rtm.Directive.Argv[0])
			}
		},
	}

	return encoder.Execute(rtm)
}
