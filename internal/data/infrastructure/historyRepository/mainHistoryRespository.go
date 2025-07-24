package historyRepository

import (
	"github.com/Wiselink/WiselinkTools/pkg/domain/mongoModels"
	"go.mongodb.org/mongo-driver/mongo"
)

type IHistoricRepository interface {
	GetHistory(db *mongo.Database, companyToken, contact string) ([]mongoModels.Activities, error)
	GetConeccion() (*mongo.Database, *mongo.Client, error)
	AddActivity(db *mongo.Database, companyToken, contactToken string, activity mongoModels.Activities) error
}

type HistoricRepositoryObj struct {
}

func GetProvider() IHistoricRepository {
	return &HistoricRepositoryObj{}
}
