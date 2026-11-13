package client_test

import (
	"net/http"
	"testing"

	"github.com/PaddleHQ/paddle-go-sdk/v6/internal/client"

	"github.com/ggicci/httpin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirectiveOperator(t *testing.T) {
	type args struct {
		name string
		req  any

		expectedURL string
		expectedErr error
	}

	cases := []args{
		{
			name: "unset",
			req:  &OperatorRequest{},

			expectedURL: "/",
		},
		{
			name: "with operator",
			req: &OperatorRequest{
				BilledAt: &operand{operator: "LT", value: "2024-01-01T00:00:00Z"},
			},

			expectedURL: "/?billed_at%5BLT%5D=2024-01-01T00%3A00%3A00Z",
		},
		{
			name: "exact match",
			req: &OperatorRequest{
				BilledAt: &operand{value: "2024-01-01T00:00:00Z"},
			},

			expectedURL: "/?billed_at=2024-01-01T00%3A00%3A00Z",
		},
		{
			name: "empty value",
			req: &OperatorRequest{
				BilledAt: &operand{operator: "GTE"},
			},

			expectedURL: "/",
		},
		{
			name: "wrong field type",
			req: &NonOperatorRequest{
				BilledAt: "2024-01-01T00:00:00Z",
			},

			expectedErr: client.ErrDirectiveRequiresOperand,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			req, err := httpin.NewRequest(http.MethodGet, "/", c.req)
			if c.expectedErr != nil {
				require.Error(t, err)
				assert.ErrorIs(t, err, c.expectedErr)

				return
			}

			require.NoError(t, err)
			assert.Equal(t, c.expectedURL, req.URL.String())
		})
	}
}

type operand struct {
	operator string
	value    string
}

func (c operand) QueryOperator() (string, string) {
	return c.operator, c.value
}

type OperatorRequest struct {
	BilledAt *operand `in:"paddle_operator=billed_at"`
}

type NonOperatorRequest struct {
	BilledAt string `in:"paddle_operator=billed_at"`
}
