package tests

import (
	"os"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")

	db.InitDB()

	assert.NotNil(t, db.DB, "DB should not be nil after initialization")
}

func TestCheckDBConnection(t *testing.T) {
	os.Setenv("DB_USER", "testuser")
	os.Setenv("DB_PASSWORD", "testpassword")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_NAME", "testdb")

	db.InitDB()

	err := db.CheckDBConnection()
	assert.NoError(t, err, "DB connection should be healthy")
}
