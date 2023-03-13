package main

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID string
	Username string
	UserID string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	mdb, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://192.168.108.170:27017"))
	if err != nil {
		fmt.Println("connect mongodb error:", err)
		return
	}
	var data bson.D
	books := mdb.Database("bookstore").Collection("books")
	cursor, _ := books.Find(ctx, bson.D{})
	for cursor.Next(ctx) {
		cursor.Decode(&data)
		fmt.Println("data:", data[6])
	}
}