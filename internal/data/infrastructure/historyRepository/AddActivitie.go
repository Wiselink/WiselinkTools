package historyRepository

import (
	"context"
	"github.com/Wiselink/WiselinkTools/pkg/domain/mongoModels"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

// AddActivity stores a new activity into the global "activities" collection.
// The document keeps references to the company and contact tokens and embeds
// the activity fields inline. Each company document remains under the
// "companies" collection.
func (h *HistoricRepositoryObj) AddActivity(db *mongo.Database, companyToken, contactToken string, activity mongoModels.Activities) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ensure company document exists in "companies" collection.
	if _, err := h.GetCompanyByToken(db, companyToken); err != nil {
		if err == mongo.ErrNoDocuments {
			company := &mongoModels.Company{Token: companyToken}
			if _, err = h.InsertCompany(db, company); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// Insert activity in the global "activities" collection.
	collection := db.Collection("activities")

	// Flatten activity fields at the document root.
	record := struct {
		CompanyToken           string `bson:"companytoken"`
		ContactToken           string `bson:"contacttoken"`
		mongoModels.Activities `bson:",inline"`
	}{
		CompanyToken: companyToken,
		ContactToken: contactToken,
		Activities:   activity,
	}

	_, err := collection.InsertOne(ctx, record)
	return err
}
