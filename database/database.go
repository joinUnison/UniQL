package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Client struct {
	mongo *mongo.Client
}

type MongoConfig struct {
	uri string
}

type ClientConfig struct {
	mongo MongoConfig
}

func NewClient(config ClientConfig) *Client {
	clientOptions := options.Client().ApplyURI(config.mongo.uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		mongo: client,
	}
}
