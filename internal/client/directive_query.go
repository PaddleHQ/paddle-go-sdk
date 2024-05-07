package client

import (
	"strings"

	"github.com/ggicci/httpin/core"
)

// DirectiveQuery extends the core.DirectiveQuery struct from the httpin package and adds support for the "omitempty"
// tag.
type DirectiveQuery struct {
	core.DirectiveQuery
}

// Encode adds the values of the input struct into the http.Request as query parameters.
// If the value is zero and the "omitempty" tag is present, the value will not be added to the query string.
func (d *DirectiveQuery) Encode(rtm *core.DirectiveRuntime) error {
	tag := rtm.Resolver.Field.Tag.Get("in")
	if rtm.Value.IsZero() && strings.Contains(tag, "omitempty") {
		return nil
	}

	return d.DirectiveQuery.Encode(rtm)
}
