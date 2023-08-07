package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/IIGabriel/btc-tx-manager/constants"
	"github.com/IIGabriel/btc-tx-manager/interfaces"
)

func NewMongoObject[T any](dirty T, collection *mongo.Collection) interfaces.RepositoryMongo[T] {
	return &MongoObject[T]{dirty: dirty, MongoConn: MongoConn{collection, constants.MongoTimeout}}
}

type MongoObject[T any] struct {
	dirty T
	MongoConn
}

type MongoConn struct {
	db      *mongo.Collection
	timeout time.Duration
}

func (m MongoObject[T]) Create(obj T) error {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	if _, err := m.db.InsertOne(ctx, obj); err != nil {
		return err
	}

	return nil
}
func (m MongoObject[T]) Delete(filter bson.D) error {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	result, err := m.db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("no matching data")
	}
	return nil
}
func (m MongoObject[T]) FindOne(filter bson.D, projection ...bson.D) (*T, error) {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	var opt options.FindOneOptions
	if len(projection) > 0 {
		opt.Projection = projection[0]
	}

	result := m.db.FindOne(ctx, filter, &opt)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var obj T
	if err := result.Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}
