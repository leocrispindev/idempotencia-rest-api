package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const uri = "mongodb://admin:admin1@localhost"

var connection *mongo.Database

func Init() {
	//open()
}

func open() {
	versionAPI := options.ServerAPI(options.ServerAPIVersion1)

	options := options.Client().ApplyURI(uri).SetServerAPIOptions(versionAPI)

	client, err := mongo.Connect(context.TODO(), options)

	fmt.Println(client.ListDatabaseNames(context.TODO(), bson.D{}))

	if err != nil {
		panic(err)
	}

	connection = client.Database("messages")
	fmt.Println(connection.Client().Ping(context.TODO(), &readpref.ReadPref{}))
	//fmt.Println(connection.Client().ListDatabaseNames(context.TODO(), bson.D{}))
}

func GetConnection() *mongo.Database {
	if connection == nil {
		open()
	}

	return connection
}
