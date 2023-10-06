package function

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB interface {
	Init() error
	CreateUser(user *User) (string, error)
	GetUser() ([]User, error)
	GetUserBy(email string) (User, error)
	DeleteUserBy(email string) error
}

type MongoDBConfig struct {
	URI      string
	Database string
}

type MongoDB struct {
	Client   *mongo.Client
	Database *mongo.Database
	Mu       sync.Mutex
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
	var err error
	// err = db.Database.Collection("users").Drop(context.Background())
	// if err != nil {
	// 	log.Println("Error dropping collection:", err)
	// }

	collection := db.Database.Collection("users")
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}},
	}
	_, err = collection.Indexes().CreateOne(context.TODO(), indexModel)
	return err
}

func (db *MongoDB) CreateUser(user *User) (string, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	collection := db.Database.Collection("users")
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Println("Error creating user:", err)
		return "", err
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (db *MongoDB) GetUser() ([]User, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	res, err := db.Database.Collection("users").Find(context.Background(), bson.D{})
	if err != nil {
		log.Println("Error getting users:", err)
		return nil, err
	}
	defer res.Close(context.Background())

	var users []User
	for res.Next(context.Background()) {
		var user User
		err := res.Decode(&user)
		if err != nil {
			log.Println("Error decoding user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (db *MongoDB) GetUserBy(email string) (User, error) {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	res, err := db.Database.Collection("users").Find(context.Background(), bson.D{
		{Key: "email", Value: email},
	})
	if err != nil {
		return User{}, err
	}

	var user User
	for res.Next(context.Background()) {
		err := res.Decode(&user)
		if err != nil {
			log.Println("Error decoding user:", err)
			return User{}, err
		}
	}

	if user.Email == "" {
		return User{}, fmt.Errorf("user not found")
	}

	return user, nil
}

func (db *MongoDB) DeleteUserBy(email string) error {
	db.Mu.Lock()
	defer db.Mu.Unlock()

	_, err := db.Database.Collection("users").DeleteOne(context.Background(), bson.D{
		{Key: "email", Value: email},
	})

	if err != nil {
		return err
	}

	return nil
}
