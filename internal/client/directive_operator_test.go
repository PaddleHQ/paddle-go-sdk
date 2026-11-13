package client_test

import (
	"iter"
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
				BilledAt: &operand{{"LT", "2024-01-01T00:00:00Z"}},
			},

			expectedURL: "/?billed_at%5BLT%5D=2024-01-01T00%3A00%3A00Z",
		},
		{
			name: "multiple operators",
			req: &OperatorRequest{
				BilledAt: &operand{{"GTE", "2024-01-01T00:00:00Z"}, {"LT", "2024-02-01T00:00:00Z"}},
			},

			expectedURL: "/?billed_at%5BGTE%5D=2024-01-01T00%3A00%3A00Z&billed_at%5BLT%5D=2024-02-01T00%3A00%3A00Z",
		},
		{
			name: "exact match",
			req: &OperatorRequest{
				BilledAt: &operand{{"", "2024-01-01T00:00:00Z"}},
			},

			expectedURL: "/?billed_at=2024-01-01T00%3A00%3A00Z",
		},
		{
			name: "no operators",
			req: &OperatorRequest{
				BilledAt: &operand{},
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

type operand [][2]string

func (o operand) QueryOperators() iter.Seq2[string, string] {
	return func(yield func(string, string) bool) {
		for _, pair := range o {
			if !yield(pair[0], pair[1]) {
				return
			}
		}
	}
}

type OperatorRequest struct {
	BilledAt *operand `in:"paddle_operator=billed_at"`
}

type NonOperatorRequest struct {
	BilledAt string `in:"paddle_operator=billed_at"`
}
