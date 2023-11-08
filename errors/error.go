package todoError

import "net/http"

var (
	InvalidStatus = NewAppError(http.StatusBadRequest, 1000)
)

var ErrorMap = map[int]string{
	1000: "Invalid status",
}

type TodoError struct {
	Code       int    `json:"code"`
	HttpStatus int    `json:"httpStatus"`
	Message    string `json:"message"`
}

func NewAppError(httpStatus int, code int) TodoError {
	return TodoError{
		Code:       code,
		HttpStatus: httpStatus,
		Message:    ErrorMap[code],
	}
}

func (e TodoError) Error() string {
	return e.Message
}
func (e TodoError) GetHttpStatus() int {
	return e.HttpStatus
}
func (e TodoError) GetCode() int {
	return e.Code
}
