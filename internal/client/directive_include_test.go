package client_test

import (
	"net/http"
	"testing"

	"github.com/PaddleHQ/paddle-go-sdk/v3/internal/client"

	"github.com/ggicci/httpin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirectiveInclude(t *testing.T) {
	type args struct {
		name string
		req  any

		expectedURL string
		expectedErr error
	}

	cases := []args{
		{
			name: "no include",
			req:  &Request{},

			expectedURL: "/",
		},
		{
			name: "include one",
			req: &Request{
				IncludeBusiness: true,
				IncludeCustomer: false,
			},

			expectedURL: "/?include=business",
		},
		{
			name: "include both",
			req: &Request{
				IncludeBusiness: true,
				IncludeCustomer: true,
			},

			expectedURL: "/?include=business&include=customer",
		},
		{
			name: "wrong field type",
			req: &NonBoolRequest{
				IncludeBusiness: "a",
			},

			expectedErr: client.ErrDirectiveRequiresBool,
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
			assert.Equal(t, req.URL.String(), c.expectedURL)
		})
	}
}

type Request struct {
	IncludeBusiness bool `in:"paddle_include=business"`
	IncludeCustomer bool `in:"paddle_include=customer"`
}

type NonBoolRequest struct {
	IncludeBusiness string `in:"paddle_include=business"`
}
