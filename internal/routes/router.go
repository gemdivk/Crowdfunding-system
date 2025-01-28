package routes

import (
	"github.com/gemdivk/Crowdfunding-system/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Campaign routes
	campaignRoutes := router.Group("/campaigns")
	{
		campaignRoutes.POST("/", handlers.CreateCampaignHandler) // Create a new campaign
		campaignRoutes.GET("/", handlers.GetCampaignsHandler)    // Get all campaigns
		campaignRoutes.GET("/:id", handlers.GetCampaignId)
		campaignRoutes.PUT("/:id", handlers.UpdateCampaignHandler)
		campaignRoutes.DELETE("/:id", handlers.DeleteCampaignHandler)

	}
	router.POST("/register", handlers.RegisterUser)
	router.POST("/login", handlers.LoginUser)

	// Donation routes
	donationRoutes := router.Group("/campaigns/:id/donations")
	{
		donationRoutes.POST("/", handlers.CreateDonation)        // Donate to a campaign
		donationRoutes.GET("/", handlers.GetDonationsByCampaign) // Get all donations for a campaign
	}

	router.GET("/donations/user/:user_id", handlers.GetDonationsByUser)
	router.PUT("/donations/:id", handlers.UpdateDonation)
	router.DELETE("/donations/:id", handlers.DeleteDonation)
	return router
}
