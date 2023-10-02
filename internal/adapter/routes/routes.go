package routes

import (
	docs "walls-user-service/docs"
	"walls-user-service/internal/adapter/api"
	configuration "walls-user-service/internal/core/helper/configuration-helper"
	errorhelper "walls-user-service/internal/core/helper/error-helper"
	logger "walls-user-service/internal/core/helper/log-helper"
	message "walls-user-service/internal/core/helper/message-helper"
	"walls-user-service/internal/core/middleware"
	services "walls-user-service/internal/core/services"
	ports "walls-user-service/internal/port"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userRepository ports.UserRepository, redisClient *redis.Client) *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies(nil)

	userService := services.NewUserService(userRepository, redisClient)

	handler := api.NewHTTPHandler(userService)

	logger.LogEvent("INFO", "Configuring Routes!")
	router.Use(middleware.LogRequest)

	corrs_config := cors.DefaultConfig()
	corrs_config.AllowAllOrigins = true

	router.Use(cors.New(corrs_config))
	//router.Use(middleware.SetHeaders)

	docs.SwaggerInfo.Description = "Walls User Service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = configuration.ServiceConfiguration.ServiceName

	router.POST("/api/user", handler.CreateUser)
	router.POST("/api/user/:user_reference/company", handler.CreateCompanyProfile)
	router.POST("/api/user/:user_reference/company/walls-badge", handler.CreateCompanyWallsBadge)
	router.POST("/api/user/:user_reference/user/walls-badge", handler.CreateUserWallsBadge)
	router.PUT("/api/user/:user_reference/company/:company_reference", handler.UpdateCompanyProfile)
	router.PUT("/api/user/:user_reference/company/:company_reference/walls-badge/:walls_badge_reference/disable", handler.DisableCompanyWallsBadge)
	router.PUT("/api/user/:user_reference/walls-badge/:walls_badge_reference/disable", handler.DisableUserWallsBadge)
	router.GET("/api/user/:user_reference/company/:company_reference/walls-badges", handler.GetCompanyWallsBadgeList)
	router.GET("/api/user/:user_reference/walls-badges", handler.GetUserWallsBadgeList)
	router.GET("/api/user/:user_reference/company/:company_reference/walls-badge/default", handler.GetDefaultCompanyWallsBadge)
	router.GET("/api/user/:user_reference/walls-badge/default", handler.GetDefaultUserWallsBadge)
	router.PUT("/api/user/:user_reference/company/:company_reference/disable", handler.DisableCompanyProfile)
	router.PUT("/api/user/:user_reference/company/:company_reference/logo", handler.UpdateCompanyLogo)
	router.PUT("/api/user/:user_reference/email/update-status", handler.UpdateUserProfileEmailStatus)
	// router.PUT("/api/user/:user_reference/phone/update-status", handler.UpdateUserProfilePhoneStatus)
	router.PUT("/api/user/:user_reference/company/:company_reference/email/update-status", handler.UpdateCompanyProfileEmailStatus)
	router.PUT("/api/user/:user_reference/bank/:bank_reference/set-default", handler.SetDefaultBank)
	router.PUT("/api/user/:user_reference/card/:card_reference/set-default", handler.SetDefaultCard)
	router.PUT("/api/user/:user_reference/name", handler.UpdateUserName)
	router.PUT("/api/user/:user_reference/email", handler.UpdateEmail)
	router.PUT("/api/user/:user_reference/dob", handler.UpdateDateOfBirth)
	router.PUT("/api/user/:user_reference/address", handler.UpdateAddress)
	router.PUT("/api/user/:user_reference/photos", handler.UpdatePhoto)
	router.PUT("/api/user/:user_reference/wallet", handler.UpdateWallet)
	router.POST("/api/user/:user_reference/bank", handler.AddBank)
	router.PUT("/api/user/:user_reference/bank/:bank_reference", handler.UpdateBank)
	router.POST("/api/user/:user_reference/card", handler.AddCard)
	router.PUT("/api/user/:user_reference/card/:card_reference", handler.UpdateCard)
	router.PUT("/api/user/:user_reference/notification-options", handler.UpdateNotificationOptions)
	router.PUT("/api/user/:user_reference/device", handler.UpdateDevice)
	router.POST("/api/user/:user_reference/identification", handler.AddDocumentation)
	router.PUT("/api/user/:user_reference/identification/:identification_reference", handler.UpdateDocumentation)
	router.POST("/api/user/:user_reference/contact", handler.AddContact)
	router.GET("/api/user/:user_reference", handler.GetUserByReference)
	router.GET("/api/user/phone/:phone", handler.GetUserByPhone)
	router.GET("/api/user/walls-tag/:wallsTag", handler.GetUserByWallsTag)
	router.GET("/api/user/walls-badge-reference/:walls_badge_reference", handler.GetUserByWallsBagdeReference)
	router.POST("/api/user/device", handler.GetUserByDevice)
	router.PUT("/api/user/:user_reference/balance", handler.UpdateBalance)
	router.PUT("/api/user/:user_reference/tier", handler.UpdateTier)
	router.PUT("/api/user/:user_reference/coupon", handler.AddCoupon)
	router.PUT("/api/user/:user_reference/reward", handler.UpdateRewards)
	router.PUT("/api/user/:user_reference/enable", handler.EnableUser)
	router.PUT("/api/user/:user_reference/disable", handler.DisableUser)
	router.POST("/api/user/send-otp", handler.CreateOtpRequest)
	router.POST("/api/user/validate-otp", handler.ValidateOtpRequest)
	router.GET("/api/user/user-reference", handler.CreateUserReference)
	router.GET("/api/user/document-reference", handler.CreateDocumentReference)
	router.POST("/api/user/create-identity", handler.CreateIdentityRequest)
	router.POST("/api/user/upgrade-tier", handler.UpgradeTierRequest)
	router.POST("/api/user/transaction", handler.CreateTransactionRequest)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			errorhelper.ErrorMessage(errorhelper.NoResourceError, message.NoResourceFound))
	})

	return router
}
