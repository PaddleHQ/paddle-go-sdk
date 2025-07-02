package paddle_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	paddle "github.com/PaddleHQ/paddle-go-sdk/v4"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Request struct {
	ID int `json:"id"`
}

func TestCollection(t *testing.T) {
	firstPageJSON := []byte(`{
		"data": [
			{"id": 1}
		],
		"meta": {
			"pagination": {
				"next_url": "https://example.com/next",
				"per_page": 1,
				"has_more": true,
				"estimated_total": 2
			}
		}
	}`)

	t.Run("successful pagination", func(t *testing.T) {
		secondPageJSON := []byte(`{
			"data": [
				{"id": 2}
			],
			"meta": {
				"pagination": {
					"next_url": "https://example.com/next",
					"per_page": 1,
					"has_more": true,
					"estimated_total": 2
				}
			}
		}`)

		thirdPageJSON := []byte(`{
			"data": [],
			"meta": {
				"pagination": {
					"next_url": "https://example.com/next",
					"per_page": 1,
					"has_more": false,
					"estimated_total": 3
				}
			}
		}`)

		c := &paddle.Collection[*Request]{}
		m := &MockDoer{t: t, pages: [][]byte{secondPageJSON, thirdPageJSON}}
		c.Wants(m)

		err := json.Unmarshal(firstPageJSON, &c)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 2, c.EstimatedTotal())
		assert.Equal(t, true, c.HasMore())
		assert.Equal(t, 1, c.PerPage())

		called := false
		i := 1

		err = c.IterErr(context.Background(), func(r *Request) error {
			called = true
			assert.Equal(t, i, r.ID)
			i++

			return nil
		})
		assert.NoError(t, err)

		assert.True(t, called)
	})

	t.Run("error on iteration", func(t *testing.T) {
		wantErr := errors.New("something happened")
		c := &paddle.Collection[*Request]{}
		m := &MockDoer{t: t, err: wantErr}
		c.Wants(m)

		err := json.Unmarshal(firstPageJSON, &c)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 2, c.EstimatedTotal())
		assert.Equal(t, true, c.HasMore())
		assert.Equal(t, 1, c.PerPage())

		called := false
		i := 1

		err = c.IterErr(context.Background(), func(r *Request) error {
			called = true
			assert.Equal(t, i, r.ID)
			i++

			return nil
		})
		assert.Error(t, wantErr)
		assert.True(t, called)
	})

	t.Run("result not mutate on next page", func(t *testing.T) {
		secondPageJSON := []byte(`{
			"data": [
				{"id": 2}
			],
			"meta": {
				"pagination": {
					"next_url": "https://example.com/next",
					"per_page": 1,
					"has_more": false,
					"estimated_total": 2
				}
			}
		}`)

		c := &paddle.Collection[*Request]{}
		m := &MockDoer{t: t, pages: [][]byte{secondPageJSON}}
		c.Wants(m)

		err := json.Unmarshal(firstPageJSON, &c)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 2, c.EstimatedTotal())
		assert.Equal(t, true, c.HasMore())
		assert.Equal(t, 1, c.PerPage())

		called := false
		var foundReq *Request

		err = c.IterErr(context.Background(), func(r *Request) error {
			if called {
				return nil
			}

			foundReq = r
			called = true

			return nil
		})

		assert.NoError(t, err)
		assert.True(t, called)
		assert.Equal(t, 1, foundReq.ID)
	})
}

type MockDoer struct {
	i     int
	t     *testing.T
	pages [][]byte
	err   error
}

func (m *MockDoer) Do(_ context.Context, method, path string, src, dst any) error {
	m.t.Helper()

	assert.Equal(m.t, "GET", method)

	typ := &paddle.Collection[*Request]{}
	require.IsType(m.t, &typ, dst)

	if m.err != nil {
		return m.err
	}

	defer func() {
		m.i++
	}()

	return json.Unmarshal(m.pages[m.i], dst)
}
