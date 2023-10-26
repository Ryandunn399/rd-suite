package bootstrap

import "github.com/ryandunn399/rd-suite/mongo"

type Application struct {
    Env                 *Env
    MongoClient         *mongo.MongoClient
}

func App() Application {
    app := &Application{}
    app.Env = NewEnv()
    app.MongoClient = EstablishConnection(app.Env)
    return *app
}


func (app *Application) DisconnectClient() {
    CloseConnection(app.MongoClient)
}
