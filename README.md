# redis-for-beginners

---
title: "KVSとしてのRedisに入門しgo-redisとrueidisから触ってみる"
emoji: "🐁"
type: "tech" # tech: 技術記事 / idea: アイデア
topics: ["go","redis"]
published: false
---

# はじめに

この記事は`Redis`に入門しつつGoライブラリの`go-redis`と`rueidis`を比較してみようという内容です。

先日、twitterを眺めていたところ`rueidis`というRedisのGoライブラリを知りました。

https://github.com/redis/rueidis

これまで私は`go-redis`を使っていたのですが、`supports client side caching.`という一文が気になり使ってみたくなりました。

https://github.com/redis/go-redis

ついでにこのタイミングでなんとなく使っていたRedisについて改めて入門してみようと思いました。
Redisに入門するにあたって[公式ドキュメント](https://redis.io/)だと私には少し分かりづらかったので以下の書籍を参考しています。

「実践Redis入門 技術の仕組みから現場の活用まで」

https://gihyo.jp/book/2022/978-4-297-13142-5

# [Redis](https://redis.io/)とは

公式サイトには以下のような記載があります。

> The open source, in-memory data store used by millions of developers as a database, cache, streaming engine, and message broker.

データベース、キャッシュ、ストリーミングエンジン、メッセージブローカーです。
インメモリのデータストアであるので高速に動作することが可能です。
データストアの形式としては`NoSQL`と呼ばれるものに分類されます。
`KVS`として気軽に利用できる一方で`RDB`のような柔軟さはないかと思います。

また、Redisはデータ型さまざまなものをサポートしています。

- String型
- List型
- Hash型
- Set型
- Sorted Set型

いまのところ、私的にString型しか需要がないので今回の記事としてはスコープをString型に絞っています。

# String型：文字列

シンプルなキーと値の組みあわせです。
Redisにおいては数値なども文字列に含まれます。
また、バイナリーセーフなため、画像や実行ファイルなど文字列以外もこの型で保存できます。（知らんかったです...）

# [go-redis](https://github.com/redis/go-redis)

```go
import goredis "github.com/redis/go-redis/v9"

func main() {
	ctx := context.Background()

	redis := goredis.NewClient(&goredis.Options{Addr: "127.0.0.1:6370"})

	if err := redis.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	defer redis.Close()
}
```

# [rueidis](https://github.com/redis/rueidis)

```go
import "github.com/redis/rueidis"

func main() {
	redis, err := rueidis.NewClient(rueidis.ClientOption{InitAddress: []string{"127.0.0.1:6370"}})
	if err != nil {
		panic(err)
	}

	if err := redis.Do(ctx, redis.B().Ping().Build()).Error(); err != nil {
		panic(err)
	}

	defer redis.Close()
}
```

# [Commands](https://redis.io/commands/)

ここから標準的なユースケースで使いそうなRedisのCommandsに入門しつつ、
そのコマンドがGoライブラリでどのように実装できるのか見ていきます。

## [SET <key> <value>](https://redis.io/commands/set/)

指定したキーに対して値を設定することができます。

```shell
127.0.0.1:6379:> SET key value
OK
127.0.0.1:6379:> GET key
"value"
```

```go
// go-redis
cmd := redis.Set(ctx, "key", "value", goredis.KeepTTL)
if err := cmd.Err(); err != nil {
    log.Println(err)
}

// rueidis
cmd := redis.B().Set().Key("key").Value("value").Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

EXオプションを使用してTTLを設定することができます。

```shell
127.0.0.1:6379> SET key value EX 10
OK
127.0.0.1:6379> GET key
"value"
# 10秒後
127.0.0.1:6379> GET key
(nil)
```

```go
// go-redis
cmd := redis.Set(ctx, "key", "value", 10*time.Second)
if err := cmd.Err(); err != nil {
    log.Println(err)
}

// rueidis
cmd := redis.B().Set().Key("key").Value("value").ExSeconds(10).Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

NXオプションを使用すると、キーが存在しない場合のみ値をセットすることができます。

```shell
127.0.0.1:6379> SET key value NX
OK
127.0.0.1:6379> SET key val NX
(nil)
127.0.0.1:6379> GET key
"value"
```

```go
// go-redis
cmd := redis.SetNX(ctx, "key", "value", 10*time.Second)
if err := cmd.Err(); err != nil {
    log.Println(err)
}

// rueidis
cmd := redis.B().Set().Key("key").Value("value").Nx().Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

XXオプションを使用すると、キーが存在する場合のみ値をセットすることができます。

```shell
127.0.0.1:6379> SET key value XX
(nil)
127.0.0.1:6379> GET key
(nil)
127.0.0.1:6379> SET key value
OK
127.0.0.1:6379> GET key
"value"
127.0.0.1:6379> SET key val XX
OK
127.0.0.1:6379> GET key
"val"
```

```go
// go-redis
cmd := redis.SetXX(ctx, "key", "value", 10*time.Second)
if err := cmd.Err(); err != nil {
    log.Println(err)
}

// rueidis
cmd := redis.B().Set().Key("key").Value("value").Xx().Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

## [GET <key>](https://redis.io/commands/get/)

SETを用いて保存したkeyの値を取得します。
指定したkeyが存在しない場合はnilを返します。

```shell
127.0.0.1:6379> SET key value
OK
127.0.0.1:6379> GET key
"value"
127.0.0.1:6379> GET k
(nil)
```

```go
// rueidis
cmd := redis.B().Get().Key(key).Build()
if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
    log.Println(err)
} else {
    log.Printf("GET key: %s, value: %s", key, val)
}
```

## [MSET <key> <value> [<key> <value> ...]](https://redis.io/commands/mset/)

指定したキーと値のペアを複数設定することができます。
MSETはアトミックであるため指定されたすべてのキーを一度に設定することができます。

```shell
127.0.0.1:6379:> MSET key1 value1 key2 value2 key3 value3
OK
```

```go
// rueidis
cmd1 := redis.B().Set().Key("key1").Value("val1").Build()
cmd2 := redis.B().Set().Key("key2").Value("val2").Build()
cmd3 := redis.B().Set().Key("key3").Value("val3").Build()
for _, res := range redis.DoMulti(ctx, cmd1, cmd2, cmd3) {
    if err := res.Error(); err != nil {
        log.Println(err)
    }
}
```

## [MGET <key> [<key> ...]](https://redis.io/commands/mget/)

指定されたキーの値を取得することができます。

```shell
127.0.0.1:6379:> MGET key1 key2 key3
1) "value1"
2) "value2"
3) "value3"
```

```go
// rueidis
cmd1 := redis.B().Get().Key("key1").Build()
cmd2 := redis.B().Get().Key("key2").Build()
cmd3 := redis.B().Get().Key("key3").Build()
for _, res := range redis.DoMulti(ctx, cmd1, cmd2, cmd3) {
    if val, err := res.ToString(); err != nil {
        log.Println(err)
    } else {
        log.Printf("GET value: %s", val)
    }
}
```

## [APPEND <key> <value>](https://redis.io/commands/append/)

キーが存在する場合は、キーの値の末尾に引数の内容を追加します。キーが存在しない場合は、SETと同じ挙動となります。
取得した結果はSETしたvalueの文字数が表示されます。

```shell
127.0.0.1:6379> APPEND key value
(integer) 5
```

```go
// rueidis
cmd := redis.B().Append().Key("key").Value("val").Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

## [STRLEN <key>](https://redis.io/commands/strlen/)

キーで指定した値の文字列の長さを取得します。

```shell
127.0.0.1:6379> STRLEN key
(integer) 5
```

```go
// rueidis
cmd := redis.B().Strlen().Key("key").Build()
if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
    log.Println(err)
} else {
    log.Printf("LEN %d", val)
}
```

## [GETRANGE <key> <start> <end>](https://redis.io/commands/getrange/)

キーで指定した値の文字列のうち、指定した範囲の文字列を取得します。

```shell
127.0.0.1:6379> GETRANGE key 1 3
"alu" # valueのうち[1]から[3]までの文字列を取得
```

```go
// rueidis
cmd := redis.B().Getrange().Key("key").Start(1).End(3).Build()
if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
    log.Println(err)
} else {
    log.Printf("GETRANGE %s", val)
}
```

## [SETRANGE <key> <offset> <value>]()

指定したオフセットの位置を起点としてキーに値をセットします。

```shell
127.0.0.1:6379> SET key HelloWorld
OK
127.0.0.1:6379> SETRANGE key 6 Redis
(integer) 11
127.0.0.1:6379> GET key
"HelloWRedis"
```

```go
// rueidis
cmd := redis.B().Setrange().Key("key").Offset(6).Value("Redis").Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

