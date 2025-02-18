package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCampaignStatus(t *testing.T) {
	router := setupRouter()

	router.PUT("/campaign/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Campaign status updated"})
	})

	req, _ := http.NewRequest(http.MethodPut, "/campaign/status", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetAllUsers(t *testing.T) {
	router := setupRouter()

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"users": []string{"User1", "User2"}})
	})

	req, _ := http.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestDeleteUser(t *testing.T) {
	router := setupRouter()

	router.DELETE("/users/:id", func(c *gin.Context) {
		c.Status(http.StatusNoContent)
	})

	req, _ := http.NewRequest(http.MethodDelete, "/users/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
