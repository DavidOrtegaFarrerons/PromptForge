package server

import (
	"log"
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/config"
	"github.com/DavidOrtegaFarrerons/promptforge/services/api-gateway/internal/middleware"
)

type Application struct {
	cfg          config.Config
	srv          *http.ServeMux
	tokenDecoder middleware.TokenDecoder
}

func NewApplication(tokenDecoder middleware.TokenDecoder, cfg config.Config) *Application {
	srv := http.NewServeMux()
	return &Application{
		srv:          srv,
		tokenDecoder: tokenDecoder,
		cfg:          cfg,
	}
}

func (app *Application) Start() error {
	app.routes(app.srv)
	log.Printf("Service running on %s \n", app.cfg.Addr)
	return http.ListenAndServe(app.cfg.Addr, app.srv)
}
