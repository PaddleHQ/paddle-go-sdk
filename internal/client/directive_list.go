package client

import (
	"errors"
	"reflect"
	"strings"

	"github.com/ggicci/httpin/core"
)

// DirectiveList is a directive that adds a list query parameter as one
// comma-separated value, the form the API documents for multi-value filters.
type DirectiveList struct{}

// Decode is not implemented. This is because we only deal with encoding
// concerns in the client.
func (*DirectiveList) Decode(_ *core.DirectiveRuntime) error {
	//nolint:err113 // typed error not needed here.
	return errors.New("not implemented")
}

// ErrDirectiveRequiresStringSlice is returned when the directive's struct
// field is not a slice of a string kind.
var ErrDirectiveRequiresStringSlice = errors.New("list directive requires a slice of strings")

// Encode sets the query parameter `name=a,b,c`. A nil or empty slice sets
// nothing, like omitempty.
// e.g.
//
//	type Request struct {
//	 Status []paddle.Status `in:"paddle_list=status"`
//	}
//
//	Req{Status: []paddle.Status{paddle.StatusActive, paddle.StatusArchived}}
//	// ?status=active,archived
func (*DirectiveList) Encode(rtm *core.DirectiveRuntime) error {
	if len(rtm.Directive.Argv) == 0 {
		return nil
	}

	if rtm.Value.Kind() != reflect.Slice || rtm.Value.Type().Elem().Kind() != reflect.String {
		return ErrDirectiveRequiresStringSlice
	}

	if rtm.Value.Len() == 0 {
		return nil
	}

	values := make([]string, rtm.Value.Len())
	for i := range values {
		values[i] = rtm.Value.Index(i).String()
	}

	rtm.GetRequestBuilder().SetQuery(rtm.Directive.Argv[0], []string{strings.Join(values, ",")})

	return nil
}
