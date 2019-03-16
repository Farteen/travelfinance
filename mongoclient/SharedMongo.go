package mongoclient

import (
	"github.com/gpmgo/gopm/modules/log"
	"github.com/mongodb/mongo-go-driver/mongo"
	"sync"
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
		sg = &singleton{
			client,
		}
	})
	return sg.mongoClient
}
