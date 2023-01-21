package main

import (
	"fmt"
	"os"

	"github.com/gustagcosta/go-api/api"
	"github.com/gustagcosta/go-api/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	storage := &storage.PgStorage{}
	server := api.NewServer(os.Getenv("PORT"), storage)
	fmt.Println("Server is up")
	server.Start()
}
