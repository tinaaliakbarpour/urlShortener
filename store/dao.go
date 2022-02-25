package store

import(
	"context"
	"github.com/go-redis/redis/v8"
)

var (
	Store storeInterface = &store{}
	ctx          = context.Background()
)

type storeInterface interface {
	InitializeStore()
	SaveUrlMapping(shortUrl string, originalUrl string, userId string)
	RetrieveInitialUrl(shortUrl string) string
}

type store struct {
	redisClient *redis.Client
}