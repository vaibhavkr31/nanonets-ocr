package client

import (
	"context"
	"fmt"
	"github.com/dripcapital/nanonets-ocr/app/config"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	logger = config.InitLogger("client")
}
func MongoClient() *mongo.Client {
	logger.Debug("Creating a mongoClient")
	// Set client options
	clientOptions := options.Client().ApplyURI(viper.GetString("mongoURI"))

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		errMsg := fmt.Sprintf("Connection to Mongo DB Failed %s", err)
		logger.Error(errMsg)
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		errMsg := fmt.Sprintf("Connection to Mongo DB Failed %s", err)
		logger.Error(errMsg)
		panic(err)
	}

	logger.Debug("Mongo Client Created")
	return client
}
