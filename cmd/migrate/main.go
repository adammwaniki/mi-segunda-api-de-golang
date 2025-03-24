package main

import (
	"log"

	"github.com/adammwaniki/mi-segunda-api-de-golang/cmd/api"
)

// This will be the entrypoint for our API

func main(){
	// Creating a new API/server instance
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}