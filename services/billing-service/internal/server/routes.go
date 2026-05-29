package server

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.healthHandler.Health)
	return mux
}
