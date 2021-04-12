// Author: yangzq80@gmail.com
// Date: 2020-11-24
//
package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

var ctx = context.Background()

func TestCluster(t *testing.T) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"172.16.20.223:7001", "172.16.20.223:7002", "172.16.20.223:7003"},
	})
	rdb.Ping(ctx)
}

func TestClient(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err.Error())
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
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
