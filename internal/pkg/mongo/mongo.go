package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConfig struct {
	Host     string `env:"MONGO_HOST" envDefault:"localhost"`
	Port     string `env:"MONGO_PORT" envDefault:"27017"`
	User     string `env:"MONGO_USER" envDefault:"user"`
	Password string `env:"MONGO_PASSWORD" envDefault:"pass"`
	Database string `env:"MONGO_DATABASE" envDefault:"test"`
}

// Connect Init new Mongo client
func Connect(config MongoConfig) mongo.Client {
	uri := "mongodb://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/" + config.Database

	//context will time out after 30s
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	log.Printf("Mongo client connected.")
	return *mongoClient
}

// Disconnect Mongo client. Should be used as defer with Connect
func Disconnect(mongoClient *mongo.Client) {
	//context will time out after 30s
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	if err := mongoClient.Disconnect(ctx); err != nil {
		panic(err)
	}

	log.Printf("Mongo client disconnected.")
}
