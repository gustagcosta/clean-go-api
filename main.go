package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gustagcosta/go-api/api"
	"github.com/gustagcosta/go-api/storage"
	"github.com/joho/godotenv"
)

func main() {
	// carrega as variaveis de ambiente do .env
	godotenv.Load()

	// define o banco de dados de nós iremos usar
	storage := &storage.MemoryStorage{}
	storage.Connect("EM MEMORIA")

	// Criamos o servidor da API, da pra abstrair mais aqui nessa etapa
	server := api.NewServer(os.Getenv("PORT"), storage)
	fmt.Println("Server is up")

	log.Fatal(server.Start())
}
