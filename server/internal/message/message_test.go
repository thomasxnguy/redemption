package message

import (
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"testing"
)

var asset1 = redemption.Asset{
	Amount:  100000000,
	TokenId: "BNB",
}

var asset2 = redemption.Asset{
	Amount:  2000000000,
	TokenId: "BUSD",
}

var asset3 = redemption.Asset{
	Amount:  3000000,
	TokenId: "ONE",
}

var asset4 = redemption.Asset{
	Amount:  556000,
	TokenId: "XRP",
}

func Test_message_GetDescription(t *testing.T) {
	description := "You have redeemed {{.Values}}"
	decimals := uint(8)
	type fields struct {
		Description string
	}
	tests := []struct {
		name     string
		fields   fields
		assets   []redemption.Asset
		decimals uint
		want     string
	}{
		{
			"test 1 coin",
			fields{description},
			[]redemption.Asset{asset1},
			decimals,
			"You have redeemed 1 BNB",
		}, {
			"test 2 coins",
			fields{description},
			[]redemption.Asset{asset1, asset2},
			decimals,
			"You have redeemed 1 BNB and 20 BUSD",
		}, {
			"test 3 coins",
			fields{description},
			[]redemption.Asset{asset1, asset2, asset3},
			decimals,
			"You have redeemed 1 BNB, 20 BUSD and 0.03 ONE",
		}, {
			"test 4 coins",
			fields{description},
			[]redemption.Asset{asset1, asset2, asset3, asset4},
			decimals,
			"You have redeemed 1 BNB, 20 BUSD, 0.03 ONE and 0.00556 XRP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := message{
				Description: tt.fields.Description,
			}
			if got := m.GetDescription(tt.assets, tt.decimals); got != tt.want {
				t.Errorf("GetDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceValues(t *testing.T) {
	text := "You have redeemed {{.Values}}"
	type args struct {
		text   string
		values string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1 coin",
			args{text, "1 BNB"},
			"You have redeemed 1 BNB",
		},
		{
			"test 2 coins",
			args{text, "20 BUSD and 1 BNB"},
			"You have redeemed 20 BUSD and 1 BNB",
		},
		{
			"test 3 coins",
			args{text, "20 BUSD, 30 ONE and 1 BNB"},
			"You have redeemed 20 BUSD, 30 ONE and 1 BNB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := replaceValues(tt.args.text, tt.args.values)
			if got != tt.want {
				t.Errorf("replaceValues() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_floatValue(t *testing.T) {
	tests := []struct {
		name     string
		volume   int64
		decimals uint
		want     float64
	}{
		{"float value 1", 10000, 8, 0.0001},
		{"float value 2", 2500, 8, 0.000025},
		{"float value 3", 9, 4, 0.0009},
		{"float value 4", 900003330, 3, 900003.33},
		{"float value 5", 343, 10, 0.0000000343},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := floatValue(tt.volume, tt.decimals); got != tt.want {
				t.Errorf("removeDecimals() = %v, want %v", got, tt.want)
			}
		})
	}
}
