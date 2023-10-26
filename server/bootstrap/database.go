package bootstrap

import (
	"fmt"
        "log"

	"github.com/ryandunn399/rd-suite/mongo"
)

// Establish a connection to the database utilizing
// local environmental variables
func EstablishConnection(env *Env) *mongo.MongoClient {

    host := env.Host
    port := env.Port
    user := env.User
    pass := env.Pass

    uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, pass, host, port)

    if user == "" || pass == "" {
        uri = fmt.Sprintf("mongodb://%s:%s", host, port)
    }
    
    log.Println(uri)

    client := mongo.CreateClient(uri)
    return client
}

// Attempt to disconnect the current client from the server.
func CloseConnection(client *mongo.MongoClient) {
    if client == nil {
        log.Fatal("Do not pass null client into paramters")
    }
    
    mongo.Disconnect(client)
}
