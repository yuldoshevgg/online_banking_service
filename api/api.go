package api

import (
	"online_banking_service/api/handlers"
	config "online_banking_service/config"

	"online_banking_service/api/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/tylfin/gin-swagger-files"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @title Swagger Example API
// SetUpRouter godoc
// @description Online Banking Service
// @termsOfService
// @version 1.0
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Version = cfg.Version
	docs.SwaggerInfo.Schemes = []string{cfg.HTTPScheme}

	r.Use(customCORSMiddleware())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/login", h.Login)
	r.POST("/sign-in", h.SignIn)

	r.Use(h.AuthMiddleware)
	{
		r.POST("/create-account", h.CreateAccount)
		r.PUT("/pay-for-account", h.PayForAccount)
		r.PUT("/withdraw-from-account", h.WithdrawFromAccount)
		r.GET("/accounts", h.GetAccounts)
		r.POST("/transfer-balance", h.TransferBalance)
	}

	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
