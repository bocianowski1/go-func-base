package function

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface{}

type MongoDBConfig struct {
	URI      string
	Database string
}

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(config MongoDBConfig) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(config.URI)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("Could not connect to MongoDB")
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Println("Could not ping MongoDB")
		return nil, err
	}

	db := client.Database(config.Database)

	return &MongoDB{
		Client:   client,
		Database: db,
	}, nil
}

func (db *MongoDB) Init() error {
	collection := db.Database.Collection("users")
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
	}
	_, err := collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}
