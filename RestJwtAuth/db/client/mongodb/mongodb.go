package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewClient(ctx context.Context, host, port string) (db *mongo.Client, err error) {
	mongoDBURI := fmt.Sprintf("mongodb://%s:%s", host, port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDBURI))
	if err != nil {
		return nil, err
	}

	return client, nil
}

func Disconnect(client *mongo.Client, ctx context.Context) error {
	return client.Disconnect(ctx)
}
