// handlers/donations.go
package handlers

import (
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateDonation allows users to donate to a campaign
func CreateDonation(c *gin.Context) {
	campaignID := c.Param("id") // Get the campaign ID from URL parameters
	var donation models.Donation
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	donation.CampaignID = campaignID // Set the campaign ID

	// Call CreateDonation function to insert it into the database
	if err := models.CreateDonation(&donation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process donation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "Donation successful",
		"donation_id": donation.ID,
		"amount":      donation.Amount,
		"campaign_id": donation.CampaignID,
	})
}

// GetDonationsByCampaign retrieves all donations for a specific campaign
func GetDonationsByCampaign(c *gin.Context) {
	campaignID := c.Param("id")
	donations, err := models.GetDonationsForCampaign(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch donations for campaign"})
		return
	}

	c.JSON(http.StatusOK, donations)
}

// GetDonationsByUser retrieves all donations made by a specific user
func GetDonationsByUser(c *gin.Context) {
	userID := c.Param("user_id")
	donations, err := models.GetDonationsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch donations for user"})
		return
	}

	c.JSON(http.StatusOK, donations)
}

// UpdateDonation allows updating donation details
func UpdateDonation(c *gin.Context) {
	donationID := c.Param("id")
	var updatedDonation models.Donation
	if err := c.ShouldBindJSON(&updatedDonation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.UpdateDonation(donationID, &updatedDonation); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update donation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Donation updated successfully",
		"donation_id": donationID,
	})
}

// DeleteDonation allows users or admins to delete a donation
func DeleteDonation(c *gin.Context) {
	donationID := c.Param("id")
	if err := models.DeleteDonation(donationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete donation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Donation deleted successfully",
		"donation_id": donationID,
	})
}
