package paddle_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
	"github.com/stretchr/testify/require"
)

func TestDateComparison(t *testing.T) {
	tests := []struct {
		name     string
		billedAt *paddle.DateComparison
		want     string
	}{
		{
			name:     "operator",
			billedAt: &paddle.DateComparison{Date: "2024-01-01T00:00:00Z", Operator: paddle.ComparisonOperatorLT},
			want:     "billed_at%5BLT%5D=2024-01-01T00%3A00%3A00Z",
		},
		{
			name:     "exact match",
			billedAt: &paddle.DateComparison{Date: "2024-01-01T00:00:00Z"},
			want:     "billed_at=2024-01-01T00%3A00%3A00Z",
		},
		{
			name:     "empty date",
			billedAt: &paddle.DateComparison{Operator: paddle.ComparisonOperatorGTE},
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
			var got string
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				got = r.URL.RawQuery
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"data":[],"meta":{"pagination":{"per_page":50,"next":"","has_more":false,"estimated_total":0}}}`))
			}))
			t.Cleanup(s.Close)

			c, err := paddle.New("pdl_test_key", paddle.WithBaseURL(s.URL))
			require.NoError(t, err)

			_, err = c.ListTransactions(t.Context(), &paddle.ListTransactionsRequest{BilledAt: tt.billedAt})
			require.NoError(t, err)
			require.Equal(t, tt.want, got)
		})
	}
}
