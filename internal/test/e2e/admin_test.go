package e2e

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"testing"

	"connectrpc.com/connect"
	mainv1 "github.com/halcyon-org/kizuna/gen/belifeline/v1"
	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/interceptor"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	"github.com/halcyon-org/kizuna/internal/test/helper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func genAPIKey(t *testing.T) string {
	t.Helper()

	cfg, err := config.NewConfigRepository()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	config := cfg.GetConfig()

	client, err := ent.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.Postgres.User, config.Postgres.Password, config.Postgres.Host, config.Postgres.Port, config.Postgres.DB))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	user, err := addAdminUser(client, "test admin user")
	if err != nil {
		log.Fatalf("failed to create admin user: %v", err)
	}

	return user.APIKey
}
func addAdminUser(client *ent.Client, name string) (*ent.AdminUser, error) {
	return client.AdminUser.Create().SetName(name).SetAPIKey(genAPIkey()).Save(context.Background())
}

func genAPIkey() string {
	randBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, randBytes)
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}

func TestClientSet(t *testing.T) {
	admin, _, _, _ := helper.NewClients(t)
	apiKey := genAPIKey(t)

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
