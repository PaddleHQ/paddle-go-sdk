package client_test

import (
	"net/http"
	"testing"

	"github.com/PaddleHQ/paddle-go-sdk/v6/internal/client"

	"github.com/ggicci/httpin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDirectiveList(t *testing.T) {
	type args struct {
		name string
		req  any

		expectedURL string
		expectedErr error
	}

	cases := []args{
		{
			name: "unset",
			req:  &ListRequest{},

			expectedURL: "/",
		},
		{
			name: "empty",
			req:  &ListRequest{Status: []listStatus{}},

			expectedURL: "/",
		},
		{
			name: "single value",
			req:  &ListRequest{Status: []listStatus{"active"}},

			expectedURL: "/?status=active",
		},
		{
			name: "multiple values",
			req:  &ListRequest{Status: []listStatus{"active", "archived"}, ID: []string{"pro_1", "pro_2"}},

			expectedURL: "/?id=pro_1%2Cpro_2&status=active%2Carchived",
		},
		{
			name: "wrong field type",
			req:  &NonListRequest{Status: "active"},

			expectedErr: client.ErrDirectiveRequiresStringSlice,
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

type listStatus string

type ListRequest struct {
	Status []listStatus `in:"paddle_list=status"`
	ID     []string     `in:"paddle_list=id"`
}

type NonListRequest struct {
	Status string `in:"paddle_list=status"`
}
