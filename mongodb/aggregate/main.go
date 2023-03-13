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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.108.170:27017"))
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	var res []bson.D
	foods := mdb.Database("inventory").Collection("food")
	cursor, err := foods.Aggregate(ctx, bson.A{bson.D{{"$match",bson.D{{"price", bson.D{{"$gt",5.00},},},},},},bson.D{{"$group",bson.D{{"_id","$name"}}},},})
	if err != nil {
		fmt.Println("aggregate error:", err)
		return
	}
	err = cursor.All(ctx, &res)
	if err != nil {
		fmt.Println("cursor.All error:", err)
		return
	}
	fmt.Println("res:", res)
}