package httptransport

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/auth-service/internal/application"
)

type AuthHandler struct {
	registerUserService *application.RegisterUserService
	loginUserService    *application.LoginUserService
}

func NewAuthHandler(registerUserService *application.RegisterUserService, loginUserService *application.LoginUserService) *AuthHandler {
	return &AuthHandler{
		registerUserService: registerUserService,
		loginUserService:    loginUserService,
	}
}

type registerUserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	log.Printf("Body is: %+v", r.Body)

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "request body is required", http.StatusBadRequest)
			return
		}

		http.Error(w, "invalid json body", http.StatusBadRequest)
		return

	}

	if input.Username == "" || input.Email == "" || input.Password == "" {
		http.Error(w, "username, email and password are required", http.StatusBadRequest)
		return
	}

	user, err := h.registerUserService.Execute(r.Context(), application.RegisterUserInput{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})

	if err != nil {
		if errors.Is(err, application.ErrDuplicateEmail) {
			http.Error(w, "email already exists", http.StatusBadRequest)
			return
		}

		http.Error(w, "could not register user", http.StatusBadRequest)
		return
	}

	response := registerUserResponse{
		ID:       string(user.ID()),
		Username: user.Username(),
		Email:    user.Email().Value(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}

type loginUserResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string
		Password string
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "request body is required", http.StatusBadRequest)
			return
		}

		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	if input.Email == "" || input.Password == "" {
		http.Error(w, "email and password are required", http.StatusBadRequest)
		return
	}

	token, err := h.loginUserService.Execute(
		r.Context(),
		application.LoginUserInput{Email: input.Email, Password: input.Password},
	)

	if err != nil {
		if errors.Is(err, application.ErrTokenCouldNotBeGenerated) {
			http.Error(w, "there has been a server error, try again later", http.StatusInternalServerError)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := loginUserResponse{Token: token}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "could not encode response", http.StatusInternalServerError)
		return
	}
}

type HealthHandler struct {
}

type healthResponse struct {
	Service string `json:"service"`
	Status  string `json:"status"`
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(healthResponse{
		Service: "auth service",
		Status:  "ok",
	})
}
