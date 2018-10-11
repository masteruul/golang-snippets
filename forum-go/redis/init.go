package redis

import (
	"log"

	"github.com/go-redis/redis"
)

func InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	log.Println(pong, err)
	if err == nil {
		return client, nil
	}
	return client, err
	// Output: PONG <nil>
}
