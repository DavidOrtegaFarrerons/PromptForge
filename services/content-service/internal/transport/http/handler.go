package httptransport

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/application"
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/domain"
	"github.com/DavidOrtegaFarrerons/promptforge/services/content-service/internal/httpcontext"
)

type PromptHandler struct {
	createPromptService *application.CreatePromptService
}

func NewPromptHandler(createPromptService *application.CreatePromptService) *PromptHandler {
	return &PromptHandler{
		createPromptService: createPromptService,
	}
}

type createPromptResponse struct {
	ID           string   `json:"id"`
	Content      string   `json:"content"`
	Placeholders []string `json:"placeholders"`
	Tags         []string `json:"tags"`
}

func (h *PromptHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
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

	ownerID := httpcontext.ContextGetUserID(r)
	if ownerID == "" {
		http.Error(w, "you must be authenticated to create a prompt", http.StatusUnauthorized)
		return
	}

	prompt, err := h.createPromptService.Execute(r.Context(), application.CreatePromptInput{

		OwnerID: ownerID,
		Title:   input.Title,
		Content: input.Content,
		Tags:    input.Tags,
	})

	if err != nil {
		switch {
		case errors.Is(err, domain.ErrTitleMinLength):
			http.Error(w, "title must be at least 4 characters", http.StatusUnprocessableEntity)
		case errors.Is(err, domain.ErrTitleMaxLength):
			http.Error(w, "title must be at most 64 characters", http.StatusUnprocessableEntity)
		case errors.Is(err, domain.ErrContentEmpty):
			http.Error(w, "content cannot be empty", http.StatusUnprocessableEntity)
		default:
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	response := createPromptResponse{
		ID:           string(prompt.PromptID()),
		Content:      prompt.Template().Content(),
		Placeholders: prompt.Template().Placeholders(),
		Tags:         prompt.Tags(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

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
		Service: "content service",
		Status:  "ok",
	})
}
