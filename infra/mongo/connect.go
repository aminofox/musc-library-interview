package mongo

import (
	"context"
	"fmt"
	"log"
	"music-libray-management/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	*mongo.Database
}

func Connect(cfg *config.MongoEnv) (*MongoDB, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s",
		cfg.MongoDbUsername,
		cfg.MongoDbPassword,
		cfg.MongoDbHost,
		cfg.MongoDbPort,
		cfg.MongoDbOption,
	)

	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err)
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	return &MongoDB{
		Database: client.Database(cfg.MongoDbDatabase),
	}, nil
}
