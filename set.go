package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

var collection *mongo.Collection
var ctx = context.TODO()

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("base").Collection("mongo")
}

func main() {
	app := &cli.App{
		Name:     "mongo",
		Usage:    "A simple CLI program to manage your tasks",
		Commands: []*cli.Command{},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