## [INCR <key>](https://redis.io/commands/incr/)

キーの値をインクリメントします。指定したキーが存在しない場合には、操作前に0がセットされます。

```shell
127.0.0.1:6379> INCR key
(integer) 1
127.0.0.1:6379> GET key
"1"
127.0.0.1:6379> INCR key
(integer) 2
127.0.0.1:6379> GET key
"2"
```

```go
// rueidis
cmd := redis.B().Incr().Key("key").Build()
if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
    log.Println(err)
} else {
    log.Printf("INCR %d", val)
}
```

## [INCRBY <key> <increment>](https://redis.io/commands/incrby/)

キーと値を指定した整数だけ増加します。指定したキーが存在しない場合には、操作前に0がセットされます。

```shell
127.0.0.1:6379> INCRBY key 5
(integer) 5
```


```go
// rueidis
cmd := redis.B().Incrby().Key("key").Increment(5).Build()
if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
    log.Println(err)
} else {
    log.Printf("INCRBY %d", val)
}
```

## [INCRBYFLOAT <key> <increment>](https://redis.io/commands/incrbyfloat/)

浮動小数点数を指定した値だけ増加します。指定したキーが存在しない場合には、操作前に0がセットされます。
数値を減らす場合は、マイナスを指定します。

```shell
127.0.0.1:6379> INCRBYFLOAT key 4.6
"4.6"
```

