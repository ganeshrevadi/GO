package controller

import "go.mongodb.org/mongo-driver/mongo/options"

const connectString = "mongodb+srv://ganeshrevadi16:Grganesh16@@cluster0.dltvwo2.mongodb.net/"
const dbName = "netflix"
const collectionName = "watchlists"

func init() {
	//Client Option
	clientOption := options.Client().ApplyURI(connectString)
}
