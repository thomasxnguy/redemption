package mongodb

import (
	"context"
	"encoding/json"
	"github.com/trustwallet/blockatlas/pkg/errors"
	"github.com/trustwallet/blockatlas/pkg/storage/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDb struct {
	*mongo.Client
	host string
}

func NewMongoDbClient(uri string) (*MongoDb, error) {
	clientOptions := options.Client().ApplyURI(uri).SetRetryWrites(false)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, errors.E(err, "Cannot connect to MongoDb")
	}
	defer cancel()
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.E(err, "MongoDb connection test failed")
	}
	return &MongoDb{
		Client: client,
		host:   uri,
	}, nil
}

func (db *MongoDb) GetValue(database, collection string, query bson.M, result interface{}) error {
	c := db.Database(database).Collection(collection)
	err := c.FindOne(context.TODO(), query).Decode(result)
	if err != nil {
		return errors.E(err, util.ErrNotFound,
			errors.Params{"database": database, "collection": collection, "query": query})
	}
	return nil
}

func (db *MongoDb) GetValues(database, collection string, query bson.M, result interface{}) error {
	c := db.Database(database).Collection(collection)
	ctx := context.TODO()
	cur, err := c.Find(ctx, query)
	if err != nil {
		return errors.E(err, util.ErrNotFound,
			errors.Params{"database": database, "collection": collection, "query": query})
	}
	defer cur.Close(ctx)

	values := make([]interface{}, 0)
	for cur.Next(ctx) {
		var value bson.M
		err := cur.Decode(&value)
		if err != nil {
			continue
		}
		values = append(values, value)
	}
	if err := cur.Err(); err != nil {
		return errors.E(err, util.ErrNotFound,
			errors.Params{"database": database, "collection": collection, "query": query})
	}

	return getResult(values, result)
}

func (db *MongoDb) InsertMany(database, collection string, value []interface{}) ([]interface{}, error) {
	c := db.Database(database).Collection(collection)
	res, err := c.InsertMany(context.TODO(), value)
	if err != nil {
		return nil, errors.E(err, util.ErrNotStored,
			errors.Params{"database": database, "collection": collection, "query": value})
	}
	return res.InsertedIDs, nil
}

func (db *MongoDb) Update(database, collection string, value interface{}, query bson.D) (interface{}, error) {
	c := db.Database(database).Collection(collection)
	update := bson.D{{"$set", value}}
	res, err := c.UpdateOne(context.Background(), query, update)
	if err != nil {
		return nil, errors.E(err, util.ErrNotStored,
			errors.Params{"database": database, "collection": collection, "query": value})
	}
	return res.UpsertedID, nil
}

func (db *MongoDb) DeleteOne(database, collection string, query bson.D) (int64, error) {
	c := db.Database(database).Collection(collection)
	res, err := c.DeleteOne(context.TODO(), query)
	if err != nil {
		return 0, errors.E(err, util.ErrNotStored,
			errors.Params{"database": database, "collection": collection, "query": query})
	}
	return res.DeletedCount, nil
}

func (db *MongoDb) DeleteMany(database, collection string, query interface{}) (int64, error) {
	c := db.Database(database).Collection(collection)
	res, err := c.DeleteMany(context.TODO(), query)
	if err != nil {
		return 0, errors.E(err, util.ErrNotStored,
			errors.Params{"database": database, "collection": collection, "query": query})
	}
	return res.DeletedCount, nil
}

func (db *MongoDb) IsReady() bool {
	err := db.Ping(context.TODO(), readpref.Primary())
	return !(err != nil)
}

func getResult(doc []interface{}, result interface{}) error {
	data, err := json.Marshal(doc)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &result)
	return err
}
