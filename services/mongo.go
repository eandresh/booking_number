package services

import (
	"context"
	"log"

	"eh-digital-shift/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongodb struct {
	cnf        *config.Config
	collection *mongo.Collection
	ctx        context.Context
}

func (m *Mongodb) init() error {
	client, err := m.getClient()
	if err != nil {
		return err
	}

	err = client.Ping(m.ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	m.collection = client.Database(m.cnf.Mongo.Database).Collection(m.cnf.Mongo.Collection)

	return nil
}

func (m *Mongodb) getClient() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(m.cnf.Mongo.Uri)
	client, err := mongo.Connect(m.ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return client, nil
}

func (m *Mongodb) Insert(ctx context.Context, obj interface{}) error {

	_, err := m.collection.InsertOne(ctx, obj)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongodb) FindOne(ctx context.Context, filter interface{}, result *interface{}) error {
	err := m.collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (m *Mongodb) FindAll(ctx context.Context, filter interface{}) ([]bson.M, error) {
	list, err := m.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var rows []bson.M
	err = list.All(ctx, &rows)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func NewMongoDB(ctx context.Context, cnf *config.Config) (*Mongodb, error) {
	m := &Mongodb{
		cnf: cnf,
		ctx: ctx,
	}
	err := m.init()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return m, nil
}
