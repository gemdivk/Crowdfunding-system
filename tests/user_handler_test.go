package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/handlers"
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	r.GET("/verify-email", handlers.VerifyEmail)
	r.POST("/logout", handlers.LogoutUser)
	return r
}

func TestRegisterUser(t *testing.T) {
	os.Setenv("SMTP_USER", "test@example.com")
	os.Setenv("SMTP_HOST", "smtp.example.com")
	os.Setenv("SMTP_PASSWORD", "password")

	router := setupRouter()

	user := models.User{
		Name:     "Test User",
		Email:    "testuser@example.com",
		Password: "testpassword",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User registered successfully")
}

func TestLoginUser(t *testing.T) {
	router := setupRouter()

	user := models.User{
		Email:    "testuser@example.com",
		Password: "testpassword",
	}

	jsonUser, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonUser))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login successful")
}

func TestVerifyEmail(t *testing.T) {
	mockToken := "your_mock_token"

	router := setupRouter()

	req, _ := http.NewRequest("GET", "/verify-email?token="+mockToken, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusFound, w.Code)
}

func TestLogoutUser(t *testing.T) {
	router := setupRouter()

	req, _ := http.NewRequest("POST", "/logout", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Logout successful")
}
