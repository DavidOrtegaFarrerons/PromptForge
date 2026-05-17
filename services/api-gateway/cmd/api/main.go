package main

import (
	"log"
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/config"
	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Register(mux, config.Load())

	log.Println("Service running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
