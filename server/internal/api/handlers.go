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
	"fmt"
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
// @Param Authorization header string false "Bearer Token" default()
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
// @Param Authorization header string false "Bearer Token" default()
// @Success 200 {object} redemption.CoinHosts
// @Error 500 {object} ginutils.ApiError
// @Router /v1/hosts [get]
func getCoinHosts(storage storage.Host) func(c *gin.Context) {
	if storage == nil {
		return nil
	}
	return func(c *gin.Context) {
		links, err := storage.GetHosts()
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
	semaphore := util.NewSemaphore(1)
	return func(c *gin.Context) {
		reqBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.nopCloser(bytes.NewReader(reqBody))
		fmt.Println(string(reqBody))

		var body redemption.Redeem
		if err := c.BindJSON(&body); err != nil {
			logger.Error(err)
			ginutils.RenderSuccess(c, createRedeemErrorResponse(err.Error()))
			return
		}

		fmt.Println(body)
		fmt.Println("Code:", body.Code)
		fmt.Println("Observers:", body.Observers)

		// Get the code from database
		link, err := storage.GetLink(body.Code)
		if err != nil || !link.Valid || link.IsOutdated() {
			logger.Error(err)
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Invalid code"))
			return
		}
		semaphore.Acquire()
		defer semaphore.Release()

		host, err := storage.GetHost(link.Asset.Coin)
		if err != nil {
			logger.Error(err)
			ginutils.RenderSuccess(c, createRedeemErrorResponse("Coin without node. You need to insert a host for this coin node"))
			return
		}

		// Get asset platform
		p, err := platform.GetTxPlatform(link.Asset.Coin, host)
		if err != nil {
			logger.Error(err, "Invalid platform API")
			ginutils.RenderSuccess(c, createRedeemErrorResponse(err.Error()))
			return
		}

		// Invalidate the link before to redeem the balance, to avoid
		// race condition and prevent hackers from explorer bugs
		link.Valid = false
		err = storage.UpdateLink(link)
		if err != nil {
			logger.Error(err)
			ginutils.RenderSuccess(c, createRedeemErrorResponse("CCannot invalidate code"))
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
			return
		}
	}
}

func createRedeemSuccessResponse(result []string, assets []redemption.Asset, decimals uint) redemption.RedeemResult {
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
