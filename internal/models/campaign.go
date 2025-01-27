package models

import (
	"database/sql"
	"fmt"
	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"log"
	"time"
)

type Campaign struct {
	CampaignID   int       `json:"campaign_id"`
	UserID       int       `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	TargetAmount float64   `json:"target_amount"`
	AmountRaised float64   `json:"amount_raised"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func CreateCampaign(campaign Campaign) error {
	var userExists bool
	checkUserQuery := `SELECT EXISTS(SELECT 1 FROM "User" WHERE user_id = $1)`
	err := db.DB.QueryRow(checkUserQuery, campaign.UserID).Scan(&userExists)
	if err != nil {
		log.Printf("Error checking user existence: %v", err)
		return err
	}
	if !userExists {
		log.Printf("User with ID %d does not exist.", campaign.UserID)
		return fmt.Errorf("user with ID %d does not exist", campaign.UserID)
	}

	query := `INSERT INTO "Campaign" (user_id, title, description, target_amount, status) 
              VALUES ($1, $2, $3, $4, $5) RETURNING campaign_id, created_at, updated_at`
	err = db.DB.QueryRow(query, campaign.UserID, campaign.Title, campaign.Description, campaign.TargetAmount, campaign.Status).
		Scan(&campaign.CampaignID, &campaign.CreatedAt, &campaign.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting campaign: %v", err)
		return err
	}
	return nil
}
func GetAllCampaigns() ([]Campaign, error) {
	rows, err := db.DB.Query(`SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, created_at, updated_at FROM "Campaign"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var campaign Campaign
		if err := rows.Scan(&campaign.CampaignID, &campaign.UserID, &campaign.Title, &campaign.Description,
			&campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.CreatedAt, &campaign.UpdatedAt); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}
func GetCampaignById(campaignid string) (*Campaign, error) {
	var campaign Campaign
	err := db.DB.QueryRow(`SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, created_at, updated_at 
		FROM "Campaign" WHERE campaign_id = $1`, campaignid).
		Scan(&campaign.CampaignID, &campaign.UserID, &campaign.Title, &campaign.Description, &campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.CreatedAt, &campaign.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Printf("Error retrieving campaign by ID: %v", err)
		return nil, fmt.Errorf("failed to retrieve campaign: %v", err)
	}
	return &campaign, nil
}
func UpdateCampaign(campaignID string, campaign Campaign) error {
	query := `
		UPDATE "Campaign" 
		SET title = $1, description = $2, target_amount = $3, status = $4, updated_at = CURRENT_TIMESTAMP
		WHERE campaign_id = $5`

	_, err := db.DB.Exec(query, campaign.Title, campaign.Description, campaign.TargetAmount, campaign.Status, campaignID)
	if err != nil {
		log.Printf("Error updating campaign: %v", err)
		return fmt.Errorf("failed to update campaign: %v", err)
	}

	return nil
}
func DeleteCampaign(campaignID string) error {
	query := `
		DELETE FROM "Campaign" WHERE campaign_id = $1`

	_, err := db.DB.Exec(query, campaignID)
	if err != nil {
		log.Printf("Error deleting campaign: %v", err)
		return fmt.Errorf("failed to delete campaign: %v", err)
	}

	return nil
}
