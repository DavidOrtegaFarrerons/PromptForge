package server

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", app.healthHandler.Health)
	mux.HandleFunc("POST /register", app.authHandler.Register)
	mux.HandleFunc("POST /login", app.authHandler.Login)

	return mux
}
