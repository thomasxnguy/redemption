package api

import (
	"github.com/chenjiandongx/ginprom"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/trustwallet/blockatlas/pkg/ginutils"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/blockatlas/pkg/metrics"
	"github.com/trustwallet/blockatlas/util"
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

// @Summary Insert coin host
// @ID insert_coin_host
// @Description Insert a host for a specific coin
// @Accept json
// @Produce json
// @Tags host
// @Header 200 {string} Authorization "Bearer test"
// @Param hosts body redemption.CoinHosts true "Hosts"
// @Success 200 {object} redemption.Success
// @Error 500 {object} ginutils.ApiError
// @Router /v1/hosts [put]
func insertCoinHosts(storage storage.Host) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var body redemption.CoinHosts
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		if len(body) == 0 {
			ginutils.ErrorResponse(c).Message("invalid payload").Render()
			return
		}

		err := storage.InsertHosts(body)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		ginutils.RenderSuccess(c, redemption.Success{Status: true})
	}
}

// @Summary Get coin host
// @ID get_coin_host
// @Description Get a host for a specific coin
// @Accept json
// @Produce json
// @Tags host
// @Header 200 {string} Authorization "Bearer test"
// @Param page query string true "Page" default(0)
// @Success 200 {object} redemption.CoinHosts
// @Error 500 {object} ginutils.ApiError
// @Router /v1/hosts [get]
func getCoinHosts(storage storage.Host) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "0")
		p, err := strconv.Atoi(page)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message("failed to parse the page number").Render()
			return
		}
		links, err := storage.GetHosts(p + 1)
		if err != nil {
			logger.Error(err)
		}
		ginutils.RenderSuccess(c, links)
	}
}

// @Summary Create code for referral
// @ID create_links
// @Description Create code and links for referral from a specific asset
// @Accept json
// @Produce json
// @Tags redeem
// @Header 200 {string} Authorization "Bearer test"
// @Param links body redemption.CreateLinks true "Links"
// @Success 200 {object} redemption.Links
// @Error 500 {object} ginutils.ApiError
// @Router /v1/links/create [post]
func createLinks(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		var body redemption.CreateLinks
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}
		if body.Assets == nil || len(body.Assets.Assets) == 0 {
			ginutils.ErrorResponse(c).Message("invalid assets").Render()
			return
		}

		links, err := code.CreateLinks(body.LinkCount, body.Provider, *body.Assets)
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
// @Header 200 {string} Authorization "Bearer test"
// @Param page query string true "Page" default(0)
// @Param provider query string true "Provider name"
// @Success 200 {object} redemption.Links
// @Error 500 {object} ginutils.ApiError
// @Router /v1/links [get]
func getAllLinks(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "0")
		provider := c.Query("provider")
		p, err := strconv.Atoi(page)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message("failed to parse the page number").Render()
			return
		}
		links, err := storage.GetLinks(p+1, provider)
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
// @Header 200 {string} Authorization "Bearer test"
// @Param code path string true "the link code"
// @Success 200 {object} redemption.Link
// @Error 500 {object} ginutils.ApiError
// @Router /v1/link/:code [get]
func getLink(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		redeemCode := c.Param("code")
		link, err := storage.GetLink(redeemCode)
		if err != nil {
			logger.Error(err)
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
// @Header 200 {string} Authorization "Bearer test"
// @Param redeem body redemption.Redeem true "Redeem"
// @Success 200 {object} redemption.RedeemResult
// @Error 500 {object} ginutils.ApiError
// @Router /v1/referral/redeem [post]
func redeemCode(storage storage.Redeem) func(c *gin.Context) {
	if storage == nil {
		return nil
	}

	// avoid race condition with semaphore
	semaphore := util.NewSemaphore(1)
	return func(c *gin.Context) {
		var body redemption.Redeem
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		// Get the code from database
		link, err := storage.GetLink(body.Code)
		if err != nil || !link.Valid {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message("invalid the code").Render()
			return
		}
		semaphore.Acquire()
		defer semaphore.Release()

		host, err := storage.GetHost(link.Asset.Coin)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message("coin without node").Render()
			return
		}

		// Get asset platform
		p, err := platform.GetTxPlatform(link.Asset.Coin, host)
		if err != nil {
			logger.Error(err, "failed to initialize platform API")
			ginutils.ErrorResponse(c).Message(err.Error()).Render()
			return
		}

		// Invalidate the link before to redeem the balance, to avoid
		// race condition and prevent hackers from explorer bugs
		link.Valid = false
		err = storage.UpdateLink(link)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message("cannot invalidate code").Render()
			return
		}

		// Verify asset is used
		if link.Asset.Used {
			ginutils.ErrorResponse(c).Message("asset already used").Render()
			return
		}
		link.Asset.Used = true

		result := make([]string, 0)
		observers := body.Observers.GetCoinObservers(link.Asset.Coin)
		for _, observer := range observers {
			to := append(observer.Addresses, observer.PublicKeys...)
			hash, err := p.TransferAssets(to, link.Asset)
			if err != nil {
				continue
			}
			result = append(result, hash)
		}

		// Return success
		ginutils.RenderSuccess(c, createRedeemSuccessResponse(result, link.Asset.Assets, p.Coin().Decimals))

		// Save assets state
		err = storage.UpdateLink(link)
		if err != nil {
			logger.Error(err)
			ginutils.ErrorResponse(c).Message("cannot invalidate the asset").Render()
			return
		}
	}
}

func createRedeemSuccessResponse(result []string, assets []*redemption.Asset, decimals uint) redemption.RedeemResult {
	msg, err := message.GetMessage()
	if err != nil {
		return redemption.RedeemResult{
			Type: redemption.RedeemResultTypeSuccess,
		}
	}
	return redemption.RedeemResult{
		Type:        redemption.RedeemResultTypeSuccess,
		Title:       msg.GetTitle(),
		Description: msg.GetDescription(assets, decimals),
		ImageURL:    msg.GetImage(),
		ResultId:    result,
	}
}
