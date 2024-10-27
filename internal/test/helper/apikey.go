package helper

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
)

func GenAPIKey(t *testing.T) string {
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
	return client.AdminUser.Create().SetName(name).SetAPIKey(genApikey()).Save(context.Background())
}

func genApikey() string {
	randBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, randBytes)
	if err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}
