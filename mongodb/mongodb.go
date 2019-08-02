package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

var MongoClient *mongo.Client

func MongoInit() error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb://localhost:27017"
	// 注意 := 会覆盖变量，包括变量范围
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	MongoClient = client
	err = MongoClient.Ping(ctx, readpref.Primary())
	return err
}

func DataBase(dbName string) *mongo.Database {
	db := MongoClient.Database(dbName)
	return db
}

func Collection(dbName string, collectionName string) *mongo.Collection {
	return MongoClient.Database(dbName).Collection(collectionName)
}
