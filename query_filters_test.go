package paddle_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"
	"github.com/stretchr/testify/require"
)

// Every SDK example lists with an empty filter, which is how typed enum query
// params shipped unencodable. These send non-empty filters and check the wire.
func TestQueryFilters(t *testing.T) {
	tests := []struct {
		name string
		call func(ctx context.Context, c *paddle.SDK) error
		want string
	}{
		{
			name: "verification type",
			call: func(ctx context.Context, c *paddle.SDK) error {
				_, err := c.ListVerifications(ctx, &paddle.ListVerificationsRequest{
					VerificationType: paddle.PtrTo(paddle.VerificationTypeDomain),
				})
				return err
			},
			want: "verification_type=domain",
		},
		{
			name: "metrics explore entity",
			call: func(ctx context.Context, c *paddle.SDK) error {
				_, err := c.ListExploreEntities(ctx, &paddle.ListExploreEntitiesRequest{
					Entity: paddle.PtrTo(paddle.MetricEntityTransactionsCompleted),
				})
				return err
			},
			want: "entity=transactions.completed",
		},
		{
			name: "enum arrays comma-joined, order by",
			call: func(ctx context.Context, c *paddle.SDK) error {
				_, err := c.ListTransactions(ctx, &paddle.ListTransactionsRequest{
					ID:      []string{"txn_1", "txn_2"},
					Status:  []paddle.TransactionStatus{paddle.TransactionStatusPaid, paddle.TransactionStatusCompleted},
					Origin:  []paddle.TransactionOriginQuery{paddle.TransactionOriginQueryAPI},
					OrderBy: paddle.PtrTo(paddle.TransactionOrderByBilledAtDesc),
				})
				return err
			},
			want: "id=txn_1%2Ctxn_2&order_by=billed_at%5BDESC%5D&origin=api&status=paid%2Ccompleted",
		},
		{
			name: "adjustment actions",
			call: func(ctx context.Context, c *paddle.SDK) error {
				_, err := c.ListAdjustments(ctx, &paddle.ListAdjustmentsRequest{
					Action: []paddle.AdjustmentAction{paddle.AdjustmentActionCredit, paddle.AdjustmentActionRefund},
				})
				return err
			},
			want: "action=credit%2Crefund",
		},
		{
			name: "price billing cycle",
			call: func(ctx context.Context, c *paddle.SDK) error {
				_, err := c.ListPrices(ctx, &paddle.ListPricesRequest{
					BillingCycleInterval:  paddle.PtrTo(paddle.IntervalMonth),
					BillingCycleFrequency: paddle.PtrTo(3),
				})
				return err
			},
			want: "billing_cycle.frequency=3&billing_cycle.interval=month",
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

			require.NoError(t, tt.call(t.Context(), c))
			require.Equal(t, tt.want, got)
		})
	}
}
