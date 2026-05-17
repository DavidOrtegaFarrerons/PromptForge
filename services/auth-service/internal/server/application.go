package server

import "github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/transport/http"

type Application struct {
	healthHandler *httptransport.HealthHandler
	authHandler   *httptransport.AuthHandler
}

func NewApplication(healthHandler *httptransport.HealthHandler, authHandler *httptransport.AuthHandler) *Application {
	return &Application{
		healthHandler: healthHandler,
		authHandler:   authHandler,
	}
}
