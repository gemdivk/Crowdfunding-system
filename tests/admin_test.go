package tests

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// GetTotalUsers fetches the total user count from the database.
func GetTotalUsers(db *sql.DB) int {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	err := row.Scan(&count)
	if err != nil {
		return 0
	}
	return count
}

func TestGetTotalUsers(t *testing.T) {
	// Initialize mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error initializing sqlmock: %v", err)
	}
	defer db.Close()

	// Define expected query and result
	rows := sqlmock.NewRows([]string{"count"}).AddRow(10)
	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users").WillReturnRows(rows)

	// Call function
	result := GetTotalUsers(db)

	// Validate result
	if result != 10 {
		t.Errorf("Expected 10 users, got %d", result)
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
