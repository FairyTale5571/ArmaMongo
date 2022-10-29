package mongo

import (
	"context"
	"time"

	"github.com/fairytale5571/mongo/pkg/logger"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	db     *mgo.Client
	logger *logger.Wrapper
}

func New(url string) (*Mongo, error) {
	// create mongo connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mgo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	return &Mongo{
		db:     client,
		logger: logger.New("Mongo"),
	}, nil
}

func (m *Mongo) Write(database, collection string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := m.db.Database(database).Collection(collection).InsertOne(ctx, data)
	if err != nil {
		m.logger.Errorf("Write: %s", err)
		return err
	}
	return nil
}

func (m *Mongo) Version() string {
	return "1.6"
}
