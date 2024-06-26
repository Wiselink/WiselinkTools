package status

import (
	"encoding/json"

	"github.com/Wiselink/WiselinkTools/pkg/domain/response"
)

func ResponseBuilder(status int, message string, data interface{}) ([]byte, error) {
	response := response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	return marshalResponse, nil

}

func ResponseStaNdard(s response.Status, err error, siteError string) response.Response {
	return response.Response{
		Resp:  &s,
		Error: err,
	}
}
func ResponseCustom(status int, message string, data interface{}, err error, siteError string) response.Response {
	return response.Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   err,
		Site:    siteError,
	}
}
