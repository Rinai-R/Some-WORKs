package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type User struct {
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

var ctx = context.Background()

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://mongo:mongo@192.168.195.129:27017"))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return
	}
	fmt.Println("mongodb connect success")

	usersCollection := client.Database("game").Collection("users")

	user := User{
		Name:     "Rinai",
		Password: "123456",
	}
	bsondata, err := bson.Marshal(user)
	if err != nil {
		panic(err)
	}
	result, err := usersCollection.InsertOne(ctx, bsondata)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Inserted a single document: %+v\n", result)
	//切片转换
	users := []User{
		{Name: "rinai", Password: "123456"},
		{Name: "clas", Password: "123123"},
	}
	var documents []interface{}
	for u := range users {
		documents = append(documents, users[u])
	}
	results, err := usersCollection.InsertMany(ctx, documents)
	if err != nil {
		panic(err)

	}
	fmt.Printf("Inserted multiple documents: %+v\n", results)
}
