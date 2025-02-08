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
	Category     string    `json:"category"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

var allowedCategories = map[string]bool{
	"Charity": true, "Health": true, "Education": true, "Environment": true,
	"Animal Welfare": true, "Film": true, "Food Security": true, "Games": true,
	"Journalism": true, "Music": true, "Photography": true, "Publishing": true,
	"Technology": true, "Discover": true,
}

func IsValidCategory(category string) bool {
	return allowedCategories[category]
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

	query := `INSERT INTO "Campaign" (user_id, title, description, target_amount, status, category) 
              VALUES ($1, $2, $3, $4, $5, $6) RETURNING campaign_id, created_at, updated_at`
	err = db.DB.QueryRow(query, campaign.UserID, campaign.Title, campaign.Description, campaign.TargetAmount, campaign.Status, campaign.Category).
		Scan(&campaign.CampaignID, &campaign.CreatedAt, &campaign.UpdatedAt)
	if err != nil {
		log.Printf("Error inserting campaign: %v", err)
		return err
	}
	return nil
}
func GetAllCampaigns(category string) ([]Campaign, error) {
	// Base query to select all campaigns
	query := `SELECT campaign_id, user_id, title, description, target_amount,amount_raised, status,category, created_at, updated_at FROM "Campaign"`
	var rows *sql.Rows
	var err error

	// If category is not empty, add a WHERE clause for filtering by category
	if category != "" {
		query += ` WHERE category = $1`
		rows, err = db.DB.Query(query, category)
	} else {
		rows, err = db.DB.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var campaigns []Campaign
	for rows.Next() {
		var campaign Campaign
		if err := rows.Scan(&campaign.CampaignID, &campaign.UserID, &campaign.Title, &campaign.Description,
			&campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.Category, &campaign.CreatedAt, &campaign.UpdatedAt); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}
	log.Printf("Returning campaigns: %+v", campaigns)
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
func SearchCampaigns(query string) ([]Campaign, error) {
	// Create a SQL query that searches by title, status, or user ID
	sqlQuery := `
		SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, created_at, updated_at
		FROM "Campaign"
		WHERE title ILIKE $1 OR status ILIKE $1 OR user_id::text ILIKE $1
	`

	// Perform the query
	rows, err := db.DB.Query(sqlQuery, "%"+query+"%")
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
	query := `SELECT campaign_id, user_id, title, description, target_amount, amount_raised, status, created_at, updated_at 
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
			&campaign.TargetAmount, &campaign.AmountRaised, &campaign.Status, &campaign.CreatedAt, &campaign.UpdatedAt); err != nil {
			return nil, err
		}
		campaigns = append(campaigns, campaign)
	}

	return campaigns, nil
}
