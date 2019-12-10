package platform

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"github.com/trustwallet/redemption/server/platform/binance"
)

var Platforms map[uint]redemption.Platform

var platformList = []redemption.Platform{
	&binance.Platform{},
}

var TxAPIs map[uint]redemption.TxApi

func Init() {
	Platforms = make(map[uint]redemption.Platform)
	TxAPIs = make(map[uint]redemption.TxApi)
	for _, platform := range platformList {
		if _, exists := Platforms[platform.Coin().ID]; exists {
			logger.Fatal("Duplicate platform")
		}

		if txApi, ok := platform.(redemption.TxApi); ok {
			TxAPIs[platform.Coin().ID] = txApi
		}
	}
}

func GetTxPlatform(coin uint, provider string) (redemption.TxApi, error) {
	p, ok := TxAPIs[coin]
	if !ok {
		return nil, errors.E("coin not supported", errors.Params{"coin": coin})
	}
	err := p.Init(provider)
	if err != nil {
		return nil, errors.E(err, "failed to initialize platform API", errors.Params{"coin": coin, "provider": provider})
	}
	return p, nil
}
