package model

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	uri    string
	client *mongo.Client
}

func CreateDb() *DB {
	Db := &DB{}
	Db.uri = "mongodb+srv://rishijha1709:rishijha1709@cluster0.2lkrw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	return Db
}

func (db *DB) Connect() *DB {
	client, err := mongo.NewClient(options.Client().ApplyURI(db.uri))
	if err != nil {
		fmt.Printf("error in connection : %v\n", err)
	}
	db.client = client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	db.client.Connect(ctx)
	defer cancel()
	return db
}
type person struct{
	ID primitive.ObjectID
	Name string
}
func (db *DB) Get() {
	var p person
	// var f person
	// f.Name ="rishi" 
	client := db.client
	coll := client.Database("Itsaboutyou_1").Collection("Itsaboutyou_1")
	filter := bson.D{{"name", "rishi"}}
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	r := coll.FindOne(ctx, filter).Decode(&p)
	fmt.Println(r)
	fmt.Println(p)
}
