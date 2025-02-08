package models

import (
	"database/sql"
	"fmt"
	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"log"
	"time"
)

type Campaign struct {
	CampaignID   int       `json:"campaign_id" form:"campaign_id"`
	UserID       int       `json:"user_id" form:"user_id"`
	Title        string    `json:"title" form:"title"`
	Description  string    `json:"description" form:"description"`
	TargetAmount float64   `json:"target_amount" form:"target_amount"`
	AmountRaised float64   `json:"amount_raised" form:"amount_raised"`
	Status       string    `json:"status" form:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MediaPath    string    `json:"media_path"`
}

func CreateCampaign(campaign Campaign) error {
	// Check if the user exists.
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

	// Insert the new campaign into the database.
	query := `INSERT INTO "Campaign" (user_id, title, description, target_amount, status, media_path) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING campaign_id, created_at, updated_at`
	err = db.DB.QueryRow(query,
		campaign.UserID,
		campaign.Title,
		campaign.Description,
		campaign.TargetAmount,
		campaign.Status,
		campaign.MediaPath).
		Scan(&campaign.CampaignID, &campaign.CreatedAt, &campaign.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting campaign: %v", err)
		return err
	}
	return nil
}

func GetAllCampaigns() ([]Campaign, error) {
	rows, err := db.DB.Query(`SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, media_path, created_at, updated_at FROM "Campaign"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var campaign Campaign
		if err := rows.Scan(&campaign.CampaignID, &campaign.UserID, &campaign.Title, &campaign.Description,
			&campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.MediaPath, &campaign.CreatedAt, &campaign.UpdatedAt); err != nil {
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
		Scan(&campaign.CampaignID, &campaign.UserID, &campaign.Title, &campaign.Description, &campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.MediaPath, &campaign.CreatedAt, &campaign.UpdatedAt)
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
		SET title = $1, description = $2, target_amount = $3, status = $4, media_path = $5, updated_at = CURRENT_TIMESTAMP
		WHERE campaign_id = $6`
	_, err := db.DB.Exec(query,
		campaign.Title,
		campaign.Description,
		campaign.TargetAmount,
		campaign.Status,
		campaign.MediaPath,
		campaignID)
	if err != nil {
		log.Printf("Error updating campaign: %v", err)
		return fmt.Errorf("failed to update campaign: %v", err)
	}
	return nil
}

func DeleteCampaign(campaignID string) error {
	query := `DELETE FROM "Campaign" WHERE campaign_id = $1`
	_, err := db.DB.Exec(query, campaignID)
	if err != nil {
		log.Printf("Error deleting campaign: %v", err)
		return fmt.Errorf("failed to delete campaign: %v", err)
	}
	return nil
}

func SearchCampaigns(queryStr string) ([]Campaign, error) {
	sqlQuery := `
		SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, created_at, updated_at
		FROM "Campaign"
		WHERE title ILIKE $1 OR status ILIKE $1 OR user_id::text ILIKE $1`
	rows, err := db.DB.Query(sqlQuery, "%"+queryStr+"%")
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

func GetCampaignByuser(userid any) ([]Campaign, error) {
	query := `SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, created_at, updated_at, media_path 
              FROM "Campaign" WHERE user_id = $1`
	rows, err := db.DB.Query(query, userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var campaign Campaign
		if err := rows.Scan(&campaign.CampaignID, &campaign.UserID, &campaign.Title, &campaign.Description,
			&campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.CreatedAt, &campaign.UpdatedAt, &campaign.MediaPath); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}
	return campaigns, nil
}
