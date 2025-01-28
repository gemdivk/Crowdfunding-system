package handlers

import (
	"net/http"
	"strconv"

	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
)

// --- User Points Handlers ---

func GetAllUserPoints(c *gin.Context) {
	points, err := models.GetAllUserPoints()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user points"})
		return
	}
	c.JSON(http.StatusOK, points)
}

func AddUserPoints(c *gin.Context) {
	var newPoints models.UserPoints
	if err := c.ShouldBindJSON(&newPoints); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.AddUserPoints(newPoints); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user points"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User points added successfully"})
}

func UpdateUserPoints(c *gin.Context) {
	userID := c.Param("user_id")
	var data struct {
		Points int `json:"points"`
	}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.UpdateUserPoints(userID, data.Points); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user points"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User points updated successfully"})
}

func DeleteUserPoints(c *gin.Context) {
	userID := c.Param("user_id")
	if err := models.DeleteUserPoints(userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user points"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User points deleted successfully"})
}

// --- Achievements Handlers ---

func GetAllAchievements(c *gin.Context) {
	achievements, err := models.GetAllAchievements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch achievements"})
		return
	}
	c.JSON(http.StatusOK, achievements)
}

func AddAchievement(c *gin.Context) {
	var newAchievement models.Achievement
	if err := c.ShouldBindJSON(&newAchievement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.AddAchievement(newAchievement); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add achievement"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Achievement added successfully"})
}

func DeleteAchievement(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid achievement ID"})
		return
	}

	if err := models.DeleteAchievement(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete achievement"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Achievement deleted successfully"})
}
