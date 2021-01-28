package Repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoMutantRepository struct {

}

func (p *MongoMutantRepository) CreateContext () context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}

func (p *MongoMutantRepository) CreateClientDatabase (ctx context.Context) *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, _ := mongo.Connect(ctx, clientOptions)
	return client
}

