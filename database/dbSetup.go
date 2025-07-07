package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "balh bhaaklashkdha"

const dbName = "ecommerce"
const UserColName = "users"
const ProductColName = "products"

func DBSet() *mongo.Client {
	clientOptions := options.Client().ApplyURI(connectionString)
	
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("failed to conenct to database")
		log.Fatal(err)
	}
	fmt.Println("successgully conencted to database")
	
	return client
}
var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection * mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection * mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}
