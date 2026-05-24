package config

import "os"

type Config struct {
	Addr                 string
	AuthenticationSecret string
	AuthServiceUrl       string
	BillingServiceUrl    string
	ContentServiceUrl    string
}

func Load() Config {
	authenticationSecret := env("AUTHENTICATION_TOKEN_SECRET", "")
	if authenticationSecret == "" {
		panic("AUTHENTICATION_TOKEN_SECRET is required")
	}

	return Config{
		Addr:                 env("ADDR", ":8080"),
		AuthenticationSecret: authenticationSecret,
		AuthServiceUrl:       env("AUTH_SERVICE_URL", "http://auth-service:8081"),
		BillingServiceUrl:    env("BILLING_SERVICE_URL", "http://billing-service:8082"),
		ContentServiceUrl:    env("CONTENT_SERVICE_URL", "http://content-service:8083"),
	}
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
