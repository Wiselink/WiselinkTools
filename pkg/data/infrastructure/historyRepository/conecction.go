package historyRepository

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (h *HistoricRepositoryObj) GetConeccion() (*mongo.Database, *mongo.Client, error) {

	uri := os.Getenv("MONGO_URL")

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		// spec 0006: cerrar el client conectado antes de salir — sin esto, cada intento
		// con Mongo caído filtraba un client (y sus goroutines) que nunca se desconectaba.
		_ = client.Disconnect(context.Background())
		return nil, nil, err
	}

	database := client.Database("history")

	// Ensure index on companyToken for faster lookups
	idx := mongo.IndexModel{Keys: bson.D{{Key: "companytoken", Value: 1}}}
	if _, err := database.Collection("activities").Indexes().CreateOne(ctx, idx); err != nil {
		_ = client.Disconnect(context.Background())
		return nil, nil, err
	}

	return database, client, nil
}
