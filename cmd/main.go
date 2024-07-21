package main

import (
	"github.com/satishcg12/echomers/internal"
	"github.com/satishcg12/echomers/internal/utils"
)

func main() {

	utils.LoadConfig()

	server := internal.NewServer(internal.ServerConfig{
		Host: utils.GetEnvOrDefault("HOST", "localhost"),
		Port: utils.GetEnvOrDefault("PORT", "8080"),
	})

	server.Start()

}
