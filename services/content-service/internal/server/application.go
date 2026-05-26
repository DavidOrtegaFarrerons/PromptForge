package server

import httptransport "github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/transport/http"

type Application struct {
	healthHandler *httptransport.HealthHandler
	promptHandler *httptransport.PromptHandler
}

func NewApplication(healthHandler *httptransport.HealthHandler, promptHandler *httptransport.PromptHandler) *Application {
	return &Application{
		healthHandler: healthHandler,
		promptHandler: promptHandler,
	}
}
