package mongoModels

type ActivityMessage struct {
	CompanyToken string     `json:"companyToken"`
	ContactToken string     `json:"contactToken"`
	Activity     Activities `json:"activity"`
}
