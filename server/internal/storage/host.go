package storage

import (
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	HostsCollection = "hosts"
)

func (s *Storage) InsertHosts(hosts redemption.CoinHosts) (err error) {
	h := make([]interface{}, 0)
	for _, host := range hosts {
		h = append(h, host)
	}
	_, err = s.InsertMany(Database, HostsCollection, h)
	return
}

func (s *Storage) GetHosts() (redemption.CoinHosts, error) {
	hosts := make(redemption.CoinHosts, 0)

	query := bson.M{}
	var result []redemption.CoinHost
	err := s.GetValues(Database, HostsCollection, query, &result)
	if err != nil {
		return hosts, err
	}
	for _, host := range result {
		hosts = append(hosts, host)
	}
	return hosts, nil
}

func (s *Storage) GetHost(coin uint) (string, error) {
	var host redemption.CoinHost
	err := s.GetValue(Database, HostsCollection, bson.M{"coin": coin}, &host)
	return host.Host, err
}
