package paddle

import (
	"encoding/json"
	"net/url"

	"github.com/PaddleHQ/paddle-go-sdk/v6/internal/response"
)

// MetricsExplorePagination is the pagination state for the series in an Explore
// query result. Explore paginates with POST over series (segments), so the next
// page is fetched by re-sending the same request body with After set.
type MetricsExplorePagination struct {
	// PerPage is the number of series per page for this response. May differ from the number requested.
	PerPage int
	// Next is the URL for the next page, carrying the `after` query parameter. Always returned, even on the last page.
	Next string
	// HasMore reports whether another page of series exists.
	HasMore bool
	// EstimatedTotal is the total number of series for this response.
	EstimatedTotal int
	// After is the cursor for the next page, parsed from Next. Nil on the last page. Assign it to the next request's After.
	After *string
}

// Ensure MetricsExplore consumes the whole response envelope so pagination isn't lost.
var (
	_ response.UnmarshalsResponses = &MetricsExplore{}
	_ json.Unmarshaler             = &MetricsExplore{}
)

// UnmarshalsResponses marks that this type unmarshals the entire response body,
// not just the unwrapped `data`, so the response's `meta.pagination` is retained.
func (r *MetricsExplore) UnmarshalsResponses() {}

// UnmarshalJSON decodes the full `{data, meta}` envelope: `data` into the result
// set fields and `meta.pagination` into Pagination, parsing the After cursor from
// the `next` URL.
func (r *MetricsExplore) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	type resultSet MetricsExplore

	var env struct {
		Data resultSet `json:"data"`
		Meta struct {
			Pagination struct {
				PerPage        int    `json:"per_page"`
				Next           string `json:"next"`
				HasMore        bool   `json:"has_more"`
				EstimatedTotal int    `json:"estimated_total"`
			} `json:"pagination"`
		} `json:"meta"`
	}

	if err := json.Unmarshal(b, &env); err != nil {
		return err
	}

	*r = MetricsExplore(env.Data)

	p := env.Meta.Pagination
	r.Pagination = MetricsExplorePagination{
		PerPage:        p.PerPage,
		Next:           p.Next,
		HasMore:        p.HasMore,
		EstimatedTotal: p.EstimatedTotal,
		After:          exploreAfterCursor(p.Next, p.HasMore),
	}

	return nil
}

// exploreAfterCursor extracts the `after` cursor from a pagination `next` URL,
// returning nil when there is no further page.
func exploreAfterCursor(next string, hasMore bool) *string {
	if !hasMore || next == "" {
		return nil
	}

	u, err := url.Parse(next)
	if err != nil {
		return nil
	}

	after := u.Query().Get("after")
	if after == "" {
		return nil
	}

	return &after
}