```go
// rueidis
cmd := redis.B().Incrbyfloat().Key("key").Increment(4.6).Build()
if val, err := redis.Do(ctx, cmd).AsFloat64(); err != nil {
    log.Println(err)
} else {
    log.Printf("INCRBYFLOAT %v", val)
}
```

## [DECR <key>](https://redis.io/commands/decr/)

キーの値を1だけ減少します。指定したキーが存在しない場合には、操作前に0がセットされます。

```shell
127.0.0.1:6379> DECR key
(integer) -1
```

```go
// rueidis
cmd := redis.B().Decr().Key("key").Build()
if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
    log.Println(err)
} else {
    log.Printf("DECR %d", val)
}
```

## [DECRBY <key> <decrement>](https://redis.io/commands/decrby/)

キーの値を指定した整数だけ減少します。指定したキーが存在しない場合には、操作前に0がセットされます。

```shell
127.0.0.1:6379> DECRBY key 4
(integer) -4
```

```go
// rueidis
cmd := redis.B().Decrby().Key("key").Decrement(4).Build()
if val, err := redis.Do(ctx, cmd).ToInt64(); err != nil {
    log.Println(err)
} else {
    log.Printf("DECRBY %d", val)
}
```

## [GETDEL <key>](https://redis.io/commands/getdel/)

キーの値を取得し、同時にそのキーを削除します。

```shell
127.0.0.1:6379> SET key value
OK
127.0.0.1:6379> GETDEL key
"value"
127.0.0.1:6379> GET key
(nil)
```

```go
// rueidis
cmd := redis.B().Getdel().Key("key").Build()
if val, err := redis.Do(ctx, cmd).ToString(); err != nil {
    log.Println(err)
} else {
    log.Printf("GETDEL %s", val)
}
```

## [KEYS <pattern>](https://redis.io/commands/keys/)

キーの一覧を取得します。パターンを指定することで、キーの一覧を絞り込むことができます。
ただし、KEYSコマンドはパフォーマンスが悪いため、本番環境ではSCAN系のコマンドを使用することが推奨されています。

```shell
127.0.0.1:6379> KEYS *
1) "key3"
2) "key1"
3) "key2"
4) "foo"
5) "key"
127.0.0.1:6379> KEYS key*
1) "key3"
2) "key1"
3) "key2"
4) "key"
```

```go
// rueidis
cmd := redis.B().Keys().Pattern("key*").Build()
if val, err := redis.Do(ctx, cmd).AsStrSlice(); err != nil {
    log.Println(err)
} else {
    for _, v := range val {
        log.Printf("KEYS %s", v)
    }
}
```
## [EXEISTS <key> [<key> ...]](https://redis.io/commands/exists/)

指定したキーが存在するかどうかを確認します。

```shell
127.0.0.1:6379> SET key velue
OK
127.0.0.1:6379> EXISTS key
(integer) 1
127.0.0.1:6379> EXISTS k
(integer) 0
```

```go
// rueidis
cmd := redis.B().Exists().Key("key").Build()
if val, err := redis.Do(ctx, cmd).AsBool(); err != nil {
    log.Println(err)
} else {
    log.Printf("EXISTS %v", val)
}
```

## [DEL <key>](https://redis.io/commands/del/)

指定したキーを削除します。キーが存在しない場合は無視されます。

```shell
127.0.0.1:6379> SET key velue
OK
127.0.0.1:6379> DEL key
(integer) 1
127.0.0.1:6379> DEL key
(integer) 0
```

```go
// rueidis
cmd := redis.B().Del().Key("key").Build()
if err := redis.Do(ctx, cmd).Error(); err != nil {
    log.Println(err)
}
```

# おわりに
