package client_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/PaddleHQ/paddle-go-sdk/v2/internal/client"
	"github.com/PaddleHQ/paddle-go-sdk/v2/internal/response"
	"github.com/PaddleHQ/paddle-go-sdk/v2/pkg/paddleerr"
)

func ptr[V any](v V) *V {
	return &v
}

func TestClient(t *testing.T) {
	type args struct {
		name string

		req  any
		res  any
		path string

		expectedRes   any
		expectedPath  string
		expectedError *paddleerr.Error
	}

	cases := []args{
		{
			name: "simple request in/out",

			req:  &Req{ID: "123", Optional: ptr("test")},
			res:  &Res{},
			path: "/",

			expectedRes:  &Res{ID: "123"},
			expectedPath: "/?id=123&optional=test",
		},
		{
			name: "path rewrite",

			req:  &ReqPath{ID: "123"},
			res:  &Res{},
			path: "/test/{id}",

			expectedRes:  &Res{ID: "123"},
			expectedPath: "/test/123",
		},
		{
			name: "simple request in/out",

			req:  &Req{ID: "123"},
			res:  &Res{},
			path: "/",

			expectedRes:  &Res{ID: "123"},
			expectedPath: "/?id=123",
		},
	}

	apiKey := "my-example-api-key-here"

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.expectedPath, r.URL.String())
				assert.Equal(t, "Bearer "+apiKey, r.Header.Get("Authorization"))

				res := response.Response[any]{
					Data: tt.expectedRes,
				}

				if tt.expectedError != nil {
					res.Data, res.Error = nil, tt.expectedError
				}

				j, err := json.Marshal(res)
				require.NoError(t, err)

				w.WriteHeader(http.StatusOK)

				_, err = w.Write(j)
				require.NoError(t, err)
			}))

			parsedURL, err := url.Parse(s.URL)
			require.NoError(t, err)

			c, err := client.New(http.DefaultClient, apiKey, parsedURL)
			require.NoError(t, err)

			err = c.Do(context.Background(), "GET", tt.path, tt.req, &tt.res)
			if tt.expectedError != nil {
				assert.ErrorIs(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedRes, tt.res)
		})
	}
}

type Req struct {
	ID       string  `in:"query=id"`
	Optional *string `in:"query=optional;omitempty"`
}

type ReqPath struct {
	ID string `in:"path=id"`
}

type Res struct {
	ID string `json:"id"`
}
