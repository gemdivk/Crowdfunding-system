package handlers

import (
	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateDonation(c *gin.Context) {
	campaignIDStr := c.Param("id")
	campaignID, err := strconv.Atoi(campaignIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	var donation models.Donation
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	donation.CampaignID = campaignID

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

func GetDonationsByCampaign(c *gin.Context) {
	campaignIDStr := c.Param("id")
	campaignID, err := strconv.Atoi(campaignIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}
	donations, err := models.GetDonationsForCampaign(campaignID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch donations for campaign"})
		return
	}

	c.JSON(http.StatusOK, donations)
}

func GetDonationsByUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.Atoi(userIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	donations, err := models.GetDonationsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch donations for user"})
		return
	}

	c.JSON(http.StatusOK, donations)
}

func UpdateDonation(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated "})
		return
	}

	donationIDStr := c.Param("id")
	donationID, err := strconv.Atoi(donationIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid donation ID"})
		return
	}

	donorid, err1 := models.GetUserByDonationID(donationID)
	if err1 == nil {
		if donorid != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have a permission to update it. Sorry xD"})
			return
		}
	}

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
		"donor_id":    donorid,
		"user_id":     userID,
	})
}

func DeleteDonation(c *gin.Context) {

	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}
	donationIDStr := c.Param("id")
	donationID, err := strconv.Atoi(donationIDStr) // Convert string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid donation ID"})
		return
	}

	var donorID int
	donorID, err = models.GetUserByDonationID(donationID)
	if err == nil {
		if donorID != userID {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to delete it"})
			return
		}
	}

	if err := models.DeleteDonation(donationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete donation"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Donation deleted successfully",
		"donation_id": donationID,
		"user_id":     userID,
		"donor_id":    donorID,
	})
}
