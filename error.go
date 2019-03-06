package face

import "encoding/json"

type ApiError struct {
	InnerErr ErrorMsg `json:"error"`
}

func (e ApiError) Error() string {
	ret, _ := json.Marshal(e)
	return string(ret)
}

type ErrorMsg struct {
	Code       string `json:"code"`
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}
