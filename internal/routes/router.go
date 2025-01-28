package routes

import (
	"github.com/gemdivk/Crowdfunding-system/internal/handlers"
	"github.com/gemdivk/Crowdfunding-system/internal/middleware"
	"github.com/gemdivk/Crowdfunding-system/internal/social"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.POST("/logout", handlers.LogoutUser)

	donationRoutes := router.Group("/campaigns/:id/donations")
	{
		donationRoutes.POST("/", handlers.CreateDonation) // Donate to a campaign
		donationRoutes.GET("/", handlers.GetDonationsByCampaign)
	}

	router.GET("/donations/user/:user_id", handlers.GetDonationsByUser)
	//	router.PUT("/donations/:id", handlers.UpdateDonation)
	//	router.DELETE("/donations/:id", handlers.DeleteDonation)
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		protectedRoutes.PUT("/donations/:id", handlers.UpdateDonation)
		protectedRoutes.DELETE("/donations/:id", handlers.DeleteDonation)
	}

	router.GET("/share", func(c *gin.Context) {
		targetURL := c.DefaultQuery("url", "https://yourcrowdfundingurl.com")
		text := c.DefaultQuery("text", "Check out this campaign!")

		facebookLink := social.GetFacebookShareLink(targetURL)
		twitterLink := social.GetTwitterShareLink(targetURL, text)
		linkedinLink := social.GetLinkedInShareLink(targetURL, "Campaign Title", "A description of the campaign")

		c.JSON(http.StatusOK, gin.H{
			"facebook": facebookLink,
			"twitter":  twitterLink,
			"linkedin": linkedinLink,
		})
	})

	return router
}
