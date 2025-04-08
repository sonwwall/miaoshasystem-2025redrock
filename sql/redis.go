package sql

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"log"
)

var ctx = context.Background()
var R *redis.Client

func Init() {
	viper.SetConfigFile("../global/redisConfig.yaml")
	R = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("DB.addr"),
		Password: viper.GetString("DB.password"),
		DB:       viper.GetInt("DB.DB"),
	})

	_, err := R.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully")
}
