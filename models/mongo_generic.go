package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (m MongoObject[T]) Create(obj T) (primitive.ObjectID, error) {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	result, err := m.db.InsertOne(ctx, obj)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	objectId, _ := result.InsertedID.(primitive.ObjectID)

	return objectId, nil
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

func (m MongoObject[T]) Count(filter bson.D) (int64, error) {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	countDocuments, err := m.db.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return countDocuments, nil

}

func (m MongoObject[T]) Find(filter bson.D, mongoParams interfaces.MongoFilter) ([]T, error) {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	opts := options.Find()
	opts.SetSkip(int64((mongoParams.Page - 1) * mongoParams.PerPage))
	opts.SetLimit(int64(mongoParams.PerPage))
	opts.SetProjection(mongoParams.Projection)
	asc := -1
	if mongoParams.Asc {
		asc = 1
	}
	opts.SetSort(bson.M{mongoParams.SortField: asc})

	cursor, err := m.db.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []T
	for cursor.Next(ctx) {
		var result T
		if err = cursor.Decode(&result); err != nil {
			return nil, err
		}

		results = append(results, result)
	}

	return results, nil
}

func (m MongoObject[T]) FindOne(filter bson.D, projection ...bson.D) (*T, error) {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	opt := options.FindOne()
	if len(projection) > 0 {
		opt.SetProjection(projection[0])
	}

	result := m.db.FindOne(ctx, filter, opt)
	if err := result.Err(); err != nil {
		return nil, err
	}

	var obj T
	if err := result.Decode(&obj); err != nil {
		return nil, err
	}

	return &obj, nil
}

func (m MongoObject[T]) Update(filters bson.D, update bson.D) error {
	ctx, c := context.WithTimeout(context.Background(), m.timeout)
	defer c()

	_, err := m.db.UpdateOne(ctx, filters, update)
	if err != nil {
		return err
	}

	return nil
}
