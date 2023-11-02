package response

type Status int

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL
const (
	InternalServerError Status = iota
	DBQueryError
	DBExecutionError
	DBRowsError
	DBScanError
	LastRowIdError
	SuccessfulCreation
	SuccessfulSearch
	SuccessfulDeletion
	SuccessfulUpdate
	CreationFailure
	BadRequest
	Unknown
	ErrorMappingData
	NothingToDelete
	NotFound
	Conflict
	Unauthorized
	RequestEntityTooLarge
	UnsupportedMediaType
	NoResults
	DBInitError
	DBRowsAffectedError
	DecodeError
	BuilderResponseError
	NoContent
	NoRowsAffected
	UpdateFailure
	FormatInvalid
	ItemAlreadyShared
	LineHasReferences
	ItemIsInterested
	NameInUse
	NumberAndVersionAlreadyExists
	BudgetInUse
	AmountNotNull
	BusinessClosedOrLost
	ActiveStageInBusiness
	ActivePipelineInBusiness
	DontHaveLineAndPipeline
	DontHaveLine
	DontHavePipeline
	AnotherEditionIsActive
	FirstCreateASchedule
	Authorized
	AccreditationInUse
	ContactDataAlreadyExists
	ContactDataDoesntExist
	EmailAlreadyExists
	DoesntMatch
	UserDoesntExist
	Found
	NothingToUpdate
	EmailDoesntExist
	SuccessfulLogin
	InvalidNameFormat
	InvalidLastNameFormat
	InvalidEmailFormat
	ErrorGeneratingUuid
	WrongType
	GRPCClientsError
	GRPCServerError
	NoBudgetProperties
	NoProfileProperties
	NoBusinessProperties
	NoEditableCounterpartProps
	WrongProperties
	WrongTypeOfProperties
	WithoutCondition
	OnePropertyByStage
)

// LOS STATUS DEBEN MANTENER SU ORDEN. LOS NUEVOS DEBEN SER AGREGADOS AL FINAL

func (s Status) String() string {
	return [...]string{
		"InternalServerError",
		"DBQueryError",
		"DBExecutionError",
		"DBRowsError",
		"DBScanError",
		"LastRowIdError",
		"SuccessfulCreation",
		"SuccessfulSearch",
		"SuccessfulDeletion",
		"SuccessfulUpdate",
		"CreationFailure",
		"BadRequest",
		"Unknown",
		"ErrorMappingData",
		"NothingToDelete",
		"NotFound",
		"Conflict",
		"Unauthorized",
		"RequestEntityTooLarge",
		"UnsupportedMediaType",
		"NoResults",
		"DBInitError",
		"DBRowsAffectedError",
		"DecodeError",
		"BuilderResponseError",
		"NoContent",
		"NoRowsAffected",
		"UpdateFailure",
		"FormatInvalid",
		"ItemAlreadyShared",
		"LineHasReferences",
		"ItemIsInterested",
		"NameInUse",
		"NumberAndVersionAlreadyExists",
		"BudgetInUse",
		"AmountNotNull",
		"BusinessClosedOrLost",
		"ActiveStageInBusiness",
		"ActivePipelineInBusiness",
		"DontHaveLineAndPipeline",
		"DontHaveLine",
		"DontHavePipeline",
		"AnotherEditionIsActive",
		"FirstCreateASchedule",
		"Authorized",
		"AccreditationInUse",
		"ContactDataAlreadyExists",
		"ContactDataDoesntExist",
		"EmailAlreadyExists",
		"DoesntMatch",
		"UserDoesntExist",
		"Found",
		"NothingToUpdate",
		"EmailDoesntExist",
		"SuccessfulLogin",
		"InvalidNameFormat",
		"InvalidLastNameFormat",
		"InvalidEmailFormat",
		"ErrorGeneratingUuid",
		"WrongType",
		"GRPCClientsError",
		"GRPCServerError",
		"NoBudgetProperties",
		"NoProfileProperties",
		"NoBusinessProperties",
		"NoEditableCounterpartProps",
		"WrongProperties",
		"WrongTypeOfProperties",
		"WithoutCondition",
		"OnePropertyByStage",
	}[s]
}

func (s Status) Index() int {
	return int(s)
}
