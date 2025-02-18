package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/handlers"
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/campaigns", handlers.CreateCampaignHandler)
	router.GET("/campaigns", handlers.GetCampaignsHandler)
	router.GET("/campaigns/:id", handlers.GetCampaignId)
	router.PUT("/campaigns/:id", handlers.UpdateCampaignHandler)
	router.DELETE("/campaigns/:id", handlers.DeleteCampaignHandler)
	return router
}

func TestCreateCampaignHandler(t *testing.T) {
	router := createRouter()

	userID := 1
	router.Use(func(c *gin.Context) {
		c.Set("userID", userID)
	})

	campaign := models.Campaign{Title: "New Campaign", Description: "This is a test campaign."}
	jsonData, _ := json.Marshal(campaign)

	req, _ := http.NewRequest("POST", "/campaigns", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusCreated, response.Code)

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	assert.True(t, res["success"].(bool))
	assert.Equal(t, "Campaign created successfully", res["message"])
}

func TestGetCampaignsHandler(t *testing.T) {
	router := createRouter()

	req, _ := http.NewRequest("GET", "/campaigns", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

}

func TestGetCampaignId(t *testing.T) {
	router := createRouter()

	req, _ := http.NewRequest("GET", "/campaigns/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var campaign models.Campaign
	json.NewDecoder(response.Body).Decode(&campaign)
	assert.Equal(t, "Expected Campaign Title", campaign.Title)
}

func TestUpdateCampaignHandler(t *testing.T) {
	router := createRouter()

	userID := 1
	router.Use(func(c *gin.Context) {
		c.Set("userID", userID)
	})

	campaign := models.Campaign{Title: "Updated Campaign", Description: "Updated description"}
	jsonData, _ := json.Marshal(campaign)

	req, _ := http.NewRequest("PUT", "/campaigns/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	assert.True(t, res["success"].(bool))
	assert.Equal(t, "Campaign updated successfully", res["message"])
}

func TestDeleteCampaignHandler(t *testing.T) {
	router := createRouter()

	userID := 1
	router.Use(func(c *gin.Context) {
		c.Set("userID", userID)
	})

	req, _ := http.NewRequest("DELETE", "/campaigns/1", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	assert.Equal(t, http.StatusOK, response.Code)

	var res map[string]interface{}
	json.NewDecoder(response.Body).Decode(&res)
	assert.Equal(t, "Campaign deleted successfully", res["message"])
}
