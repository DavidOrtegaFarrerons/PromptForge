package httptransport

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthHandler_Register(t *testing.T) {
	authHandler := &AuthHandler{}

	tests := []struct {
		name                   string
		body                   string
		expectedResponseText   string
		expectedResponseStatus int
	}{
		{
			name:                   "Correct data sent",
			body:                   `{"username":"david","email":"david@example.com","password":"secret"}`,
			expectedResponseText:   "",
			expectedResponseStatus: http.StatusOK,
		},
		{
			name:                   "Missing username",
			body:                   `{"email":"david@example.com","password":"secret"}`,
			expectedResponseText:   "username, email and password are required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name:                   "Missing email",
			body:                   `{"username":"david","password":"secret"}`,
			expectedResponseText:   "username, email and password are required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name:                   "Missing password",
			body:                   `{"username":"david","email":"david@example.com"}`,
			expectedResponseText:   "username, email and password are required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name:                   "No body sent",
			body:                   "",
			expectedResponseText:   "request body is required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name:                   "Empty JSON object",
			body:                   `{}`,
			expectedResponseText:   "username, email and password are required\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
		{
			name:                   "Invalid JSON",
			body:                   `{"username":"david"`,
			expectedResponseText:   "invalid json body\n",
			expectedResponseStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(tt.body))
			authHandler.Register(rr, req)

			if tt.expectedResponseText != rr.Body.String() {
				t.Errorf("Expected response: %s got %s", tt.expectedResponseText, rr.Body.String())
			}

			if tt.expectedResponseStatus != rr.Code {
				t.Errorf("Expected http status code: %d got %d", tt.expectedResponseStatus, rr.Code)
			}
		})
	}
}
