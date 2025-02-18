package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDonationCreation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.POST("/campaigns/:id/donate", handlers.CreateDonation)

	userID := 1
	router.Use(func(c *gin.Context) {
		c.Set("userID", userID)
	})

	tests := []struct {
		campaignID       string
		requestBody      interface{}
		expectedStatus   int
		expectedResponse map[string]interface{}
	}{
		{
			campaignID: "1",
			requestBody: map[string]interface{}{
				"amount":            100.0,
				"stripe_payment_id": "test_payment_id",
			},
			expectedStatus: http.StatusCreated,
			expectedResponse: map[string]interface{}{
				"message": "Donation successful",
			},
		},
		{
			campaignID: "invalid",
			requestBody: map[string]interface{}{
				"amount":            100.0,
				"stripe_payment_id": "test_payment_id",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"error": "Invalid campaign ID",
			},
		},
		{
			campaignID: "1",
			requestBody: map[string]interface{}{
				"amount":            100.0,
				"stripe_payment_id": "",
			},
			expectedStatus: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"error": "Stripe payment ID is required",
			},
		},
	}

	for _, test := range tests {
		body, _ := json.Marshal(test.requestBody)
		req, _ := http.NewRequest(http.MethodPost, "/campaigns/"+test.campaignID+"/donate", bytes.NewBuffer(body))
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, test.expectedStatus, res.Code)
		var response map[string]interface{}
		json.Unmarshal(res.Body.Bytes(), &response)
		for key, expectedValue := range test.expectedResponse {
			assert.Equal(t, expectedValue, response[key])
		}
	}
}

func TestMyDonationsHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/users/:userID/donations", handlers.MyDonationsHandler)

	tests := []struct {
		userID         string
		expectedStatus int
	}{
		{
			userID:         "1",
			expectedStatus: http.StatusOK,
		},
		{
			userID:         "invalid",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest(http.MethodGet, "/users/"+test.userID+"/donations", nil)
		res := httptest.NewRecorder()

		router.ServeHTTP(res, req)

		assert.Equal(t, test.expectedStatus, res.Code)
	}
}
