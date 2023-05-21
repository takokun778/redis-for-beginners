package main

import (
	"context"
	"log"

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
	// {
	// 	cmd := redis.Set(ctx, "key", "value", goredis.KeepTTL)
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// SET <key> <value> EX <seconds>
	// {
	// 	cmd := redis.Set(ctx, "key", "value", 10*time.Second)
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// SET <key> <value> NX
	// {
	// 	cmd := redis.SetNX(ctx, "key", "value", 10*time.Second)
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// SET <key> <value> XX
	// {
	// 	cmd := redis.SetXX(ctx, "key", "value", 10*time.Second)
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// GET <key>
	// {
	// 	cmd := redis.Get(ctx, "key")
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%s", res)
	// 	}
	// }

	// MSET <key> <value>
	// {
	// 	cmd := redis.MSet(ctx, "key1", "value1", "key2", "value2", "key3", "value3")
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// MGET <key> <value>
	// {
	// 	cmd := redis.MGet(ctx, "key1", "key2", "key3")
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%v", res)
	// 	}
	// }

	// APPEND <key> <value>
	// {
	// 	cmd := redis.Append(ctx, "key", "value")
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// STRLEN <key>
	// {
	// 	cmd := redis.StrLen(ctx, "key")
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%v", res)
	// 	}
	// }

	// GETRANGE <key> <start> <end>
	// {
	// 	cmd := redis.GetRange(ctx, "key", 1, 3)
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%v", res)
	// 	}
	// }

	// SETRANGE <key> <offset> <value>
	// {
	// 	cmd := redis.SetRange(ctx, "key", 6, "Redis")
	// 	if err := cmd.Err(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// INCR <key>
	// {
	// 	cmd := redis.Incr(ctx, "key")
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%v", res)
	// 	}
	// }

	// INCRBY <key> <increment>
	// {
	// 	cmd := redis.IncrBy(ctx, "key", 5)
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%v", res)
	// 	}
	// }

	// INCRBYFLOAT <key> <increment>
	// {
	// 	cmd := redis.IncrByFloat(ctx, "key", 4.6)
	// 	if res, err := cmd.Result(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("%v", res)
	// 	}
	// }

	// DECR <key>
	{
		cmd := redis.Decr(ctx, "key")
		if res, err := cmd.Result(); err != nil {
			log.Println(err)
		} else {
			log.Printf("%v", res)
		}
	}

	// DECRBY <key> <decrement>
	{
		cmd := redis.DecrBy(ctx, "key", 4)
		if res, err := cmd.Result(); err != nil {
			log.Println(err)
		} else {
			log.Printf("%v", res)
		}
	}

	// GETDEL <key>
	{
		cmd := redis.GetDel(ctx, "key")
		if res, err := cmd.Result(); err != nil {
			log.Println(err)
		} else {
			log.Printf("%v", res)
		}
	}

	// KEYS <pattern>
	{
		cmd := redis.Keys(ctx, "key*")
		if res, err := cmd.Result(); err != nil {
			log.Println(err)
		} else {
			log.Printf("%v", res)
		}
	}

	// EXISTS <key>
	{
		cmd := redis.Exists(ctx, "key")
		if res, err := cmd.Result(); err != nil {
			log.Println(err)
		} else {
			log.Printf("%v", res)
		}
	}

	// DEL <key>
	{
		cmd := redis.Del(ctx, "key")
		if err := cmd.Err(); err != nil {
			log.Println(err)
		}
	}
}
