package interfaces

import (
	"go.mongodb.org/mongo-driver/bson"
)

type RepositoryMongo[T any] interface {
	Create(T) error
	Delete(filter bson.D) error
	FindOne(filter bson.D, projection ...bson.D) (*T, error)
}
