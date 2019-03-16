package redisclient

import (
	"github.com/go-redis/redis"
	"sync"
)

type singleton struct {
	redisClient *redis.Client
}

var once = sync.Once{}
var sg *singleton = nil

func init() {
	RedisClient()
}

func RedisClient() *redis.Client {
	once.Do(func() {
		redisClient := redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",//"redisclient-16639.c114.us-east-1-4.ec2.cloud.redislabs.com:16639",
			Password: "",//"5r04BNROKJr4p2cvJ5nOCXT6jnMROb44",
			DB:       0,
		})
		sg = &singleton{
			redisClient: redisClient,
		}
	})
	return sg.redisClient
}
