package main

import (
	"fmt"
	"math/rand"
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"os"
)

func query_db_with_topic(topic string) []bson.M {
	var links []bson.M

	//TODO: Research Go Context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(build_connection_string()))
	

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	//CHANGE THIS LINE
	collection := client.Database("saga").Collection("links")
	//collection := client.Database(os.Getenv("DB_NAME")).Collection("links")
	
	cur, err := collection.Find(ctx, bson.M{"topic": topic})
	if err != nil { 
		fmt.Println(err)
		return links
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &links); err != nil {
		fmt.Println(err)
		return links
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return links
	}
	return links
}

func get_topics(category string) []bson.M {
	var topics []bson.M
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(build_connection_string()))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()

	//CHANGE THIS LINE
	collection := client.Database("saga").Collection("topics")
	//collection := client.Database(os.Getenv("DB_NAME")).Collection("topics")

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil { 
		fmt.Println(err)
		return topics
	}
	defer cur.Close(ctx)

	if err = cur.All(ctx, &topics); err != nil {
		fmt.Println(err)
		return topics
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
		return topics
	}
	return topics
}

// TO DO
func connect_to_db(connection_string string) {
	//
}

func build_connection_string() string {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_USER_PASS")
	db_name := os.Getenv("DB_NAME")
	endpoint := os.Getenv("DB_ENDPOINT")

	return "mongodb://" +username + ":" + password + "@" + endpoint + "/" + db_name
}

func recommend_randon_number(highest int) int {
	return rand.Intn(highest)
}