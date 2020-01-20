package storage

import (
	"github.com/trustwallet/redemption/server/internal/config"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"github.com/trustwallet/redemption/server/pkg/storage/mongodb"
	"net/url"
	"strings"
)

var (
	Database = "redemption"
)

type Storage struct {
	*mongodb.MongoDb
}

func New() (*Storage, error) {
	u, err := url.Parse(config.Configuration.Mongo.Uri)
	if err == nil {
		dbName := strings.Replace(u.Path, "/", "", -1)
		if len(dbName) > 0 {
			Database = dbName
		}
	}
	mongo, err := mongodb.NewMongoDbClient(config.Configuration.Mongo.Uri)
	if err != nil {
		return nil, err
	}
	return &Storage{mongo}, nil
}

type Redeem interface {
	InsertLinks(links redemption.Links) error
	UpdateLink(link *redemption.Link) error
	GetLinks(provider string) (redemption.Links, error)
	GetLink(code string) (*redemption.Link, error)
	GetLinksByProvider(provider string) (redemption.Links, error)
}
