package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var jwtKey = []byte("your_secret_key")

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.New()

	router.Use(middleware.AuthMiddleware())

	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	})

	tests := []struct {
		name         string
		authHeader   string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "No Authorization Header",
			authHeader:   "",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"Authorization header is required"}`,
		},
		{
			name:         "Invalid Header Format",
			authHeader:   "InvalidHeaderFormat",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"Authorization header must be Bearer token"}`,
		},
		{
			name:         "Missing Token",
			authHeader:   "Bearer ",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"Token is missing"}`,
		},
		{
			name:         "Invalid Token",
			authHeader:   "Bearer invalidtoken",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"Invalid or expired token"}`,
		},
		{
			name: "Valid Token",
			authHeader: func() string {
				claims := jwt.MapClaims{
					"userID": 1,
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, _ := token.SignedString(jwtKey)
				return "Bearer " + tokenString
			}(),
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"success"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}

			resp := httptest.NewRecorder()
			router.ServeHTTP(resp, req)

			assert.Equal(t, tt.expectedCode, resp.Code)
			assert.JSONEq(t, tt.expectedBody, resp.Body.String())
		})
	}
}
