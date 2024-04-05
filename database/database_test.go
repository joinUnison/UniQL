package database

import "testing"

func TestNewClient_Mongo_ConnectError(t *testing.T) {
	config := ClientConfig{}

	client := NewClient(config)

	if client != nil {
		t.Fatal("NewClient returned non-nil, expected a nil database client instance")
	}
}

func TestNewClient_Mongo_Success(t *testing.T) {
	config := ClientConfig{
		mongo: MongoConfig{
			uri: "mongodb://user:pass@sample.host:27017",
		},
	}

	client := NewClient(config)

	if client == nil {
		t.Fatal("NewClient returned nil, expected a database client instance")
	}

	if client.mongo == nil {
		t.Fatal("client.mongo is nil, expected a mongo.Client instance")
	}
}
