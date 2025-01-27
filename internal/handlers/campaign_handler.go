package handlers

import (
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCampaignHandler(c *gin.Context) {
	var campaign models.Campaign
	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := models.CreateCampaign(campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create campaign"})
		return
	}

	c.JSON(http.StatusCreated, campaign)
}

// GetCampaignsHandler handles retrieving all campaigns
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
	var campaign models.Campaign
	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	campaignID := c.Param("id") // Get the campaign ID from the URL parameter

	err := models.UpdateCampaign(campaignID, campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update campaign"})
		return
	}

	c.JSON(http.StatusOK, campaign)
}
func DeleteCampaignHandler(c *gin.Context) {
	campaignID := c.Param("id")

	err := models.DeleteCampaign(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete campaign"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
