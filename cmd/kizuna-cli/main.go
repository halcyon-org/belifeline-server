package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"log"

	"github.com/halcyon-org/kizuna/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	_ "github.com/lib/pq"
)

func main() {
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

	fmt.Println("Make Admin User")

	var name string

	fmt.Print("name: ")
	fmt.Scanln(&name)

	user, err := AddAdminUser(client, name)
	if err != nil {
		log.Fatalf("failed to create admin user: %v", err)
	}
	fmt.Printf("created admin user: %+v\n", user)
}

func AddAdminUser(client *ent.Client, name string) (*ent.AdminUser, error) {
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

func HashApiKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))
	return hex.EncodeToString(hash[:])
}
