package obj

import "encoding/json"

func UnmarshalErrorResponse(data []byte) (ErrorResponse, error) {
	var r ErrorResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
