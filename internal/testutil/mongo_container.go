package testutil

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoContainer(ctx context.Context) (testcontainers.Container, *mongo.Client, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForLog("Waiting for connections").WithStartupTimeout(60 * time.Second),
	}

	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Printf("Error creating MongoDB container: %v\n", err)
		return nil, nil, err
	}

	host, err := mongoC.Host(ctx)
	if err != nil {
		log.Printf("Error getting host for MongoDB container: %v\n", err)
		return nil, nil, err
	}

	port, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		log.Printf("Error getting mapped port for MongoDB container: %v\n", err)
		return nil, nil, err
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s", host, port.Port())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("Error connecting to MongoDB: %v\n", err)
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Error pinging MongoDB: %v\n", err)
		return nil, nil, err
	}

	log.Printf("Successfully set up MongoDB container at %s\n", mongoURI)
	return mongoC, client, nil
}

func TeardownMongoContainer(ctx context.Context, container testcontainers.Container) error {
	return container.Terminate(ctx)
}
