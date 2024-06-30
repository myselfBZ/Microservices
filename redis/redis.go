package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)




func RedisInit() *redis.Client{
    Client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        DB: 0,
        Password: "",
    })
    _, err := Client.Ping(context.TODO()).Result()
    if err != nil{
        panic("You suck at initializing Redis client")
    }
    return Client 
}
