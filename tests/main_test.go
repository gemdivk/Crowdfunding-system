package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/db"
	"github.com/gemdivk/Crowdfunding-system/internal/routes"
	"github.com/joho/godotenv"
)

func TestMainFunctionality(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Fatalf("Failed to load .env file: %v", err)
	}

	db.InitDB()
	db.CheckDBConnection()

	router := routes.SetupRouter()

	ts := httptest.NewServer(router)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}
}
