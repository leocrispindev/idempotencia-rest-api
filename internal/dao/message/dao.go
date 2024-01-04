package message

import (
	"api-handle/internal/database"
	"api-handle/internal/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// TODO registro no banco
// TODO consulta no banco

var connection = database.GetConnection()

func AlreadyProcessed(key string) (bool, error) {
	// TODO implementar busca no banco
	collection := connection.Collection("message")

	result := collection.FindOne(context.TODO(), bson.M{"idempotenciaKey": key})

	var messageAlready model.Message

	err := result.Decode(&messageAlready)

	if err != nil {
		log.Println(err)
		return false, err
	}

	return messageAlready.IdempotenciaKey != "", err
}

func GetAllMessages() []interface{} {
	var result []interface{}

	coll := connection.Collection("message")

	cursor, err := coll.Find(context.TODO(), bson.D{})

	if err != nil {
		log.Println(err.Error())
		return result

	}

	cursor.All(context.TODO(), &result)

	return result
}

func SaveMessage(message model.Message) error {
	wc := writeconcern.Majority()
	txnOptions := options.Transaction().SetWriteConcern(wc)

	mongoSession, err := connection.Client().StartSession()
	if err != nil {
		return err
	}
	defer mongoSession.EndSession(context.TODO())

	coll := connection.Collection("message")

	_, err = mongoSession.WithTransaction(context.TODO(), func(ctx mongo.SessionContext) (interface{}, error) {
		log.Println("Transaction started")

		result, err := coll.InsertOne(ctx, message)
		if err != nil {
			log.Println("Error inserting into collection:", err)
			return nil, err
		}

		log.Printf("Document inserted: %v", result.InsertedID)

		return result, nil
	}, txnOptions)

	if err != nil {
		log.Println("Transaction error:", err)
		return fmt.Errorf("transaction failed: %w", err)

	}

	return nil

}
