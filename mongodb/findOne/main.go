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
		fmt.Println("connect mongodb error:", err)
		 return
	}
	foods := mdb.Database("inventory").Collection("food")
	var res bson.D
	err = foods.FindOne(ctx, bson.M{}).Decode(&res)
	fmt.Println("res:", res[1].Value)
}