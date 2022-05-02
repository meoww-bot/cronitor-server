package db

import (
	"context"
	"cronitor-server/lib"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetMonitorsCount() (int64, error) {

	client := Conn()

	MonitorCollection := client.Database("cronitor").Collection("Monitor")

	filter := bson.D{}

	count, err := MonitorCollection.CountDocuments(context.TODO(), filter)

	return count, err
}

func GetAllMonitors(page string, monitors []lib.Monitor) ([]lib.Monitor, error) {

	client := Conn()

	MonitorCollection := client.Database("cronitor").Collection("Monitor")

	opts := options.Find()

	opts.SetLimit(40)

	pageInt, _ := strconv.ParseInt(page, 10, 64)

	opts.SetSkip((pageInt - 1) * 50)

	cur, err := MonitorCollection.Find(context.TODO(), bson.D{}, opts)

	if err != nil { //hndling with global error
		return monitors, err //exit
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var row lib.Monitor
		err := cur.Decode(&row)

		if err != nil {
			return monitors, err
		}

		monitors = append(monitors, row)
	}

	return monitors, err
}

// func AddEvent(event lib.Event) (*mongo.InsertOneResult, error) {

// 	client := Conn()

// 	EventCollection := client.Database("cronitor").Collection("Event")

// 	result, err := EventCollection.InsertOne(context.TODO(), event)

// 	return result, err
// }

func AddUser(user lib.User) (*mongo.InsertOneResult, error) {

	client := Conn()

	UserCollection := client.Database("cronitor").Collection("User")

	result, err := UserCollection.InsertOne(context.TODO(), user)

	return result, err
}

func AddInvocation(invocation lib.Invocation) (*mongo.InsertOneResult, error) {

	client := Conn()

	InvocationCollection := client.Database("cronitor").Collection("Invocation")

	result, err := InvocationCollection.InsertOne(context.TODO(), invocation)

	return result, err
}

func GetRunningInvocation(api_key string, code string, series string, invocation *lib.Invocation) error {

	client := Conn()

	InvocationCollection := client.Database("cronitor").Collection("Invocation")

	filter := bson.M{"state": "running", "api_key": api_key, "code": code, "series": series}

	err := InvocationCollection.FindOne(context.TODO(), filter).Decode(&invocation)

	return err
}

func UpdateRunningInvocation(api_key string, code string, series string, updateInvocation primitive.M) (*mongo.UpdateResult, error) {

	client := Conn()

	InvocationCollection := client.Database("cronitor").Collection("Invocation")

	filter := bson.M{"state": "running", "api_key": api_key, "code": code, "series": series}

	result, err := InvocationCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": updateInvocation})

	return result, err
}

func UpdateMonitorStatus(code string, updateMonitorStatus primitive.M) (*mongo.UpdateResult, error) {

	client := Conn()

	MonitorCollection := client.Database("cronitor").Collection("Monitor")

	filter := bson.M{"code": code}

	result, err := MonitorCollection.UpdateOne(context.TODO(), filter, bson.M{"$set": updateMonitorStatus})

	return result, err
}

func GetUser(Username string, user *lib.User) error {

	client := Conn()

	UserCollection := client.Database("cronitor").Collection("User")

	err := UserCollection.FindOne(context.TODO(), bson.M{"username": Username}).Decode(&user)

	return err
}
