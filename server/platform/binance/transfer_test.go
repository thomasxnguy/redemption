package binance

import (
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"testing"
)

var account = &types.BalanceAccount{
	Number:  707542,
	Address: "tbnb1qxm48ndhmh7su0r7zgwmwkltuqgly57jdf8yf8",
	Balances: []types.TokenBalance{
		{
			Symbol: "BNB",
			Free:   19979110000,
		},
		{
			Symbol: "USDT.B-B7C",
			Free:   309113030,
		},
	},
	PublicKey: []uint8{2, 51, 243, 103, 161, 196, 20, 67, 40, 236, 165, 206, 5, 149, 118, 185, 195, 214, 254, 249, 161, 171, 72, 86, 40, 149, 17, 98, 60, 71, 41, 233, 231},
	Sequence:  16,
	Flags:     0,
}

var asset1 = redemption.Asset{
	Amount:  100000000,
	TokenId: "BNB",
}

var asset2 = redemption.Asset{
	Amount:  200000,
	TokenId: "USDT.B-B7C",
}

var asset3 = redemption.Asset{
	Amount:  3000000,
	TokenId: "ONE",
}

func Test_verifyBalance(t *testing.T) {
	type args struct {
		account *types.BalanceAccount
		asset   redemption.Asset
		repeat  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"asset1 repeat 1 time", args{account, asset1, 1}, true},
		{"asset2 repeat 1 time", args{account, asset2, 1}, true},
		{"asset1 repeat 3 times", args{account, asset1, 3}, true},
		{"asset2 repeat 3 times", args{account, asset2, 3}, true},
		{"invalid value", args{account, asset1, 30000}, false},
		{"invalid asset", args{account, asset3, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyBalance(tt.args.account, tt.args.asset, tt.args.repeat); got != tt.want {
				t.Errorf("verifyBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
