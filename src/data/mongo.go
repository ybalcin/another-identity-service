package data

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CONNECTED = "Successfully connected to database: %v"

type (
	MongoConfig struct {
		Uri      string
		Database string
	}

	MongoStore struct {
		Db      *mongo.Database
		Session *mongo.Client
	}
)

//	MgoConfig keeps mongo config
var MgoConfig MongoConfig

//	mgoStore keeps mongo db and session
var mgoStore MongoStore

//	InitMongo initializes mongo store
func InitMongo() {
	session, database := connect()
	mgoStore = MongoStore{
		Db:      database,
		Session: session,
	}
}

func connect() (*mongo.Client, *mongo.Database) {
	var connectOnce sync.Once
	var session *mongo.Client
	var database *mongo.Database
	connectOnce.Do(func() {
		session, database = connectToMongo()
	})

	return session, database
}

func connectToMongo() (*mongo.Client, *mongo.Database) {
	var err error
	var session *mongo.Client
	session, err = mongo.NewClient(options.Client().ApplyURI(MgoConfig.Uri))
	if err != nil {
		log.Fatalf("[log_mongo_connectomongo_newclient]: %s", err)
	}

	if err = session.Connect(context.Background()); err != nil {
		log.Fatalf("[log_mongo_connectomongo_connect]: %s", err)
	}

	database := session.Database(MgoConfig.Database)
	log.Printf(CONNECTED, MgoConfig.Database)

	return session, database
}