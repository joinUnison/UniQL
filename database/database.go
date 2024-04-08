package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Client struct {
	Mongo *mongo.Client
}

type MongoConfig struct {
	URI string
}

type ClientConfig struct {
	Mongo MongoConfig
}

func NewClient(config ClientConfig) *Client {
	clientOptions := options.Client().ApplyURI(config.Mongo.URI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil
	}

	return &Client{
		Mongo: client,
	}
}
