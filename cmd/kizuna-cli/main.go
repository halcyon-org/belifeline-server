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
	"net/url"

	"github.com/halcyon-org/kizuna/gen/ent"
	"github.com/halcyon-org/kizuna/internal/adapter/repository/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.NewConfigRepository()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	config := cfg.GetConfig()

	dataSource := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(config.Postgres.User, config.Postgres.Password),
		Host:     config.Postgres.Host,
		Path:     config.Postgres.DB,
		RawQuery: "sslmode=disable",
	}

	client, err := ent.Open(
		"postgres",
		dataSource.String(),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	log.Println("Make Admin User")

	var name string

	log.Print("name: ")

	if _, err = fmt.Scanln(&name); err != nil {
		log.Fatalf("failed to read name: %v", err)
	}

	user, err := AddAdminUser(client, name)
	if err != nil {
		log.Fatalf("failed to create admin user: %v", err)
	}

	log.Printf("created admin user: %+v\n", user)
}

func AddAdminUser(client *ent.Client, name string) (*ent.AdminUser, error) {
	return client.AdminUser.Create().SetName(name).SetAPIKey(genAPIkey()).Save(context.Background())
}

const length = 64

func genAPIkey() string {
	randBytes := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, randBytes); err != nil {
		panic(err)
	}

	return base64.RawURLEncoding.WithPadding(base64.NoPadding).EncodeToString(randBytes)
}

func HashAPIKey(apiKey string) string {
	hash := sha256.Sum256([]byte(apiKey))

	src := hash[:]
	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	return string(dst)
}
