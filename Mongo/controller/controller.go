package controller

import (
	"context"
	"fmt"
	"ganeshrevadi/GO/Mongo/model"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectString = "mongodb+srv://ganeshrevadi16:Grganesh16@@cluster0.dltvwo2.mongodb.net/"
const dbName = "netflix"
const collectionName = "watchlists"

var collection *mongo.Collection

func Init() {
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
	}

	fmt.Println("Inserted 1 Movie in db with id : ", inserted.InsertedID)
}

func updateOneMovie(moiveId string) {
	id, _ := primitive.ObjectIDFromHex(moiveId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"Watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count : ", result.ModifiedCount)
}

func DeleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Count : ", result.DeletedCount)
}

func DeleteAllMovie() {
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Deleted Count is : ", result.DeletedCount)
}
