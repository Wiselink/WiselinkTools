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
	}[s]
}

func (s Status) Index() int {
	return int(s)
}
