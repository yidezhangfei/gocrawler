package mongodb

import (
	"fmt"
	"testing"
)

func TestMongoInit(t *testing.T) {
	err := MongoInit()
	if err != nil {
		t.Fatalf("err: %v", err)
		return
	}
}

func TestDataBase(t *testing.T) {
	err := MongoInit()
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

func TestCollection(t *testing.T) {
	err := MongoInit()
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
