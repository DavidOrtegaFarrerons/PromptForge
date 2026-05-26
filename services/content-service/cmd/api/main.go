package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/infrastructure/postgres"
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/infrastructure/uuid"
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/server"
	httptransport "github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/transport/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseDSN := os.Getenv("CONTENT_DATABASE_DSN")
	if databaseDSN == "" {
		log.Fatal("CONTENT_DATABASE_DSN is required")
	}

	migrationsPath := os.Getenv("CONTENT_MIGRATIONS_PATH")
	if migrationsPath == "" {
		migrationsPath = "file://migrations"
	}

	err := postgres.RunMigrations(databaseDSN, migrationsPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("pgx", databaseDSN)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	promptRepository := postgres.NewPostgresPromptRepository(db)
	promptIDGenerator := uuid.NewPromptIdGenerator()
	createPromptService := application.NewCreatePromptService(
		promptIDGenerator,
		promptRepository,
	)

	healthHandler := &httptransport.HealthHandler{}
	promptHandler := httptransport.NewPromptHandler(createPromptService)
	app := server.NewApplication(healthHandler, promptHandler)

	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatal("ADDR env var required")
	}
	log.Printf("Content Service running on %s", addr)
	err = http.ListenAndServe(addr, app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
