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
	/*	donationRoutes := router.Group("/donations")
		{
			donationRoutes.POST("/:campaign_id", handlers.DonateToCampaignHandler) // Donate to a campaign
			donationRoutes.GET("/:campaign_id", handlers.GetDonationsHandler)      // Get all donations for a campaign
		}
	*/
	return router
}
