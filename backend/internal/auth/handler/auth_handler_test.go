package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/vnkot/piklnk/configs"
	"github.com/vnkot/piklnk/internal/auth/handler"
	"github.com/vnkot/piklnk/pkg/apierr"
	"github.com/vnkot/piklnk/pkg/jwt"
)

type MockAuthService struct {
	LoginFunc    func(email, password string) (uint, error)
	RegisterFunc func(email, password, name string) (uint, error)
}

func (m *MockAuthService) Login(email, password string) (uint, error) {
	return m.LoginFunc(email, password)
}

func (m *MockAuthService) Register(email, password, name string) (uint, error) {
	return m.RegisterFunc(email, password, name)
}

func TestAuthHandler_Login_Success(t *testing.T) {
	secret := "test_secret"
	userPassword := "password123"
	userEmail := "test@example.com"

	authService := &MockAuthService{
		LoginFunc: func(email, password string) (uint, error) {
			if email == userEmail && password == userPassword {
				return 1, nil
			}
			return 0, errors.New("invalid credentials")
		},
	}

	config := &configs.Config{
		Auth: configs.AuthConfig{
			Secret: secret,
		},
	}

	h := handler.AuthHandler{
		AuthService: authService,
		Config:      config,
	}

	body := bytes.NewBufferString(fmt.Sprintf(`{"email":"%s","password":"%s"}`, userEmail, userPassword))

	req := httptest.NewRequest("POST", "/auth/login", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.Login()(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response handler.LoginResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Failed to parse response:", err)
	}

	_, valid := jwt.NewJWT(config.Auth.Secret).Parse(response.Token)
	if !valid {
		t.Errorf("Generated token is invalid")
	}
}

func TestAuthHandler_Login_InvalidCredentials(t *testing.T) {
	userPassword := "password123"
	userEmail := "test@example.com"

	authService := &MockAuthService{
		LoginFunc: func(email, password string) (uint, error) {
			return 0, errors.New("invalid credentials")
		},
	}

	h := handler.AuthHandler{
		AuthService: authService,
		Config:      &configs.Config{},
	}

	body := bytes.NewBufferString(fmt.Sprintf(`{"email":"%s","password":"%s"}`, userEmail, userPassword))
	req := httptest.NewRequest("POST", "/auth/login", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	h.Login()(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}

	var errResp apierr.APIError
	if err := json.Unmarshal(w.Body.Bytes(), &errResp); err != nil {
		t.Fatal("Failed to parse error response:", err)
	}
}

func TestAuthHandler_Login_InvalidJSON(t *testing.T) {
	h := handler.AuthHandler{}

	body := bytes.NewBufferString(`invalid json`)

	req := httptest.NewRequest("POST", "/auth/login", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.Login()(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400, got %d", w.Code)
	}
}

func TestAuthHandler_Register_Success(t *testing.T) {
	secret := "test_secret"
	userName := "user_name"
	userPassword := "password123"
	userEmail := "test@example.com"

	authService := &MockAuthService{
		RegisterFunc: func(email, password, name string) (uint, error) {
			return 1, nil
		},
	}

	config := &configs.Config{
		Auth: configs.AuthConfig{
			Secret: secret,
		},
	}

	h := handler.AuthHandler{
		AuthService: authService,
		Config:      config,
	}

	body := bytes.NewBufferString(fmt.Sprintf(`{"email":"%s","password":"%s","name":"%s"}`, userEmail, userPassword, userName))

	req := httptest.NewRequest("POST", "/auth/register", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.Register()(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response handler.RegisterResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatal("Failed to parse response:", err)
	}

	_, valid := jwt.NewJWT(config.Auth.Secret).Parse(response.Token)
	if !valid {
		t.Errorf("Generated token is invalid")
	}
}

func TestAuthHandler_Register_ServiceError(t *testing.T) {
	userName := "user_name"
	userPassword := "password123"
	userEmail := "test@example.com"

	authService := &MockAuthService{
		RegisterFunc: func(email, password, name string) (uint, error) {
			return 0, errors.New("user already exists")
		},
	}

	h := handler.AuthHandler{
		AuthService: authService,
		Config:      &configs.Config{},
	}

	body := bytes.NewBufferString(fmt.Sprintf(`{"email":"%s","password":"%s","name":"%s"}`, userEmail, userPassword, userName))

	req := httptest.NewRequest("POST", "/auth/register", body)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	h.Register()(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status 401, got %d", w.Code)
	}
}
