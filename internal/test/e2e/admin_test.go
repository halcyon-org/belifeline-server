package e2e

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/test/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestClientSet(t *testing.T) {
	admin, _, _, _ := helper.NewClients(t)
	apiKey := helper.GenAPIKey(t)

	for _, client := range admin {
		grp, ctx := errgroup.WithContext(context.Background())
		req := &connect.Request[mainv1.ClientSetRequest]{Msg: &mainv1.ClientSetRequest{Username: "test"}}
		req.Header().Add(interceptor.AuthAPIKeyHeader, apiKey)
		res, err := client.ClientSet(ctx, req)
		if err != nil {
			t.Fatalf("client.ClientSet() = %v", err)
		}

		require.NoError(t, grp.Wait())
		assert.Equal(t, req.Msg.Username, *res.Msg.ClientInformation.Username)
		assert.NotNil(t, res.Msg.ClientInformation.ClientId)
		assert.NotNil(t, res.Msg.ClientInformation.ApiKey)
		assert.NotNil(t, res.Msg.ClientInformation.CreatedAt)
		assert.NotNil(t, res.Msg.ClientInformation.LastUsedAt)
		assert.NotNil(t, res.Msg.ClientInformation.LastUpdatedAt)
	}
}
