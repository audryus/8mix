package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	client *mongo.Client
	db     *mongo.Database
}

func New() (*Mongo, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	url := "mongodb://admin:VRuAd2Nvmp4ELHh5@localhost:%s/"
	//port := os.Getenv("8MIX_MONGO_PORT")

	clientOptions := options.Client().ApplyURI(fmt.Sprintf(url, "27017"))
	clientOptions.SetServerAPIOptions(serverAPI)
	clientOptions.SetConnectTimeout(3 * time.Second)
	clientOptions.SetSocketTimeout(30 * time.Second)
	clientOptions.SetServerSelectionTimeout(15 * time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database("test")

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		panic(err)
	}

	fmt.Println("mongo connected")

	return &Mongo{
		client: client,
		db:     db,
	}, nil
}

func (m *Mongo) Collection(col string) *mongo.Collection {
	return m.db.Collection(col)
}

func (m *Mongo) Close() error {
	return m.client.Disconnect(context.TODO())
}
