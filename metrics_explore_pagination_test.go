package paddle_test

import (
	"encoding/json"
	"os"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v6"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMetricsExploreUnmarshal(t *testing.T) {
	t.Run("unwraps data and retains pagination with a next cursor", func(t *testing.T) {
		b, err := os.ReadFile("testdata/metrics_explore.json")
		require.NoError(t, err)

		var res paddle.MetricsExplore
		require.NoError(t, json.Unmarshal(b, &res))

		assert.Equal(t, paddle.MetricEntityTransactionsCompleted, res.Entity)
		assert.Equal(t, paddle.MetricsBucketIntervalMonth, res.Interval)
		require.Len(t, res.Series, 1)
		assert.Equal(t, "pro_01h1vjfevh5etwq3rb416a23h2", res.Series[0].Dimensions["product"])
		assert.Equal(t, "120000", res.Series[0].Timeseries[0].Measures["sum_gross_revenue"])

		assert.Equal(t, 5, res.Pagination.PerPage)
		assert.True(t, res.Pagination.HasMore)
		assert.Equal(t, 42, res.Pagination.EstimatedTotal)
		require.NotNil(t, res.Pagination.After)
		assert.Equal(t,
			"eyJ2IjoxLCJxaWQiOiIwMWpuOHg0azJwN3E5ciIsIm9mZiI6NSwicWgiOiJzaGEyNTY6OWYyYzRkMWUiLCJleHAiOjE3ODU1Njc2MDB9",
			*res.Pagination.After,
		)
	})

	t.Run("after is nil on the last page", func(t *testing.T) {
		b := []byte(`{
			"data": {"entity": "checkouts", "interval": "day", "series": []},
			"meta": {"pagination": {
				"per_page": 5,
				"next": "https://api.paddle.com/metrics/explore?after=should_be_ignored",
				"has_more": false,
				"estimated_total": 3
			}}
		}`)

		var res paddle.MetricsExplore
		require.NoError(t, json.Unmarshal(b, &res))

		assert.Equal(t, paddle.MetricEntityCheckouts, res.Entity)
		assert.False(t, res.Pagination.HasMore)
		assert.Equal(t, 3, res.Pagination.EstimatedTotal)
		assert.Nil(t, res.Pagination.After)
	})
}
