package mongoclient

import (
	"context"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/mongodb/mongo-go-driver/mongo"
	"sync"
	"time"
)

type singleton struct {
	mongoClient *mongo.Client
}

var once = sync.Once{}
var sg *singleton = nil

func init() {
	MongoClient()
}

func MongoClient() *mongo.Client {
	once.Do(func() {
		client, err := mongo.NewClient("mongodb://localhost:27017")
		if err != nil {
			log.Fatal("fata error init mongo client")
		}
		ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Fatal("mongodb connect err")
		}
		sg = &singleton{
			client,
		}
	})
	return sg.mongoClient
}

func MongoDBCollectionWithName(collectionName string) *mongo.Collection {
	collection := MongoClient().Database("travelfinance").Collection(collectionName)
	return collection
}
