package mongodb

import (
	"fmt"
	"testing"
)

var uriCases = []string{"mongodb://localhost:27017", "mongodb://139.199.36.63:27017"}

func TestMongoInit(t *testing.T) {
	for _, uri := range uriCases {
		err := MongoInit(uri)
		if err != nil {
			t.Fatalf("err: %v", err)
			return
		}
	}

}

func TestDataBase(t *testing.T) {
	for _, uri := range uriCases {
		err := MongoInit(uri)
		if err != nil {
			t.Fatalf("err: %v", err)
			return
		}
		db := DataBase("test")
		if db != nil {
			fmt.Printf("db: %v", db)
		} else {
			t.Fatal("err: create or get db failed")
		}
	}

}

func TestCollection(t *testing.T) {
	for _, uri := range uriCases {
		err := MongoInit(uri)
		if err != nil {
			t.Fatalf("err: %v", err)
			return
		}
		collection := Collection("test", "test")
		if collection != nil {
			fmt.Printf("collection: %v", collection)
		} else {
			t.Fatal("err: create or get collection failed")
		}
	}
}
