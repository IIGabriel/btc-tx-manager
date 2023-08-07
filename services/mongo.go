package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/mgocompat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"

	"github.com/IIGabriel/eth-tx-manager/constants"
	"github.com/IIGabriel/eth-tx-manager/utils"
)

var mongoInstance *mongo.Database

func Mongo() *mongo.Database {
	if mongoInstance == nil {
		clientOpts := options.Client().
			SetRegistry(mgocompat.Registry).
			ApplyURI(utils.EnvString(constants.MongoEnvKey))

		client, err := mongo.Connect(context.Background(), clientOpts)
		if err != nil {
			zap.L().Fatal("failed create client of mongo", zap.Error(err))
		}
		ctx, c := context.WithTimeout(context.Background(), constants.MongoTimeout)
		defer c()

		err = client.Connect(ctx)
		if err != nil {
			zap.L().Fatal("failed to connect to mongo", zap.Error(err))
		}

		ctx, c = context.WithTimeout(context.Background(), constants.MongoTimeout)
		defer c()
		if err := client.Ping(ctx, nil); err != nil {
			zap.L().Fatal("failed to ping mongo", zap.Error(err))
		}
		mongoInstance = client.Database(utils.EnvString(constants.MongoDataBaseKey))
	}

	return mongoInstance
}
