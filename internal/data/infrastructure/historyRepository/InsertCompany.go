package historyRepository

import (
	"context"
	"github.com/Wiselink/WiselinkTools/pkg/domain/mongoModels"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (h HistoricRepositoryObj) InsertCompany(db *mongo.Database, company *mongoModels.Company) (*mongo.InsertOneResult, error) {
	// Establecer un tiempo límite de 10 segundos para la conexión a la base de datos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Insertar la empresa en la colección
	result, err := db.Collection("companies").InsertOne(ctx, company)
	if err != nil {
		return nil, err
	}

	return result, nil
}
