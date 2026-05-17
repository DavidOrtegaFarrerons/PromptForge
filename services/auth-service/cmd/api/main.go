package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/postgres"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/security"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/uuid"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/server"
	httptransport "github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/transport/http"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseDSN := os.Getenv("AUTH_DATABASE_DSN")
	if databaseDSN == "" {
		log.Fatal("AUTH_DATABASE_DSN is required")
	}

	migrationsPath := os.Getenv("AUTH_MIGRATIONS_PATH")
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

	userRepository := postgres.NewPostgresUserRepository(db)
	passwordHasher := security.NewBcryptPasswordHasher()
	userIDGenerator := uuid.NewUserIDGenerator()
	registerUserService := application.NewRegisterUserService(
		userRepository,
		passwordHasher,
		userIDGenerator,
	)

	healthHandler := &httptransport.HealthHandler{}
	authHandler := httptransport.NewAuthHandler(registerUserService)
	app := server.NewApplication(healthHandler, authHandler)

	addr := os.Getenv("ADDR")
	if addr == "" {
		log.Fatal("ADDR env var required")
	}
	log.Printf("Auth Service running on %s", addr)
	err = http.ListenAndServe(addr, app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
