package handlers

import (
	"fmt"
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateCampaignHandler(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in first"})
		return
	}

	var campaign models.Campaign
	if err := c.ShouldBind(&campaign); err != nil {
		log.Printf("Binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Override the UserID from the form with the authenticated user's ID.
	if uid, ok := userIDValue.(int); ok {
		campaign.UserID = uid
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID in context"})
		return
	}

	// Continue with file upload processing...
	file, err := c.FormFile("media")
	if err != nil && err.Error() != "http: no such file" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload file"})
		return
	}
	if file != nil {
		filePath := fmt.Sprintf("./uploads/%s", file.Filename)
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
			return
		}
		campaign.MediaPath = filePath
	}

	if err := models.CreateCampaign(campaign); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create campaign"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success":  true,
		"message":  "Campaign created successfully",
		"campaign": campaign,
	})
}

func GetCampaignsHandler(c *gin.Context) {
	campaigns, err := models.GetAllCampaigns()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaigns"})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func GetCampaignId(c *gin.Context) {
	campaignID := c.Param("id")
	campaign, err := models.GetCampaignById(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaign"})
		return
	}

	if campaign == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func UpdateCampaignHandler(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in first"})
		return
	}

	var campaign models.Campaign
	// Use ShouldBind to allow binding from multipart or JSON
	if err := c.ShouldBind(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	campaignID := c.Param("id") // Get the campaign ID from the URL parameter
	err := models.UpdateCampaign(campaignID, campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success":  true,
		"message":  "Campaign updated successfully",
		"campaign": campaign,
	})
}

func DeleteCampaignHandler(c *gin.Context) {
	_, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in first"})
		return
	}
	campaignID := c.Param("id")

	err := models.DeleteCampaign(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}

func SearchCampaignsHandler(c *gin.Context) {
	// Get search query from request URL (e.g., ?query=keyword)
	query := c.DefaultQuery("query", "")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	// Call the model function to search campaigns
	campaigns, err := models.SearchCampaigns(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search campaigns"})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}

func GetCampaignsbyUser(c *gin.Context) {
	userid, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in first"})
		return
	}
	userIDParam := c.Param("id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		log.Println("Invalid user ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if userid != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to do that"})
		return
	}
	campaigns, err := models.GetCampaignByuser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaigns"})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}
