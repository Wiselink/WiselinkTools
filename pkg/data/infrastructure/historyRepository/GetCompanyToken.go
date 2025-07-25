package historyRepository

import (
	"context"
	"github.com/Wiselink/WiselinkTools/pkg/domain/mongoModels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (h HistoricRepositoryObj) GetCompanyByToken(db *mongo.Database, token string) (*mongoModels.Company, error) {
	// Establecer un tiempo límite de 10 segundos para la conexión a la base de datos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Buscar el token de la empresa
	var result mongoModels.Company
	err := db.Collection("companies").FindOne(ctx, bson.M{"token": token}).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
