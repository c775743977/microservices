package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var coll *mongo.Collection

	// Insert documents {name: "Alice"} and {name: "Bob"}.
	// Set the Ordered option to false to allow both operations to happen even
	// if one of them errors.
	docs := []interface{}{
		bson.D{{"name", "Alice"}},
		bson.D{{"name", "Bob"}},
	}
	opts := options.InsertMany().SetOrdered(false)
	res, err := coll.InsertMany(context.TODO(), docs, opts)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("inserted documents with IDs %v\n", res.InsertedIDs)
}