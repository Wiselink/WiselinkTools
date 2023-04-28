package status

import (
	"WiselinkTools/pkg/domain/response"
	"net/http"
)

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
	case response.SuccessfulSearch, response.SuccessfulUpdate, response.SuccessfulDeletion, response.NotFound, response.NothingToDelete:
		(*w).WriteHeader(http.StatusOK)
		(*w).Write(responseWriter)
		return
	case response.SuccessfulCreation:
		(*w).WriteHeader(http.StatusCreated)
		(*w).Write(responseWriter)
		return
	case response.InternalServerError, response.DBQueryError, response.DBRowsAffectedError, response.DBExecutionError, response.DBRowsError, response.DBScanError, response.DBInitError, response.DecodeError, response.LastRowIdError, response.CreationFailure, response.ErrorMappingData:
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
	case response.Unauthorized:
		(*w).WriteHeader(http.StatusUnauthorized)
		(*w).Write([]byte(responseWriter))
		return
	case response.Conflict, response.NoRowsAffected:
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
