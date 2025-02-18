package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCampaign(t *testing.T) {
	client := resty.New()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Ik1pa2FlbEBleGFtcGxlLmNvbSIsIlVzZXJJRCI6MzYsImV4cCI6MTczOTI5OTUwMH0.0D2KDBWBZ1NB4iipkM1AchlJ5r9_JzIMEBUTS38bis0"

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+token).
		SetBody(`{"title": "Test Campaign", "goal": 1000}`).
		Post("http://localhost:8080/campaigns")
	assert.NoError(t, err)
	assert.Equal(t, 201, resp.StatusCode())
}
