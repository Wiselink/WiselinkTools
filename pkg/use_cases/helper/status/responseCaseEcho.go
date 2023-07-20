package status

import (
	"net/http"

	"github.com/Wiselink/WiselinkTools/pkg/domain/response"

	"github.com/labstack/echo/v4"
)

func ResponseCaseEcho(c echo.Context, status response.Status, data interface{}) error {

	var respData interface{}

	if data != nil {
		respData = data
	} else {
		respData = echo.Map{}
	}

	switch status {
	case response.SuccessfulSearch, response.SuccessfulUpdate, response.SuccessfulDeletion, response.NotFound, response.NothingToDelete:
		return c.JSON(http.StatusOK, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.SuccessfulCreation:
		return c.JSON(http.StatusCreated, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.InternalServerError, response.DBQueryError, response.DBRowsAffectedError, response.DBExecutionError, response.DBRowsError, response.DBScanError, response.DBInitError, response.DecodeError, response.LastRowIdError, response.CreationFailure, response.UpdateFailure, response.ErrorMappingData:
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.NoContent:
		return c.NoContent(http.StatusNoContent)
	case response.BadRequest:
		return c.JSON(http.StatusBadRequest, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.Unauthorized:
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.Conflict, response.NoRowsAffected, response.LineHasReferences, response.ItemAlreadyShared, response.ItemIsInterested, response.FormatInvalid, response.NoResults,
		response.NameInUse, response.NumberAndVersionAlreadyExists, response.BudgetInUse, response.AmountNotNull, response.BusinessClosedOrLost, response.ActiveStageInBusiness,
		response.ActivePipelineInBusiness, response.DontHaveLineAndPipeline, response.DontHaveLine, response.DontHavePipeline:
		return c.JSON(http.StatusConflict, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.RequestEntityTooLarge:
		return c.JSON(http.StatusRequestEntityTooLarge, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	case response.UnsupportedMediaType:
		return c.JSON(http.StatusUnsupportedMediaType, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	default:
		status = response.Unknown
		return c.JSON(http.StatusNotImplemented, echo.Map{
			"status":  status.Index(),
			"message": status.String(),
			"data":    respData,
		})
	}
}
