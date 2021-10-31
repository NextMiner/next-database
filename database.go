package next_database

import (
	"context"
	"github.com/go-redis/redis/v8"
	logging "github.com/ipfs/go-log/v2"
	"strings"
)

var log = logging.Logger("database")

type Database struct {
	*redis.Client
	coin string
}

func InitDatabase(coinName string, options *DatabaseConfig) *Database {
	client := redis.NewClient(options.ToRedisOptions())
	if client == nil {
		log.Panic("Failed to connect to the redis server.")
		return nil
	}

	result, err := client.Ping(context.Background()).Result()
	if err != nil || strings.ToLower(result) != "pong" {
		log.Panicf("Failed to connect to the redis server: %s %s", result, err)
	}

	return &Database{
		Client: client,
		coin:   coinName,
	}
}