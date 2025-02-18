package tests

import (
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateCampaign(t *testing.T) {
	campaign := models.Campaign{
		UserID:       1,
		Title:        "Test Campaign",
		Description:  "This is a test campaign.",
		TargetAmount: 1000.00,
		Status:       "active",
		Category:     "Social Impact",
		MediaPath:    "path/to/media",
	}

	err := models.CreateCampaign(campaign)
	assert.NoError(t, err, "Expected no error while creating campaign")

}

func TestGetAllCampaigns(t *testing.T) {
	campaigns, err := models.GetAllCampaigns("", "", 0, 0)
	assert.NoError(t, err, "Expected no error while fetching all campaigns")
	assert.NotNil(t, campaigns, "Expected campaigns to be returned")
}

func TestGetCampaignById(t *testing.T) {
	campaignID := 1
	campaign, err := models.GetCampaignById(campaignID)
	assert.NoError(t, err, "Expected no error while fetching campaign by ID")
	assert.NotNil(t, campaign, "Expected campaign to be found")
}

func TestUpdateCampaign(t *testing.T) {
	campaign := models.Campaign{
		Title:        "Updated Title",
		Description:  "Updated description.",
		TargetAmount: 2000.00,
		Status:       "active",
		Category:     "Education & Research",
		MediaPath:    "path/to/updated/media",
	}

	err := models.UpdateCampaign("1", campaign)
	assert.NoError(t, err, "Expected no error while updating campaign")
}

func TestDeleteCampaign(t *testing.T) {
	campaignID := 1
	err := models.DeleteCampaign(campaignID)
	assert.NoError(t, err, "Expected no error while deleting campaign")
}

func TestSearchCampaigns(t *testing.T) {
	queryStr := "Test"
	campaigns, err := models.SearchCampaigns(queryStr)
	assert.NoError(t, err, "Expected no error while searching campaigns")
	assert.NotNil(t, campaigns, "Expected campaigns to be returned")
}

func TestGetCampaignByuser(t *testing.T) {
	userID := 1
	campaigns, err := models.GetCampaignByuser(userID)
	assert.NoError(t, err, "Expected no error while fetching campaigns by user")
	assert.NotNil(t, campaigns, "Expected campaigns to be returned")
}

func TestGetUserEmailByID(t *testing.T) {
	userID := 1
	email, err := models.GetUserEmailByID(userID)
	assert.NoError(t, err, "Expected no error while fetching user email")
	assert.NotEmpty(t, email, "Expected email to be returned")
}

func TestGetMediaByID(t *testing.T) {
	campaignID := "1"
	media, err := models.GetMediaByID(campaignID)
	assert.NoError(t, err, "Expected no error while fetching media by campaign ID")
	assert.NotEmpty(t, media, "Expected media path to be returned")
}
