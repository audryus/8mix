package container

import (
	"context"
	"fmt"

	"github.com/audryus/8mix/http/pkg/logger"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
)

var mongoContainer *mongodb.MongoDBContainer

func InitMongo() {
	fmt.Println("initializing mongo")
	ctx := context.Background()

	logger := logger.New()

	container, err := mongodb.Run(ctx, "mongo:6",
		testcontainers.CustomizeRequest(testcontainers.GenericContainerRequest{
			Reuse: true,

			ContainerRequest: testcontainers.ContainerRequest{
				Name: "8mix-mongo",

				ExposedPorts: []string{"27017:27017"},
				Env: map[string]string{
					"MONGO_INITDB_ROOT_USERNAME": "admin",
					"MONGO_INITDB_ROOT_PASSWORD": "VRuAd2Nvmp4ELHh5",
					"MONGO_INITDB_DATABASE":      "test"},
			},
		}))

	if err != nil {
		logger.Error("connect to mongo: %s", err)
		panic(err)
	}

	port, err := container.MappedPort(ctx, "27017")

	if err != nil {
		fmt.Println("Could not obtain cockroach port")
		panic(err)
	}

	fmt.Printf("mongo up and running port %s...\n\n", port.Port())
}

func TerminateMongo(logger *logger.Log) {
	if mongoContainer != nil {
		if err := mongoContainer.Terminate(context.Background()); err != nil {
			logger.Error("terminate mongo error: %s", err)
		}
	}
}
