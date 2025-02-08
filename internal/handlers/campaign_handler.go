package handlers

import (
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func CreateCampaignHandler(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Please log in first"})
		return
	}
	var campaign models.Campaign
	campaign.UserID = userID.(int)

	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := models.CreateCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create campaign"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Campaign created successfully",
		"campaign": gin.H{
			"campaign_id":   campaign.CampaignID,
			"user_id":       campaign.UserID,
			"title":         campaign.Title,
			"description":   campaign.Description,
			"target_amount": campaign.TargetAmount,
			"status":        campaign.Status,
			"category":      campaign.Category, // Ensure category is part of the response
			"created_at":    campaign.CreatedAt,
			"updated_at":    campaign.UpdatedAt,
		},
	})
}

func GetCampaignsHandler(c *gin.Context) {
	category := c.DefaultQuery("category", "")
	campaigns, err := models.GetAllCampaigns(category)
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
	if err := c.BindJSON(&campaign); err != nil {
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
		"campaign": campaign, // Returning the created campaign details
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
		c.JSON(http.StatusForbidden, gin.H{"error": "you don't have any permission do it"})
		return
	}
	campaigns, err := models.GetCampaignByuser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaigns"})
		return
	}

	c.JSON(http.StatusOK, campaigns)
}
