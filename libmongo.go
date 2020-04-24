package libmongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type (
	// DBConfig is used to ...
	DBConfig struct {
		Host     string
		Port     int
		DBName   string
		Username string
		Password string
	}
)

// Connect is used to ...
func Connect(config DBConfig) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)))
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database(config.DBName), nil
}

// ConnectWithScram is used to ...
func ConnectWithScram(config DBConfig) (*mongo.Database, error) {
	credential := options.Credential{
		Username: config.Username,
		Password: config.Password,
	}
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		return nil, err
	}
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}
	return client.Database(config.DBName), nil
}
