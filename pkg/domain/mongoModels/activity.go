package mongoModels

type Option string

type History struct {
	Companies []Company
}

type Company struct {
	Token    string //token de la empresa
	Contacts []Contact
}

type Contact struct {
	ContactToken string       `json:"contactToken"`
	Activities   []Activities `json:"activities"`
}

type Activities struct {
	Types               string             `json:"types"`       // business, meeting, task
	GenericType         string             `json:"genericType"` // tipo de actividad ejemplo note, meeting, task
	CreateUser          User               `json:"createUser"`
	TypeActivity        ActivityType       `json:"typeActivity"`
	EmployeeToken       string             `json:"employeeToken"`
	IsFromClient        bool               `json:"isFromClient"` //el cliente  ejecuto la actividad
	UnixTime            int64              `json:"unixTime"`
	Data                DataActivities     `json:"dataHelper"`
	ContactDataResponse ActivitiesResponse `json:"data"`
	Point               int                `json:"point"`
}

type DataActivities struct {
	ReasonUpdateTokenBusiness string `json:"reasonUpdateToken"`
	LostBusinesstoken         string `json:"lostBusinessToken"`
	ReasonBlocked             string `json:"reasonBlocked"`
	StageTokenLast            string `json:"stageTokenLast"`
	StageTokenNow             string `json:"stageTokenNow"`
	TokenTaskObjet            string `json:"tokenTaskObjet"`
	TokenCommunicationObjet   string `json:"tokenCommunicationObjet"`
	TokenMeetingObjet         string `json:"tokenMeetingObjet"`
	StageToken                string `json:"stageToken"`
	MeetID                    int    `json:"meetId"`
	TokenMeeting              string `json:"tokenMeeting"`
	TokenBusiness             string `json:"tokenBusiness"`
	TokenTask                 string `json:"tokenTask"`
	TokenLine                 string `json:"tokenLine"`
	TokenProduct              string `json:"tokenProduct"`
	TokenService              string `json:"tokenService"`
	TokenEmployee             string `json:"tokenEmployee"`
	TokenCompany              string `json:"tokenCompany"`
	TokenNote                 string `json:"tokeNote"`
	TokenInteraction          string `json:"interactionToken"`
	TokenBudgetVersion        string `json:"tokenBudgetVersion"`
	TokenBudgetOption         string `json:"tokenBudgetOption"`
	TokenPreAccreditation     string `json:"tokenPreAccreditation"`
}

type Budget struct {
	Name    string `json:"name"`
	Number  string `json:"number"`
	Version string `json:"version"`
}

type ActivitiesResponse struct {
	Budget                Budget     `json:"budget"`
	User                  User       `json:"createUser"`
	Exposition            Exposition `json:"exposition"`
	TokenNote             string     `json:"tokenNote"`
	InteractionName       string     `json:"interactionName"`
	BusinessName          string     `json:"businessName"`
	TokenBusiness         string     `json:"tokenBusiness"`
	Stages                Stages     `json:"stages"`
	CommunicationToken    string     `json:"communicationToken"`
	StageName             string     `json:"stageName"`
	MeetingName           string     `json:"meetingName"`
	ObjectInteractionName string     `json:"objectInteractionName"`
	BudgetName            string     `json:"budgetName"`
	BudgetVersionNumber   string     `json:"budgetVersionNumber"`
	TaskName              string     `json:"taskName"`
	Item                  Item       `json:"item"`
}

type Item struct {
	Name  string `json:"name"`
	Token string `json:"token"`
}

type Stages struct {
	OldStage            Stage  `json:"oldStage"`
	NewStage            Stage  `json:"newStage"`
	Reason              string `json:"reasonId"`
	ReasonDescription   string `json:"reasonDescription" `
	TimeStampBlockedEnd string `json:"timeStampBlockedEnd"`
}

type Stage struct {
	Name      string `json:"name"`
	StageName string `json:"stageName"`
	Token     string `json:"token"`
}

type Exposition struct {
	Name    string `json:"name"`
	Edition string `json:"edition"`
}

type User struct {
	Name               string `json:"name"`
	Avatar             string `json:"avatar"`
	SectorORProfession string `json:"sectorORProfession"`
	CompanyName        string `json:"companyName"`
}

type ActivityTypeResponse struct {
	AmountActivities int          `json:"amountActivities"`
	TotalPoints      int          `json:"totalPoints"`
	Activities       []Activities `json:"activities"`
}

type TaskHistory struct {
	Token   string      `json:"token"`
	Name    string      `json:"name"`
	History *[]Historyt `json:"history"`
}

type Historyt struct {
	EmployeeToken string        `json:"employeeToken"`
	Name          string        `json:"name"`
	Avatar        string        `json:"avatar"`
	Activity      *LastActivity `json:"activity"`
}

type LastActivity struct {
	Date   string `json:"date"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
}
