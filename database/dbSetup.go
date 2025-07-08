package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "ecommerce"
const UserColName = "Users"
const ProductColName = "Products"

func DBSet() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("MONGO_DB_CONN")
	if connectionString == "" {
		log.Fatal("MONGO_DB_CONN is empty or not found in .env")
	}

	fmt.Println("MONGO_DB_CONN:", connectionString)

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("failed to connect to database")
		log.Fatal(err)
	}
	fmt.Println("successfully connected to database")

	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database(dbName).Collection(collectionName)
	return collection
}