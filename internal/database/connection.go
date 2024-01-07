package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const uri = "mongodb://127.0.0.1:27017/sample_messages?authSource=admin"

var connection *mongo.Database

func Init() {
	open()
}

func open() {

	// Configure as opções de conexão
	clientOptions := options.Client().ApplyURI(uri)

	// Conecte-se ao MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verifique a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	connection = client.Database("sample_messages")

	fmt.Println("Conectado ao MongoDB!")
}

func GetConnection() *mongo.Database {
	if connection == nil {
		open()
	}

	return connection
}
