package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/halcyon-org/kizuna/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Fprint(os.Stderr, "Error loading .env file")
	}

	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	server := api.NewServer()
	r := gin.Default()
	api.RegisterHandlers(r, server)

	s := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:" + port,
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
