package paddle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
	"github.com/stretchr/testify/require"
)

const (
	jan = "2024-01-01T00:00:00Z"
	feb = "2024-02-01T00:00:00Z"
)

func TestDateFilter(t *testing.T) {
	tests := []struct {
		name     string
		billedAt *paddle.DateFilter
		want     string
	}{
		{
			name:     "range",
			billedAt: &paddle.DateFilter{GTE: paddle.PtrTo(jan), LT: paddle.PtrTo(feb)},
			want:     "billed_at%5BGTE%5D=2024-01-01T00%3A00%3A00Z&billed_at%5BLT%5D=2024-02-01T00%3A00%3A00Z",
		},
		{
			name:     "exact match",
			billedAt: &paddle.DateFilter{Equals: paddle.PtrTo(jan)},
			want:     "billed_at=2024-01-01T00%3A00%3A00Z",
		},
		{
			name:     "empty date",
			billedAt: &paddle.DateFilter{GTE: paddle.PtrTo("")},
			want:     "",
		},
		{
			name:     "unset",
			billedAt: nil,
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, got := newQueryRecorder(t)

			_, err := c.ListTransactions(t.Context(), &paddle.ListTransactionsRequest{BilledAt: tt.billedAt})
			require.NoError(t, err)
			require.Equal(t, tt.want, *got)
		})
	}
}

func TestDateRange(t *testing.T) {
	c, got := newQueryRecorder(t)

	_, err := c.ListSubscriptionHistory(t.Context(), &paddle.ListSubscriptionHistoryRequest{
		SubscriptionID: "sub_01",
		OccurredAt:     &paddle.DateRange{GTE: paddle.PtrTo(jan), LTE: paddle.PtrTo(feb)},
	})
	require.NoError(t, err)
	require.Equal(t, "occurred_at%5BGTE%5D=2024-01-01T00%3A00%3A00Z&occurred_at%5BLTE%5D=2024-02-01T00%3A00%3A00Z", *got)
}

// newQueryRecorder returns a client whose list requests record their raw query.
func newQueryRecorder(t *testing.T) (*paddle.SDK, *string) {
	t.Helper()

	var got string
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.URL.RawQuery
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"data":[],"meta":{"pagination":{"per_page":50,"next":"","has_more":false,"estimated_total":0}}}`))
	}))
	t.Cleanup(s.Close)

	c, err := paddle.New("pdl_test_key", paddle.WithBaseURL(s.URL))
	require.NoError(t, err)

	return c, &got
}
