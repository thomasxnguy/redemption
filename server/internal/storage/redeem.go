package storage

import (
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/redemption/server/pkg/redemption"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	LinksCollection = "links"
)

func (s *Storage) InsertLinks(links redemption.Links) (err error) {
	l := make([]interface{}, 0)
	for _, link := range links {
		l = append(l, link)
	}
	_, err = s.InsertMany(Database, LinksCollection, l)
	return
}

func (s *Storage) UpdateLink(link *redemption.Link) error {
	_, err := s.Update(Database, LinksCollection, link, bson.D{{"code", link.Code}})
	if err != nil {
		return errors.E(err, "cannot update links")
	}
	return nil
}

func (s *Storage) GetLinks(page int, provider string) (redemption.Links, error) {
	links := make(redemption.Links, 0)

	query := bson.M{}
	if len(provider) > 0 {
		query = bson.M{"provider": provider}
	}
	var result []*redemption.Link
	err := s.GetValues(Database, LinksCollection, query, &result)
	if err != nil {
		return links, err
	}
	count := redemption.LinksPageCount + 1
	pagination := page * count
	for i, link := range result {
		if i > (pagination - 1) {
			return links, nil
		}
		if i < pagination-count {
			continue
		}
		links = append(links, link)
	}
	return links, nil
}

func (s *Storage) GetLink(code string) (result *redemption.Link, err error) {
	err = s.GetValue(Database, LinksCollection, bson.M{"code": code}, &result)
	return
}

func (s *Storage) GetLinksByProvider(provider string) (redemption.Links, error) {
	var result []*redemption.Link
	err := s.GetValues(Database, LinksCollection, bson.M{"provider": provider}, &result)
	if err != nil {
		return nil, err
	}
	links := make(redemption.Links, 0)
	for _, link := range result {
		links = append(links, link)
	}
	return links, nil
}
