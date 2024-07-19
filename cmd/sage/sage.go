package cmd

import (
	"github.com/ronymmoura/spending-sage-api/internal/api"
	"github.com/ronymmoura/spending-sage-api/internal/util"
)

func Run() {
	config, err := util.LoadConfig(".env")
	if err != nil {
		panic("Error loading config file")
	}

	server, err := api.NewServer(config)
	if err != nil {
		panic("Opa")
	}

	server.Start(":8080")
}
