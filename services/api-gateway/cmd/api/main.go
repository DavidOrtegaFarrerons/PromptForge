package main

import (
	"log"

	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/config"
	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/middleware"
	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/server"
)

func main() {
	cfg := config.Load()

	tokenDecoder := middleware.NewJwtTokenDecoder(cfg.AuthenticationSecret)
	err := server.NewApplication(tokenDecoder, cfg).Start()
	if err != nil {
		log.Fatal(err)
	}
}
