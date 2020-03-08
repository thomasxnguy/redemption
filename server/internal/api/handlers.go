package api

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/metrics"
	"github.com/trustwallet/blockatlas/pkg/semaphore"
	"github.com/trustwallet/redemption/server/internal/code"
	"github.com/trustwallet/redemption/server/internal/message"
	"github.com/trustwallet/redemption/server/internal/storage"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"github.com/trustwallet/redemption/server/platform"
	"strconv"
)

// @Summary Get Status
// @ID health
// @Description Get application status
// @Tags status
// @Success 200 {object} redemption.Success
// @Router /status [get]
func statusHandler(c *gin.Context) {
	ginutils.RenderSuccess(c, redemption.Success{Status: true})
}

// @Summary Get Metrics
// @ID metrics
// @Description Get application metrics
// @Tags status
// @Router /metrics [get]
func makeMetricsRoute(router gin.IRouter) {
	router.Use(metrics.PromMiddleware())
	m := router.Group("/metrics")
	m.GET("/", ginprom.PromHandler(promhttp.Handler()))
}

// @Summary Create code for referral
// @ID create_links
// @Description Create code and links for referral from a specific asset
// @Accept json
// @Produce json
// @Tags redeem
// @Param Authorization header string false "Bearer Token" default()
// @Param links body redemption.CreateLinks true "Links"
// @Success 200 {object} redemption.Links
// @Error 500 {object} ginutils.ApiError
// @Router /v1/links/create [post]
func createLinks(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		body := redemption.CreateLinks{LinkCount: 50}
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		if len(body.Provider) == 0 || len(body.Asset.Assets) == 0 {
			ginutils.ErrorResponse(c).Message("invalid payload").Render()
			return
		}

		links, err := code.CreateLinks(body.LinkCount, body.Provider, body.Asset)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		err = storage.InsertLinks(links)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, links)
	}
}

// @Summary Get all links
// @ID get_all_links
// @Description Get all referral links
// @Accept json
// @Produce json
// @Tags redeem
// @Param Authorization header string false "Bearer Token" default()
// @Param provider query string true "Provider name"
// @Success 200 {object} redemption.Links
// @Error 500 {object} ginutils.ApiError
// @Router /v1/links [get]
func getAllLinks(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		provider := c.Query("provider")
		links, err := storage.GetLinks(provider)
		if err != nil {
			logger.Error(err)
		}
		ginutils.RenderSuccess(c, links)
	}
}

// @Summary Get specific link
// @ID get_link
// @Description Get a specific link
// @Accept json
// @Produce json
// @Tags redeem
// @Param Authorization header string false "Bearer Token" default()
// @Param code path string true "the link code"
// @Success 200 {object} redemption.Link
// @Error 500 {object} ginutils.ApiError
// @Router /v1/link/{code} [get]
func getLink(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		redeemCode := c.Param("code")
		link, err := storage.GetLink(redeemCode)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, link)
	}
}

// @Summary Get public address
// @ID public_address
// @Description Get public address from a main wallet
// @Accept json
// @Produce json
// @Tags account
// @Param Authorization header string false "Bearer Token" default()
// @Param platform path string true "the platform Id" default(714)
// @Success 200 {object} redemption.Address
// @Error 500 {object} ginutils.ApiError
// @Router /v1/address/{platform} [get]
func getPublicAddress() func(c *gin.Context) {
	return func(c *gin.Context) {
		platformId := c.Param("platform")
		if len(platformId) == 0 {
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Platform cannot be empty"))
			return
		}
		id, err := strconv.Atoi(platformId)
		if err != nil {
			logger.Error(err, "Invalid platform id")
			ginutils.RenderSuccess(c, createRedeemErrorResponse(err.Error()))
			return
		}

		p, err := platform.GetPlatform(uint(id))
		if err != nil {
			logger.Error(err, "Invalid platform API")
			ginutils.RenderSuccess(c, createRedeemErrorResponse(err.Error()))
			return
		}

		address, err := p.GetPublicAddress()
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, redemption.Address{Address: address, Coin: p.Coin().ID, CoinName: p.Coin().Name})
	}
}

