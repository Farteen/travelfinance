package mongoclient

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
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
			log.Println("fatal error init mongo client")
		}
		ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
		err = client.Connect(ctx)
		if err != nil {
			log.Println("mongodb connect err")
		}
		sg = &singleton{
			client,
		}
	})
	return sg.mongoClient
}

func Collection(collectionName string) *mongo.Collection {
	collection := MongoClient().Database("travelfinance").Collection(collectionName)
	return collection
}
