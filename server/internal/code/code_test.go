package code

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	_ = os.Setenv("LINK_URL", "https://redeem.trustwallet.com/redeem?code={{.Code}}&provider={{.Provider}}")
	_ = os.Setenv("CODE_PREFIX", "trust-")
	_ = os.Setenv("CODE_PATTERN", "####-####-####")
	_ = os.Setenv("CODE_CHARSET", "0123456789")
	config.InitConfig()
	os.Exit(m.Run())
}

func Test_getUrl(t *testing.T) {
	type args struct {
		code     string
		provider string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{"trust-VLQ-lEt-cRr", "redeem.trustwallet.com"},
			"https://redeem.trustwallet.com/redeem?code=trust-VLQ-lEt-cRr&provider=redeem.trustwallet.com",
		}, {
			"test 2",
			args{"Ha4-let-TYr", "redeem.binance.com"},
			"https://redeem.trustwallet.com/redeem?code=Ha4-let-TYr&provider=redeem.binance.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUrl(tt.args.code, tt.args.provider)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

var assets = redemption.Assets{
	Coin: 714,
	Used: false,
	Assets: []redemption.Asset{
		{
			Amount:  25000000,
			TokenId: "BUSD-BD1",
		}, {
			Amount:  10000000,
			TokenId: "BNB",
		},
	},
}

func TestCreateLinks(t *testing.T) {
	type args struct {
		count    int
		provider string
		assets   redemption.Assets
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"generate 10 codes", args{10, "redeem.trustwallet.com", assets}, false},
		{"generate 100 codes", args{100, "redeem.trustwallet.com", assets}, false},
		{"generate 1000 codes", args{1000, "redeem.trustwallet.com", assets}, false},
		{"generate code error", args{0, "redeem.trustwallet.com", assets}, true},
		{"generate code error", args{10, "redeem.trustwallet.com", redemption.Assets{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateLinks(tt.args.count, tt.args.provider, tt.args.assets)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.args.count, len(got))
			for _, gotAssets := range got {
				assert.Equal(t, tt.args.assets, gotAssets.Asset)
			}
		})
	}
}

func Test_generateCodes(t *testing.T) {
	tests := []struct {
		name  string
		count int
	}{
		{"generate 10 codes", 10},
		{"generate 100 codes", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateCodes(tt.count)
			assert.Nil(t, err)
			assert.Equal(t, tt.count, len(got))
		})
	}
}
