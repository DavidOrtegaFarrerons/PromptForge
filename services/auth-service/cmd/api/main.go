package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/postgres"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/rabbitmq"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/security"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/token"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/infrastructure/uuid"
	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/server"
	httptransport "github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/transport/http"
	_ "github.com/jackc/pgx/v5/stdlib"
	amqp "github.com/rabbitmq/amqp091-go"
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

	amqpUser := os.Getenv("AMQP_USER")
	amqpPassword := os.Getenv("AMQP_PASSWORD")
	amqpHost := os.Getenv("AMQP_HOST")
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", amqpUser, amqpPassword, amqpHost))
	if err != nil {
		log.Fatalf("connecting to AMQP failed: %s", err.Error())
	}
	defer conn.Close()
	amqpChannel, err := conn.Channel()
	if err != nil {
		log.Fatalf("getting AMQP channel failed: %s", err.Error())
	}

	exchangeName := "promptforge.events"
	err = amqpChannel.ExchangeDeclare(
		exchangeName,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatalf("Declaring AMQP exchange %s failed: %s", exchangeName, err.Error())
	}

	userRepository := postgres.NewPostgresUserRepository(db)
	passwordHasher := security.NewBcryptPasswordHasher()
	userIDGenerator := uuid.NewUserIDGenerator()
	eventPublisher := rabbitmq.NewRabbitEventPublisher(amqpChannel)
	registerUserService := application.NewRegisterUserService(
		userRepository,
		passwordHasher,
		userIDGenerator,
		eventPublisher,
	)

	tokenSecret := os.Getenv("TOKEN_SECRET")
	if tokenSecret == "" {
		log.Fatal("TOKEN_SECRET env variable is missing")
	}

	tokenGenerator := token.NewJwtTokenGenerator([]byte(tokenSecret), 15*time.Minute)
	loginUserService := application.NewLoginUserService(
		userRepository,
		passwordHasher,
		tokenGenerator,
	)

	healthHandler := &httptransport.HealthHandler{}
	authHandler := httptransport.NewAuthHandler(registerUserService, loginUserService)
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
