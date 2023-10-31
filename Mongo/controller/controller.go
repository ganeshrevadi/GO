package controller

import (
	"context"
	"fmt"
	"ganeshrevadi/GO/Mongo/model"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString = "mongodb+srv://ganeshrevadi16:Grganesh16@@cluster0.dltvwo2.mongodb.net/"
const dbName = "netflix"
const collectionName = "watchlists"

var collection *mongo.Collection

func init() {
	//Client Option
	clientOption := options.Client().ApplyURI(connectString)

	//connect mongoDB
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)

	//collection instance
	fmt.Println("Collection instance created!")

}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	fmt.Println(inserted)
}
