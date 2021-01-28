package Repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateContext () context.Context
	CreateClientDatabase(ctx context.Context) *mongo.Client
}
