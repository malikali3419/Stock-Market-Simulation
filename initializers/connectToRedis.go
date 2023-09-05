package initializers

import (
	"context"
	"github.com/go-redis/redis/v8"
	"stock_market_simulation/m/constants"
)

var REDISCLIENT *redis.Client

func ConnectToRedis() {
	REDISCLIENT = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := REDISCLIENT.Ping(context.Background()).Result()
	if err != nil {
		panic(constants.FailedToConnectRedis + err.Error())
	}
}
