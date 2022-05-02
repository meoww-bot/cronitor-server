package db

import (
	"context"
	"cronitor-server/config"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {

	clientOptions := options.Client().ApplyURI(config.MongoURI) //.SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB ,err:" + err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")

}

func Conn() *mongo.Client {

	clientOptions := options.Client().ApplyURI(config.MongoURI) //.SetServerAPIOptions(serverAPIOptions)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Failed to connect to MongoDB ,err:" + err.Error())
		os.Exit(1)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
