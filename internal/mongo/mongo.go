package mongo

import (
	"context"
	"log"
	"time"

	"hexa-example-go/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect Init new Mongo client
func Connect(conf config.MongoConfig) mongo.Client {
	uri := "mongodb://" + conf.User + ":" + conf.Password + "@" + conf.Host + ":" + conf.Port + "/" + conf.Database

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
