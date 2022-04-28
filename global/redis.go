package global

import "github.com/go-redis/redis/v8"

var RedisDB *redis.Client

func InitRedis() {
	RedisDB = redis.NewClient(&redis.Options{
		Addr:	  RedisInfo.Host+":6379",
		Password: "", // no password set
		DB:		  0,  // use default DB
	})
}
