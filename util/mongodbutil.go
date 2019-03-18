package util

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func MongoDBHexID(result *mongo.InsertOneResult) string {
	return result.InsertedID.(primitive.ObjectID).Hex()
}

func MongoDBOID(result *mongo.InsertOneResult) primitive.ObjectID {
	return result.InsertedID.(primitive.ObjectID)
}
