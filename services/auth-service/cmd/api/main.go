package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/postgres"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/security"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/uuid"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/server"
	httptransport "github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/transport/http"
	"golang.org/x/crypto/bcrypt"

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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	userRepository := postgres.NewPostgresUserRepository(db)
	passwordHasher := security.NewBcryptPasswordHasher(bcrypt.DefaultCost)
	userIDGenerator := uuid.NewUserIDGenerator()
	registerUserService := application.NewRegisterUserService(
		userRepository,
		passwordHasher,
		userIDGenerator,
	)

	healthHandler := &httptransport.HealthHandler{}
	authHandler := httptransport.NewAuthHandler(registerUserService)
	app := server.NewApplication(healthHandler, authHandler)

	log.Println("Auth Service running on :8081")
	err = http.ListenAndServe(":8081", app.Routes())
	if err != nil {
		log.Fatal(err)
	}
}
