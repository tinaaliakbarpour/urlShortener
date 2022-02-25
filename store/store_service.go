package store

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)



const CacheDuration = 6 * time.Hour

//InitializeStore
func (s *store)InitializeStore()  {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	s.redisClient = redisClient
}

// if user id was not provided generate one on the fly : case for not logged in users

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/
func (s *store)SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	err := s.redisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}

	fmt.Printf("Saved shortUrl: %s - originalUrl: %s\n", shortUrl, originalUrl)
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/
func (s *store)RetrieveInitialUrl(shortUrl string) string {
	result, err := s.redisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	return result
}
