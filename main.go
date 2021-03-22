package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/rishi-org-stack/ItsaboutYou-b/model"
)
 
func connect() *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://rishijha1709:rishijha1709@cluster0.2lkrw.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database("Itsaboutyou_1").Collection("Itsaboutyou_1")
	// defer client.Disconnect(ctx)
	return coll

}

type person struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name,omitempty"`
}

func insert() {
	var p person
	// p.Name = "rishi"
	coll := connect()
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	i, e := coll.InsertOne(ctx, p)
	fmt.Println(e)
	fmt.Println(i.InsertedID)
}

func get(){
	var p person
	var f person
	f.Name ="rishi"
	coll := connect()
	filter := bson.D{primitive.E{"name",f.Name}}
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Second)
	r:= coll.FindOne(ctx,filter).Decode(&p)
	fmt.Println(r)
	fmt.Println(p)
}
func main() {
	db := &model.DB{

	}
	db = model.Instantiate()
	db = db.Connect()
	p := &person{}
	p.Name = "jhaji"
	db.Type =p
	db.CreateDb("okgettingit")
	db.LinkToCollection("jihaja")
	db.Insert()
	// get()
}
