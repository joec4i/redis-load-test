package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/thanhpk/randstr"
	"math/rand"
	"strconv"
	"time"
)

func main()  {
	connections := flag.Int("c", 10, "The number of clients to start")
	addr := flag.String("addr", "10.133.136.177:63793", "The redis address")

	flag.Parse()

	fmt.Printf("Will create %d clients\n", *connections)

	for i := 0; i < *connections; i ++ {
		go func(clientNumber int) {
			var ctx = context.Background()

			for {
				rdb := redis.NewClient(&redis.Options{
					Addr:     *addr,
					Password: "", // no password set
					DB:       0,  // use default DB
				})

				time.Sleep(50 * time.Millisecond)

				readKey := "test:key:" + strconv.Itoa(rand.Intn(5000))
				writeKey := "test:key:" + strconv.Itoa(rand.Intn(5000))

				val, err := rdb.Get(ctx, readKey).Result()
				if err != nil {
					fmt.Println("#", clientNumber, "reading", readKey, ":", err)
					continue
				} else {
					fmt.Println("#", clientNumber, "read", readKey, ":", len(val))
				}

				value := randstr.Hex(rand.Intn(1024))

				err = rdb.Set(ctx, writeKey, value, 0).Err()
				if err != nil {
					fmt.Println("#", clientNumber, "error writing", writeKey, ":", err)
				} else {
					fmt.Println("#", clientNumber, "write", writeKey, ":", len(value))
				}

				err = rdb.Close()

				if err != nil {
					fmt.Println("#", clientNumber, "error disconnecting from redis:", err)
				}
			}
		}(i)
	}

	time.Sleep(2 * time.Hour)
}
