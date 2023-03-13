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
	foods := mdb.Database("inventory").Collection("food")
	//bson.M会导入的文档数据顺序混乱，因为bson.M是无序的
	// res, err := foods.UpdateOne(ctx, bson.M{"name":"orange"}, bson.M{"$set":bson.M{"name":"orange_chen", "price":5.29, "birth":"1998-05-29"}})
	res, err := foods.UpdateOne(ctx, bson.M{"name":"orange"}, bson.D{{"$set", bson.D{{"name", "orange_chen"}, {"price", 5.29}, {"birth","1998-05-29"},},}})
	if err != nil {
		fmt.Println("updateOne error:", err)
		return
	}
	fmt.Println("res:", res)
	var result bson.M
	err = foods.FindOne(ctx, bson.M{"name":"orange_chen"}).Decode(&result)
	if err != nil {
		fmt.Println("FindOne error:", err)
		return
	}
	fmt.Println("result:", result)
}