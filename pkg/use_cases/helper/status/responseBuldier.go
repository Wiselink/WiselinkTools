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

func ResponseStaNdard(s response.Status) response.Response {
	return response.Response{
		Resp: &s,
	}
}
func ResponseCustom(status int, message string, data interface{}) response.Response {
	return response.Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func CompareSTatus(resp response.Response, status response.Status) bool {
	if resp.Resp != nil {
		if *resp.Resp != status {
			return true
		} else {
			return false
		}
	} else {

		if resp.Status < 300 {
			return false
		} else {
			return true
		}

	}
	return true

}
