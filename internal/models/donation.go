package models

import (
	"database/sql"
	"fmt"
	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"time"
)

type Donation struct {
	ID         int       `json:"id"`
	CampaignID int       `json:"campaign_id"` // Changed to int
	UserID     int       `json:"user_id"`     // Changed to int
	Amount     float64   `json:"amount"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func CreateDonation(donation *Donation) error {
	query := `INSERT INTO "Donation" (user_id, campaign_id, amount, donation_date) 
          VALUES ($1, $2, $3, CURRENT_TIMESTAMP) RETURNING donation_id`

	err := db.DB.QueryRow(query, donation.UserID, donation.CampaignID, donation.Amount).
		Scan(&donation.ID)
	if err != nil {
		return fmt.Errorf("Failed to insert donation: %v", err)
	}
	return nil
}

func GetDonationsForCampaign(campaignID int) ([]Donation, error) {
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

func GetUserByDonationID(donationID int) (int, error) {
	var userID int
	query := `SELECT user_id from "Donation" where donation_id = $1`
	err := db.DB.QueryRow(query, donationID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("no donation found for donation_id: %d", donationID)
		}
		return 0, err
	}
	return userID, nil
}

func GetDonationsByUser(userID int) ([]Donation, error) {
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

func UpdateDonation(id int, updatedDonation *Donation) error { // id is now int
	query := `UPDATE "Donation" SET amount = $1, donation_date = CURRENT_TIMESTAMP 
          WHERE donation_id = $2`
	_, err := db.DB.Exec(query, updatedDonation.Amount, id)
	if err != nil {
		return fmt.Errorf("Failed to update donation: %v", err)
	}
	return nil
}

func DeleteDonation(id int) error {

	query := `DELETE FROM "Donation" WHERE donation_id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to delete donation: %v", err)
	}
	return nil
}
