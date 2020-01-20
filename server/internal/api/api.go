package api

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	_ "github.com/trustwallet/redemption/server/docs"
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/internal/storage"
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

func reverseProxy(c *gin.Context) {
	c.Request.Header.Del("Forwarded")
	c.Request.Header.Del("X-Forwarded-Proto")
	c.Request.Header.Del("X-Forwarded-Host")
	c.Request.Header.Del("X-Forwarded-For")
}

func addMiddleware(engine *gin.Engine) {
	sg := sentrygin.New(sentrygin.Options{})
	engine.Use(reverseProxy, sg)
}

func makeRoutes(engine *gin.Engine, storage *storage.Storage) {
	makeMetricsRoute(engine)
	engine.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/status", statusHandler)

	// Serve frontend static files
	engine.Use(static.Serve("/", static.LocalFile(config.Configuration.Client.Path, true)))

	dashboard := engine.Group("/v1")
	dashboard.Use(ginutils.TokenAuthMiddleware(config.Configuration.Dashboard.Token))

	// Address
	dashboard.GET("/address/:platform", getPublicAddress())

	// Dashboard
	dashboard.GET("/links", getAllLinks(storage))
	dashboard.GET("/link/:code", getLink(storage))
	dashboard.POST("/link/:code", updateLink(storage))
	dashboard.POST("/links/create", createLinks(storage))

	// Redeem
	api := engine.Group("/v1")
	api.Use(ginutils.TokenAuthMiddleware(config.Configuration.Api.Token))
	api.POST("/referral/redeem", redeemCode(storage))
}
