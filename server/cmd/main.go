package main

import (
	"context"
	"fmt"

	"github.com/ryandunn399/rd-suite/bootstrap"
	"github.com/ryandunn399/rd-suite/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type Book struct {
    Title       string `bson:"title"`
    Author      string `bson:"author"`
}

func main() {
    
    app := bootstrap.App()

    db := app.MongoClient
    defer app.DisconnectClient()

    // Spike

    var result Book

    collection := mongo.GetCollection(db, "mydb", "books") 
    filter := bson.D{{ }}

    err := collection.FindOne(context.TODO(), filter).Decode(&result)

    if err != nil {
        fmt.Println("Error somewhere...")
    }

    fmt.Printf("%s and %s\n", result.Author, result.Title)
}
