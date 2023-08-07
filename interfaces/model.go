package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryMongo[T any] interface {
	Create(T) (primitive.ObjectID, error)
	Count(filter bson.D) (int64, error)
	Delete(filter bson.D) error
	FindOne(filter bson.D, projection ...bson.D) (*T, error)
	Find(filter bson.D, mongoParams MongoFilter) ([]T, error)
	Update(filters bson.D, update bson.D) error
}
