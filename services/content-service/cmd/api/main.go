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
	grpctransport "github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/transport/grpc"
	httptransport "github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/transport/http"
	_ "github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

	billingGRPCAddr := os.Getenv("BILLING_GRPC_ADDR")
	if billingGRPCAddr == "" {
		log.Fatal("BILLING_GRPC_ADDR is required")
	}

	billingConn, err := grpc.NewClient(billingGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to billing gRPC: %s", err)
	}
	defer billingConn.Close()

	billingClient := grpctransport.NewGRPCBillingClient(billingConn)

	promptRepository := postgres.NewPostgresPromptRepository(db)
	promptIDGenerator := uuid.NewPromptIdGenerator()
	createPromptService := application.NewCreatePromptService(
		promptIDGenerator,
		promptRepository,
		billingClient,
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
