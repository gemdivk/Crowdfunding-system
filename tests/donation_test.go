package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

type Donation struct {
	ID              int
	CampaignID      int
	UserID          int
	Amount          float64
	DonationDate    time.Time
	StripePaymentID string
}

func TestCreateDonation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	donation := Donation{
		CampaignID:      1,
		UserID:          1,
		Amount:          100.0,
		StripePaymentID: "payment_123",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "Donation" (.+) VALUES \($1, $2, $3, CURRENT_TIMESTAMP, $4) RETURNING donation_id`).
		WithArgs(donation.UserID, donation.CampaignID, donation.Amount, donation.StripePaymentID).
		WillReturnRows(sqlmock.NewRows([]string{"donation_id"}).AddRow(1))
	mock.ExpectCommit()

	err = CreateDonation(&donation)
	if err != nil {
		t.Errorf("CreateDonation() returned an error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func CreateDonation(donation *Donation) error {
	return nil
}
