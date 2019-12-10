package mongodb

import (
	"github.com/stretchr/testify/assert"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

var bsonAsset1 = bson.M{
	"coin":     714,
	"token_id": "BUSD-BD1",
	"amount":   25000000,
}

var bsonAsset2 = bson.M{
	"coin":     714,
	"token_id": "BNB",
	"amount":   10000000,
}

var bsonInvalid = bson.M{
	"assets":     nil,
	"link_count": 50,
	"provider":   "block.trust.com",
}

var asset1 = redemption.Asset{
	Amount:  25000000,
	TokenId: "BUSD-BD1",
}

var asset2 = redemption.Asset{
	Amount:  10000000,
	TokenId: "BNB",
}

func Test_getResult(t *testing.T) {
	type args struct {
		doc    []interface{}
		result interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"test asset 1",
			args{doc: []interface{}{bsonAsset1}, result: &[]redemption.Asset{}},
			&[]redemption.Asset{asset1},
			false,
		}, {
			"test asset 2",
			args{doc: []interface{}{bsonAsset2}, result: &[]redemption.Asset{}},
			&[]redemption.Asset{asset2},
			false,
		}, {
			"test 2 assets",
			args{doc: []interface{}{bsonAsset1, bsonAsset2}, result: &[]redemption.Asset{}},
			&[]redemption.Asset{asset1, asset2},
			false,
		}, {
			"want error",
			args{doc: []interface{}{bsonAsset1}, result: &[]redemption.Asset{}},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := getResult(tt.args.doc, tt.args.result)
			if tt.wantErr {
				assert.NotNil(t, tt.args.result)
				return
			}
			assert.Nil(t, err)
			assert.Equal(t, tt.want, tt.args.result)
		})
	}
}

func TestNewMongoDbClientError(t *testing.T) {
	tests := []struct {
		name string
		uri  string
	}{
		{"error", "..trust.."},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoDbClient(tt.uri)
			assert.NotNil(t, err)
			assert.Nil(t, got)
		})
	}
}
