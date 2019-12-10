package api

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/internal/storage"

	swaggerFiles "github.com/swaggo/files"
	_ "github.com/trustwallet/redemption/server/docs"
)

func Provide(storage *storage.Storage) {
	gin.SetMode(config.Configuration.Api.Mode)
	engine := gin.Default()

	addMiddleware(engine)
	makeRoutes(engine, storage)

	port := config.Configuration.Port
	logger.Info("Running application", logger.Params{"port": port})
	if err := engine.Run(":" + port); err != nil {
		logger.Fatal("Application failed", err)
	}
}

func addMiddleware(engine *gin.Engine) {
	sg := sentrygin.New(sentrygin.Options{})
	engine.Use(sg)
}

// @title TrustWallet Redeem API
// @version 1.0
// @description Provide a redemption API

// @contact.name Trust Wallet
// @contact.url https://t.me/wallecore

// @license.name MIT License
// @license.url https://raw.githubusercontent.com/trustwallet/redemption/master/LICENSE
func makeRoutes(engine *gin.Engine, storage *storage.Storage) {
	makeMetricsRoute(engine)
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/status", statusHandler)

	// Serve frontend static files
	engine.Use(static.Serve("/", static.LocalFile("../client/build", true)))

	v1 := engine.Group("/v1")
	v1.Use(ginutils.TokenAuthMiddleware(config.Configuration.Api.Auth_Token))

	// Redeem
	v1.GET("/links", getAllLinks(storage))
	v1.GET("/link/:code", getLink(storage))
	v1.POST("/links/create", createLinks(storage))
	v1.POST("/referral/redeem", redeemCode(storage))

	// Hosts
	v1.PUT("/hosts", insertCoinHosts(storage))
	v1.GET("/hosts", getCoinHosts(storage))
}
