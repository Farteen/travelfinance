package commands

import (
	"context"
	"github.com/Farteen/travelfinance/mongoclient"
)

func ResetMongo(collections ...string) {
	for _, collection := range collections{
		mongoclient.Collection(collection).Drop(context.Background())
	}
}