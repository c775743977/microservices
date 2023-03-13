package main

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.108.170:27017"))
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	data := []interface{}{
		bson.D{
			{"name","coconut"},
			{"price",5.00},
		},
		bson.D{
			{"name","pizza"},
			{"price",15.00},
		},
	}
	foods := mdb.Database("inventory").Collection("food")
	res, err := foods.InsertMany(ctx, data)
	if err != nil {
		fmt.Println("insertOne error:", err)
		return
	}
	fmt.Println("res:", res) // res就是插入数据的_id
}