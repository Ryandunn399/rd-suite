package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Interface defining the methods to connect, disconnect,
// and manipulate data in a MongoDB server
type MongoClient struct {
    Client      *mongo.Client
}

// Will create a client object and attempt to establish a connection
// with a mongo db server. This will process error handling and will
// panic if there is a problem exiting the program.
func CreateClient(url string) *MongoClient {

    log.Println("Attempting to establish connection...")

    ctx := context.TODO()
    clientOptions := options.Client().ApplyURI(url)

    client, err := mongo.Connect(ctx, clientOptions)

    if err != nil {
        panic(err)
    }

    err = client.Ping(ctx, nil)

    if err != nil {
        panic(err)
    }

    log.Println("Connection to mongo database established!")

    return &MongoClient{client}
}

// Will attempt to close the connection between our Mongo client
// and the MongoDB server the client is currently connected to.
func Disconnect(mc *MongoClient) {
    ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    mc.Client.Disconnect(ctx)
}

// Will retrieve a collection of the given MongoDB client based on the 
// database and collection name passed as arguments.
func GetCollection(mc *MongoClient, db string, collection string) *mongo.Collection {
    _, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
    defer cancel()
    return mc.Client.Database(db).Collection(collection)
}
