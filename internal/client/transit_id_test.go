package client_test

import (
	"context"
	"testing"

	"github.com/PaddleHQ/paddle-go-sdk/v4/internal/client"

	"github.com/stretchr/testify/assert"
)

func TestTransitID(t *testing.T) {
	ctx := context.Background()
	ctx = client.ContextWithTransitID(ctx, "id")

	transitID := client.TransitIDFromContext(ctx)

	assert.Equal(t, "id", transitID)
}
