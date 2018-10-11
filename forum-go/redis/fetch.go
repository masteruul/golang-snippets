package redis

import (
	"fmt"
	"log"

	"github.com/go-redis/redis"
)

func UseRedis() {
	client, err := InitRedis()
	if err != nil {
		log.Println("something")
	}
	client.Set("key", "value", 0).Err()
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
