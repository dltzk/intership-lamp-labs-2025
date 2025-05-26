package redisclient

import (
    "context"
    "log"
    redis "github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func InitRedis() *redis.Client {
    rdb = redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", 
        Password: "",               
        DB:       0,               
    })

    if err := rdb.Ping(ctx).Err(); err != nil {
        log.Fatalf("Cant access Redis: %v", err)
    }

    log.Println("Successfuly accessed Redis")

    return rdb
}