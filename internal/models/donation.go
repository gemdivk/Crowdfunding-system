// models/donation.go
package models

import (
	"fmt"
	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"time"
)

type Donation struct {
	ID         int       `json:"id"`
	CampaignID string    `json:"campaign_id"`
	UserID     string    `json:"user_id"`
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// CreateDonation saves a new donation in the database
// CreateDonation saves a new donation in the database
func CreateDonation(donation *Donation) error {
	query := `INSERT INTO "Donation" (user_id, campaign_id, amount, donation_date) 
          VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING donation_id`

	err := db.DB.QueryRow(query, donation.UserID, donation.CampaignID, donation.Amount).
		Scan(&donation.ID) // PostgreSQL will return the auto-generated donation_id here
	if err != nil {
		return fmt.Errorf("Failed to insert donation: %v", err)
	}
	return nil
}

// GetDonationsForCampaign fetches all donations for a specific campaign
func GetDonationsForCampaign(campaignID string) ([]Donation, error) {
	var donations []Donation
	query := `SELECT donation_id, user_id, campaign_id, amount, donation_date 
              FROM "Donation" WHERE campaign_id = $1`
	rows, err := db.DB.Query(query, campaignID)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch donations for campaign: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var donation Donation
		if err := rows.Scan(&donation.ID, &donation.UserID, &donation.CampaignID, &donation.Amount, &donation.CreatedAt); err != nil {
			return nil, fmt.Errorf("Failed to scan donation: %v", err)
		}
		donations = append(donations, donation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return donations, nil
}

// GetDonationsByUser fetches all donations made by a specific user
func GetDonationsByUser(userID string) ([]Donation, error) {
	var donations []Donation
	query := `SELECT donation_id, user_id, campaign_id, amount, donation_date 
              FROM "Donation" WHERE user_id = $1`
	rows, err := db.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch donations for user: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var donation Donation
		if err := rows.Scan(&donation.ID, &donation.UserID, &donation.CampaignID, &donation.Amount, &donation.CreatedAt); err != nil {
			return nil, fmt.Errorf("Failed to scan donation: %v", err)
		}
		donations = append(donations, donation)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return donations, nil
}

// UpdateDonation updates a specific donation in the database
func UpdateDonation(id string, updatedDonation *Donation) error {
	query := `UPDATE "Donation" SET amount = $1, donation_date = CURRENT_TIMESTAMP 
          WHERE donation_id = $2`

	_, err := db.DB.Exec(query, updatedDonation.Amount, id)
	if err != nil {
		return fmt.Errorf("Failed to update donation: %v", err)
	}
	return nil
}

// DeleteDonation deletes a specific donation from the database
func DeleteDonation(id string) error {
	query := `DELETE FROM "Donation" WHERE donation_id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete donation: %v", err)
	}
	return nil
}