// @Summary Update specific link
// @ID update_link
// @Description Update a specific link
// @Accept json
// @Produce json
// @Tags redeem
// @Param Authorization header string false "Bearer Token" default()
// @Param code path string true "the link code"
// @Param link body redemption.UpdateLink true "Link"
// @Success 200 {object} redemption.Link
// @Error 500 {object} ginutils.ApiError
// @Router /v1/link/{code} [post]
func updateLink(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		redeemCode := c.Param("code")
		link, err := storage.GetLink(redeemCode)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		var body redemption.UpdateLink
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		link.MergeLinks(body)
		err = storage.UpdateLink(link)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, link)
	}
}

// @Summary Redeem code
// @ID create_links
// @Description Redeem the referral code
// @Accept json
// @Produce json
// @Tags redeem
// @Param Authorization header string false "Bearer Token" default()
// @Param redeem body redemption.Redeem true "Redeem"
// @Success 200 {object} redemption.RedeemResult
// @Error 500 {object} redemption.RedeemResult
// @Router /v1/referral/redeem [post]
func redeemCode(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}

	// avoid race condition with semaphore
	semaphore := semaphore.NewSemaphore(1)
	return func(c *gin.Context) {
		var body redemption.Redeem
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.RenderSuccess(c, createRedeemErrorResponse(err.Error()))
			return
		}
		if body.Observers == nil || len(body.Observers) == 0 {
			ginutils.RenderSuccess(c, createRedeemErrorResponse("No accounts available to redeem"))
			return
		}

		// Get the code from database
		link, err := storage.GetLink(body.Code)
		if err != nil {
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Invalid code"))
			return
		}
		if !link.Valid {
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Code already been redeemed"))
			return
		}
		if link.IsOutdated() {
			ginutils.RenderSuccess(c, createRedeemErrorResponse("The code is outdated"))
			return
		}

		semaphore.Acquire()
		defer semaphore.Release()

		// Get asset platform
		p, err := platform.GetPlatform(link.Asset.Coin)
		if err != nil {
			ginutils.RenderSuccess(c, createRedeemErrorResponse(err.Error()))
			return
		}

		// Invalidate the link before to redeem the balance, to avoid
		// race condition and prevent hackers from explorer bugs
		link.Valid = false
		err = storage.UpdateLink(link)
		if err != nil {
			logger.Error("Cannot update link before redemption", err)
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Cannot invalidate code"))
			return
		}

		// Verify asset is used
		if link.Asset.Used {
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Asset already used"))
			return
		}
		link.Asset.Used = true

		result := make([]string, 0)
		observers := body.Observers.GetCoinObservers(link.Asset.Coin)
		for _, observer := range observers {
			if len(observer.Addresses) == 0 {
				continue
			}
			hash, err := p.TransferAssets([]string{observer.Addresses[0]}, link.Asset)
			if err != nil {
				logger.Error(err)
				continue
			}
			result = append(result, hash)
		}

		// Return success
		ginutils.RenderSuccess(c, createRedeemSuccessResponse(result, link.Asset, p.Coin().Decimals))

		// Save assets state
		err = storage.UpdateLink(link)
		if err != nil {
			logger.Error("Cannot update link after redemption", err)
			return
		}
	}
}

func createRedeemSuccessResponse(result []string, assets redemption.Assets, decimals uint) redemption.RedeemResult {
	msg, err := message.GetMessage()
	if err != nil {
		return redemption.RedeemResult{
			Type: redemption.RedeemResultTypeSuccess,
		}
	}
	return redemption.RedeemResult{
		Type:        redemption.RedeemResultTypeSuccess,
		Title:       msg.GetTitle(),
		Description: msg.GetDescription(assets.Assets, decimals),
		ImageURL:    msg.GetImage(),
		Assets:      assets,
		ResultId:    result,
	}
}

func createRedeemErrorResponse(description string) redemption.RedeemResult {
	result := redemption.RedeemResult{
		Type:        redemption.RedeemResultTypeError,
		Title:       "Failed Redeem",
		Description: description,
	}
	msg, err := message.GetMessage()
	if err == nil {
		result.ImageURL = msg.GetImage()
	}
	return result
}
