package binance

import (
	sdk "github.com/binance-chain/go-sdk/client"
	"github.com/binance-chain/go-sdk/client/transaction"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/binance-chain/go-sdk/types/msg"
	"github.com/spf13/viper"
	"github.com/trustwallet/blockatlas/coin"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/logger"
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/pkg/redemption"
)

const (
	CurrentNetwork = types.ProdNetwork
)

type Platform struct {
	Client     sdk.DexClient
	KeyManager keys.KeyManager
	Address    string
}

func (p *Platform) Init() error {
	provider := viper.GetString("binance.rpc")
	mnemonic := config.Configuration.Wallet.Mnemonic

	types.Network = CurrentNetwork

	var err error
	p.KeyManager, err = keys.NewMnemonicKeyManager(mnemonic)
	if err != nil || p.KeyManager == nil {
		return errors.E(err, "unable to create a NewMnemonicKeyManager")
	}
	p.Address = p.KeyManager.GetAddr().String()
	p.Client, err = sdk.NewDexClient(provider, CurrentNetwork, p.KeyManager)
	if err != nil {
		return errors.E(err, "cannot connect to client", logger.Params{"provider": provider})
	}
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.BNB]
}

func (p *Platform) TransferAssets(addresses []string, assets redemption.Assets) (string, error) {
	logParams := logger.Params{"to": addresses, "assets": assets}
	logger.Info("sending assets...", logParams)

	logger.Info("getting main account balance...", logger.Params{"address": p.Address})
	account, err := p.Client.GetAccount(p.Address)
	if err != nil {
		return "", errors.E(err, "failed to get main account balance", logger.Params{"address": p.Address})
	}

	// Create transfer objects
	transfers := make([]msg.Transfer, 0)
	for _, asset := range assets.Assets {
		if !verifyBalance(account, asset, len(addresses)) {
			return "", errors.E("main account doesn't have enough balance",
				logger.Params{"asset": asset})
		}

		for _, address := range addresses {
			addr, err := types.GetFromBech32(address, CurrentNetwork.Bech32Prefixes())
			if err != nil {
				logger.Info(err, "AccAddressFromBech32 decode failed", logger.Params{"address": address})
				continue
			}
			transfers = append(transfers, msg.Transfer{
				ToAddr: addr,
				Coins: types.Coins{
					types.Coin{
						Denom:  asset.TokenId,
						Amount: asset.Amount,
					},
				},
			})
		}
	}

	// send multi-transfer
	sendResult, err := p.Client.SendToken(transfers, true, transaction.WithMemo(config.Configuration.Transaction.Memo))
	if err != nil || !sendResult.Ok {
		return "", errors.E(err, "failed to send transactions", logParams, logger.Params{"result": sendResult})
	}
	logger.Info("txs sent!", logger.Params{"result": sendResult.Hash, "log": sendResult.Log, "Code": sendResult.Code})
	return sendResult.TxCommitResult.Hash, nil
}

func (p *Platform) GetPublicAddress() (string, error) {
	types.Network = CurrentNetwork
	mnemonic := config.Configuration.Wallet.Mnemonic
	km, err := keys.NewMnemonicKeyManager(mnemonic)
	if err != nil || km == nil {
		return "", errors.E(err, "unable to create a NewMnemonicKeyManager")
	}
	return km.GetAddr().String(), nil
}

func verifyBalance(account *types.BalanceAccount, asset redemption.Asset, repeat int) bool {
	var tokenBalance int64 = 0
	for _, balance := range account.Balances {
		if balance.Symbol == asset.TokenId {
			balanceLog := logger.Params{"balance": balance.Free, "asset": asset}
			logger.Info("main account balance", balanceLog)
			tokenBalance = balance.Free.ToInt64()
			break
		}
	}
	return tokenBalance >= (asset.Amount * int64(repeat))
}
