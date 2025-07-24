package historyRepository

import (
	"context"
	"github.com/Wiselink/WiselinkTools/pkg/domain/mongoModels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

// GetHistory retrieves all activities of a company from the global
// "activities" collection. If contact is not empty, the results are filtered by
// contact token as well.
func (h HistoricRepositoryObj) GetHistory(db *mongo.Database, companyToken, contact string) ([]mongoModels.Activities, error) {
	var activities []mongoModels.Activities

	filter := bson.M{"companytoken": companyToken}
	if contact != "" {
		filter["contacttoken"] = contact
	}
	collection := db.Collection("activities")
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var result struct {
			CompanyToken           string `bson:"companytoken"`
			ContactToken           string `bson:"contacttoken"`
			mongoModels.Activities `bson:",inline"`
		}
		if err := cursor.Decode(&result); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		activities = append(activities, result.Activities)
	}

	return activities, cursor.Err()
}
