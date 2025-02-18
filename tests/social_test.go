package tests

import (
	"net/url"
	"testing"

	"github.com/gemdivk/Crowdfunding-system/internal/social" // Import the package
)

func TestGetFacebookShareLink(t *testing.T) {
	targetURL := "https://example.com"
	description := "Check this out!"

	expected := "https://www.facebook.com/sharer/sharer.php?u=" + url.QueryEscape(targetURL) + "&quote=" + url.QueryEscape(description)
	result := social.GetFacebookShareLink(targetURL, description) // Call the function with package prefix

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
