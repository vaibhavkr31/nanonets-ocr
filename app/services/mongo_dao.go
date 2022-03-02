package services

import (
	"context"
	"fmt"
	clt "github.com/dripcapital/nanonets-ocr/app/client"
	"github.com/dripcapital/nanonets-ocr/app/config"
	st "github.com/dripcapital/nanonets-ocr/app/structs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

var client *mongo.Client
var logger *zap.Logger

func init() {
	logger = config.InitLogger("dao")
	client = clt.MongoClient()
}

func StoreNanonetResponse(bldocResp st.BlDocResponse) {
	s := fmt.Sprintf("Inserting Nanonets Response in bl-doc colletion")
	logger.Debug(s)
	collection := client.Database("nanonets").Collection("bl-doc")
	_, err := collection.InsertOne(context.TODO(), bldocResp)
	if err != nil {
		logger.Error(fmt.Sprintf("Error Returned While Inserting in Mongo Database : %s", err))
	}
	logger.Debug("Successfully Inserted the record into Mongo Db")
}
