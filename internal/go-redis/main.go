package main

import (
	"context"
	"log"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

func main() {
	ctx := context.Background()

	redis := goredis.NewClient(&goredis.Options{Addr: "127.0.0.1:6370"})

	if err := redis.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	defer redis.Close()

	// SET <key> <value>
	{
		cmd := redis.Set(ctx, "key", "value", goredis.KeepTTL)
		if err := cmd.Err(); err != nil {
			log.Println(err)
		}
	}

	// SET <key> <value> EX <seconds>
	{
		cmd := redis.Set(ctx, "key", "value", 10*time.Second)
		if err := cmd.Err(); err != nil {
			log.Println(err)
		}
	}

	// SET <key> <value> NX
	{
		cmd := redis.SetNX(ctx, "key", "value", 10*time.Second)
		if err := cmd.Err(); err != nil {
			log.Println(err)
		}
	}

	// SET <key> <value> XX
	{
		cmd := redis.SetXX(ctx, "key", "value", 10*time.Second)
		if err := cmd.Err(); err != nil {
			log.Println(err)
		}
	}
}
