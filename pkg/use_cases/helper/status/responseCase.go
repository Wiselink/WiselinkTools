package status

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Wiselink/WiselinkTools/pkg/domain/response"
)

func ResponseCasev2(w *http.ResponseWriter, r response.Response) {

	if r.Resp == nil {
		marshalResponse, err := json.Marshal(r)
		if err != nil {
			log.Println(" ResponseCasev2", err.Error())
		}
		(*w).WriteHeader(r.Status)
		(*w).Write(marshalResponse)
	} else {
		ResponseCase(w, *r.Resp, r.Data)
	}

}

func ResponseCase(w *http.ResponseWriter, status response.Status, data interface{}) {

	var responseWriter []byte
	var err error

	responseWriter, err = ResponseBuilder(status.Index(), status.String(), data)
	if err != nil {
		status = response.BuilderResponseError
		responseWriter, err = ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			(*w).WriteHeader(http.StatusInternalServerError)
			(*w).Write([]byte("500: Internal server error"))
			return
		}
		(*w).WriteHeader(http.StatusInternalServerError)
		(*w).Write([]byte(responseWriter))
		return
	}

	switch status {
	case response.SuccessfulSearch, response.SuccessfulUpdate, response.SuccessfulDeletion, response.NotFound, response.NothingToDelete, response.Authorized, response.ContactDataDoesntExist, response.ContactDataAlreadyExists, response.EmailAlreadyExists, response.DoesntMatch,
		response.Found, response.UserDoesntExist, response.NothingToUpdate, response.EmailDoesntExist, response.SuccessfulLogin, response.InvalidNameFormat, response.InvalidLastNameFormat, response.InvalidEmailFormat, response.ErrorGeneratingUuid, response.WrongType:
		(*w).WriteHeader(http.StatusOK)
		(*w).Write(responseWriter)
		return
	case response.SuccessfulCreation:
		(*w).WriteHeader(http.StatusCreated)
		(*w).Write(responseWriter)
		return
	case response.InternalServerError, response.DBQueryError, response.DBRowsAffectedError, response.DBExecutionError, response.DBRowsError, response.DBScanError, response.DBInitError, response.DecodeError, response.LastRowIdError, response.CreationFailure, response.UpdateFailure, response.ErrorMappingData,
		response.GRPCServerError, response.GRPCClientsError:
		(*w).WriteHeader(http.StatusInternalServerError)
		(*w).Write(responseWriter)
		return
	case response.NoContent:
		(*w).WriteHeader(http.StatusNoContent)
		(*w).Write([]byte(responseWriter))
		return
	case response.BadRequest:
		(*w).WriteHeader(http.StatusBadRequest)
		(*w).Write([]byte(responseWriter))
		return
	case response.Unauthorized, response.PermissionDenied:
		(*w).WriteHeader(http.StatusUnauthorized)
		(*w).Write([]byte(responseWriter))
		return
	case response.Conflict, response.NoRowsAffected, response.LineHasReferences, response.ItemAlreadyShared, response.ItemIsInterested, response.FormatInvalid, response.NoResults,
		response.NameInUse, response.NumberAndVersionAlreadyExists, response.BudgetInUse, response.AmountNotNull, response.BusinessClosedOrLost, response.ActiveStageInBusiness,
		response.ActivePipelineInBusiness, response.DontHaveLineAndPipeline, response.DontHaveLine, response.DontHavePipeline, response.AnotherEditionIsActive, response.FirstCreateASchedule,
		response.AccreditationInUse, response.WrongTypeOfProperties, response.OnePropertyByStage, response.NoBudgetProperties, response.NoBusinessProperties, response.NoProfileProperties,
		response.NoEditableCounterpartProps, response.WrongProperties, response.WithoutCondition:
		(*w).WriteHeader(http.StatusConflict)
		(*w).Write([]byte(responseWriter))
		return
	case response.RequestEntityTooLarge:
		(*w).WriteHeader(http.StatusRequestEntityTooLarge)
		(*w).Write([]byte(responseWriter))
		return
	case response.UnsupportedMediaType:
		(*w).WriteHeader(http.StatusUnsupportedMediaType)
		(*w).Write([]byte(responseWriter))
		return
	default:
		status = response.Unknown
		responseWriter, err = ResponseBuilder(status.Index(), status.String(), nil)
		if err != nil {
			(*w).WriteHeader(http.StatusInternalServerError)
			(*w).Write([]byte("500: Internal server error"))
			return
		}
		(*w).WriteHeader(http.StatusNotImplemented)
		(*w).Write([]byte(responseWriter))
		return
	}
}
