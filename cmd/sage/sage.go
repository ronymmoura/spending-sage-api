package cmd

import (
	"log"

	"github.com/ronymmoura/spending-sage-api/internal/api"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

func Run() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}

	server, err := api.NewServer(config)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}

	server.Start(":8080")
}
