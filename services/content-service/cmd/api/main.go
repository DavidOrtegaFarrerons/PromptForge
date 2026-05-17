package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type healthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(healthResponse{
			Service: "content service",
			Status:  "ok",
		})
	})

	log.Println("Content Service running on :8083")
	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		log.Fatal(err)
	}
}
