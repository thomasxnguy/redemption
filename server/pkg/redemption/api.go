package redemption

import (
	"github.com/trustwallet/blockatlas/coin"
)

type Platform interface {
	Init(provider string) error
	Coin() coin.Coin
	GetPublicAddress() (string, error)
}

type TxApi interface {
	Platform
	TransferAssets(addresses []string, assets Assets) (string, error)
}
