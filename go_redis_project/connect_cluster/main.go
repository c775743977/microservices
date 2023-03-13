package main

import (
	"github.com/redis/go-redis/v9"
	"fmt"
	"context"
)

func main() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"192.168.108.165:6381", "192.168.108.165:6388", "192.168.108.165:6383", "192.168.108.165:6384", "192.168.108.165:6385", "192.168.108.165:6386"},
	})
	ctx := context.Background()
	err := rdb.ForEachShard(ctx, func(ctx context.Context, shard *redis.Client) error {
		return shard.Ping(ctx).Err()
	})
	if err != nil {
		panic(err)
	}
	// res, err := rdb.ClusterInfo(ctx).Result()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println("res:", res)
	// }
	// res, err = rdb.ClusterNodes(ctx).Result()
	// if err != nil {
	// 	fmt.Println("error:", err)
	// } else {
	// 	fmt.Println("res:", res)
	// }
	err = rdb.Set(ctx, "k4", "v4", -1).Err()
	if err != nil {
		fmt.Println("error:", err)
	} 
	v, err := rdb.Get(ctx, "k3").Result()
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("v=", v)
	}
}