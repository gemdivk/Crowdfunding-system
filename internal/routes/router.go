package routes

import (
	"net/http"

	"github.com/gemdivk/Crowdfunding-system/internal/handlers"
	"github.com/gemdivk/Crowdfunding-system/internal/middleware"
	"github.com/gemdivk/Crowdfunding-system/internal/social"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Static("/static", "./frontend")
	router.Static("/uploads", "./uploads")
	router.POST("/upload", handlers.UploadFileHandler)
	// Route to serve the HTML file
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.POST("/create-payment-intent", handlers.CreatePaymentIntent)

	// Load HTML templates (if you have multiple .html files)
	router.LoadHTMLGlob("frontend/*.html")
	// Campaign routes
	campaignRoutes := router.Group("/campaigns")
	{
		campaignRoutes.Use(middleware.AuthMiddleware())
		campaignRoutes.POST("/", handlers.CreateCampaignHandler) // Create a new campaign
		campaignRoutes.PUT("/:id", handlers.UpdateCampaignHandler)
		campaignRoutes.DELETE("/:id", handlers.DeleteCampaignHandler)
		campaignRoutes.GET("/user/:id", handlers.GetCampaignsbyUser)
	}
	router.GET("/campaigns/search", handlers.SearchCampaignsHandler)
	router.GET("/campaigns/", handlers.GetCampaignsHandler)
	router.GET("/campaigns/:id", handlers.GetCampaignId)
	router.POST("/register", handlers.RegisterUser)
	router.GET("/verify-email", handlers.VerifyEmail)
	router.POST("/login", handlers.LoginUser)
	router.POST("/logout", handlers.LogoutUser)

	donationRoutes := router.Group("/campaigns/:id/donations")
	{
		donationRoutes.Use(middleware.AuthMiddleware())
		donationRoutes.POST("/", handlers.CreateDonation) // Donate to a campaign
		donationRoutes.GET("/", handlers.GetDonationsByCampaign)
	}
	donation := router.Group("/donations")
	{
		donation.Use(middleware.AuthMiddleware())
		donation.GET("/my/:userID", handlers.MyDonationsHandler)
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
		publicURL := c.DefaultQuery("url", "https://yourcrowdfundingurl.com")
		text := c.DefaultQuery("text", "Check out this campaign!")

		// For demo, using fixed campaign data; in practice, pass actual campaign info
		campaignTitle := "Campaign Title"
		campaignDescription := "This campaign is for raising funds to protect the Amazon rainforest."

		facebookLink := social.GetFacebookShareLink(publicURL, campaignDescription)
		twitterLink := social.GetTwitterShareLink(publicURL, text)
		linkedinLink := social.GetLinkedInShareLink(publicURL, campaignTitle, campaignDescription)

		c.JSON(http.StatusOK, gin.H{
			"facebook": facebookLink,
			"twitter":  twitterLink,
			"linkedin": linkedinLink,
		})
	})

	router.GET("/gamification/leaderboard", handlers.GetLeaderboard)
	router.POST("/gamification/update", middleware.AuthMiddleware(), handlers.UpdateUserPoints)

	router.GET("/admin", func(c *gin.Context) {
		c.HTML(http.StatusOK, "admin.html", nil)
	})

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware())

	setupAdminRoutes(router)

	return router
}

func setupAdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.AuthMiddleware())
	{
		adminRoutes.GET("/users", handlers.GetUsersHandler)
		adminRoutes.DELETE("/users/:id", handlers.DeleteUserHandler)
		adminRoutes.GET("/campaigns", handlers.GetCampaignsHandler)
		adminRoutes.DELETE("/campaigns/:id", handlers.DeleteCampaignHandler)
		adminRoutes.PUT("/campaigns/:id/status", handlers.UpdateCampaignStatusHandler)
		adminRoutes.GET("/dashboard", handlers.GetAdminDashboardHandler)
		adminRoutes.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "admin.html", nil)
		})
	}
}
