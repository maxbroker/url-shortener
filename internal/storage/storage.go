package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Storage struct {
	db *mongo.Collection
}

type StorageFile struct {
	id    int
	alias string
	url   string
}

func New(CollectionName string) (*Storage, error) {

	const op = "storage.mongoDB.New"

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("mydb").Collection(CollectionName)

	first_link := StorageFile{1, "alias", "https://www.youtube.com/watch?v=oEupPPSes2I&list=PLNkWIWHIRwMFJ-3-gI7GC5JDg1ivbIKNR&index=6"}

	InsertOneResult, err := collection.InsertOne(context.TODO(), first_link)

	if err != nil {
		return nil, fmt.Errorf("%s: %v", op, err)
	}

	fmt.Println("Inserted a single document: ", InsertOneResult.InsertedID)

	first_collection := Storage{collection}

	return &first_collection, nil
}

func main() {
	const op = "storage.main"

	storage, err := New("url-shortener")
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}
	_ = storage
}
