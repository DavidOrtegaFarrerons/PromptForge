package server

import "github.com/DavidOrtegaFarrerons/promptforge/services/billing-service/internal/transport/http"

type Application struct {
	healthHandler *httptransport.HealthHandler
}

func NewApplication(healthHandler *httptransport.HealthHandler) *Application {
	return &Application{
		healthHandler: healthHandler,
	}
}
