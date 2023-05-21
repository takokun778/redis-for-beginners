package main

import (
	"context"
	"log"

	"github.com/redis/rueidis"
)

func main() {
	ctx := context.Background()

	redis, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6370"}})
	if err != nil {
		panic(err)
	}

	if err := redis.Do(ctx, redis.B().Ping().Build()).Error(); err != nil {
		panic(err)
	}

	defer redis.Close()

	// 自動補完を有効にしていると
	// redis.Do()の引数をコンパイル通すために redis.B()以降の既述をすぐに閉じようとしてキャストされてしまう

	{
		// SET <key> <value>
		// cmd := redis.B().Set().Key("key").Value("value").Build()

		// SET <key> <value> EX <seconds>
		// cmd := redis.B().Set().Key("key").Value("value").ExSeconds(10).Build()

		// SET <key> <value> NX
		// cmd := redis.B().Set().Key("key").Value("value").Nx().Build()

		// SET <key> <value> XX
		// cmd := redis.B().Set().Key("key").Value("value").Xx().Build()

		// if err := redis.Do(ctx, cmd).Error(); err != nil {
		// 	log.Println(err)
		// }
	}

	{
		// GET <key>
		// cmd := redis.B().Get().Key("key").Build()
		// if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
		// 	log.Println(err)
		// } else {
		// 	log.Printf("GET key: %s, value: %s", "key", val)
		// }
	}

	{
		// cmd1 := redis.B().Set().Key("key1").Value("val1").Build()

		// cmd2 := redis.B().Set().Key("key2").Value("val2").Build()

		// cmd3 := redis.B().Set().Key("key3").Value("val3").Build()

		// // MSET <key1> <value1> <key2> <value2> ...
		// for _, res := range redis.DoMulti(ctx, cmd1, cmd2, cmd3) {
		// 	if err := res.Error(); err != nil {
		// 		log.Println(err)
		// 	}
		// }

		// // MGET <key1> <key2> ...
		// get1 := redis.B().Get().Key(key1).Build()

		// get2 := redis.B().Get().Key(key2).Build()

		// get3 := redis.B().Get().Key(key3).Build()

		// for _, res := range redis.DoMulti(ctx, get1, get2, get3) {
		// 	if val, err := res.ToString(); err != nil {
		// 		log.Println(err)
		// 	} else {
		// 		log.Printf("GET value: %s", val)
		// 	}
		// }
	}

	// APPEND <key> <value>
	{
		// cmd := redis.B().Set().Key("key").Value("val").Build()

		// if err := redis.Do(ctx, cmd).Error(); err != nil {
		// 	log.Println(err)
		// }

		// cmd = redis.B().Append().Key("key").Value("val").Build()

		// if err := redis.Do(ctx, cmd).Error(); err != nil {
		// 	log.Println(err)
		// }

		// cmd = redis.B().Get().Key("key").Build()

		// if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
		// 	log.Println(err)
		// } else {
		// 	log.Printf("GET key: %s, value: %s", "key", val)
		// }
	}

	// STRLEN <key>
	// {
	// 	cmd := redis.B().Strlen().Key("key").Build()
	// 	if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("LEN %d", val)
	// 	}
	// }

	// GETRANGE <key> <start> <end>
	// {
	// 	cmd := redis.B().Getrange().Key("key").Start(1).End(3).Build()
	// 	if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("GETRANGE %s", val)
	// 	}
	// }

	// SETRANGE <key> <offset> <value>
	// {
	// 	cmd := redis.B().Setrange().Key("key").Offset(6).Value("Redis").Build()
	// 	if err := redis.Do(ctx, cmd).Error(); err != nil {
	// 		log.Println(err)
	// 	}
	// }

	// INCR <key>
	// {
	// 	cmd := redis.B().Incr().Key("key").Build()
	// 	if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("INCR %d", val)
	// 	}
	// }

	// INCRBY <key> <increment>
	// {
	// 	cmd := redis.B().Incrby().Key("key").Increment(5).Build()
	// 	if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("INCRBY %d", val)
	// 	}
	// }

	// INCRBYFLOAT <key> <increment>
	// {
	// 	cmd := redis.B().Incrbyfloat().Key("key").Increment(4.6).Build()
	// 	if val, err := redis.Do(ctx, cmd).AsFloat64(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("INCRBYFLOAT %v", val)
	// 	}
	// }

	// DECR <key>
	// {
	// 	cmd := redis.B().Decr().Key("key").Build()
	// 	if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("DECR %d", val)
	// 	}
	// }

	// DECRBY <key> <decrement>
	{
		cmd := redis.B().Decrby().Key("key").Decrement(4).Build()
		if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
			log.Println(err)
		} else {
			log.Printf("DECRBY %d", val)
		}
	}

	// GETDEL <key>
	// {
	// 	cmd := redis.B().Getdel().Key("key").Build()
	// 	if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
	// 		log.Println(err)
	// 	} else {
	// 		log.Printf("GETDEL %s", val)
	// 	}
	// }

	// keys <pattern>
	{
		cmd := redis.B().Keys().Pattern("key*").Build()
		if val, err := redis.Do(ctx, cmd).AsStrSlice(); err != nil {
			log.Println(err)
		} else {
			for _, v := range val {
				log.Printf("KEYS %s", v)
			}
		}
	}

	// EXISTS <key>
	{
		cmd := redis.B().Exists().Key("key").Build()
		if val, err := redis.Do(ctx, cmd).AsBool(); err != nil {
			log.Println(err)
		} else {
			log.Printf("EXISTS %v", val)
		}
	}

	{
		cmd := redis.B().Del().Key("key").Build()
		if err := redis.Do(ctx, cmd).Error(); err != nil {
			log.Println(err)
		}
	}
}
