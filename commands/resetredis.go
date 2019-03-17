package commands

import "github.com/Farteen/travelfinance/redisclient"

func ResetRedis() {
	redisclient.RedisClient().FlushAll()
}