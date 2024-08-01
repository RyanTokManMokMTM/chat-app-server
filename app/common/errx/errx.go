package errx

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	code    InternalCode `json:"code"`
	message string       `json:"message"`
}

type ErrorResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewCustomError(code InternalCode, message string) *CustomError {
	return &CustomError{
		code:    code,
		message: message,
	}
}

func NewCustomErrMessage(message string) *CustomError {
	return &CustomError{
		code:    SERVER_COMMON_ERROR,
		message: message,
	}
}

func NewCustomErrCode(code InternalCode) *CustomError {
	return &CustomError{
		code:    code,
		message: MapErrMsg(code),
	}
}

func (err *CustomError) Error() string {
	return fmt.Sprintf("Code %v, Message: %v", err.code, err.message)
}

func (err *CustomError) GetCode() InternalCode {
	return err.code
}

func (err *CustomError) GetMessage() string {
	return err.message
}

func (err *CustomError) StatusCode() int {
	switch err.code {
	case SUCCESS:
		return http.StatusOK
	case REQ_PARAM_ERROR:
		return http.StatusBadRequest

	case SERVER_COMMON_ERROR:
		fallthrough
	case TOKEN_GENERATE_ERROR:
		fallthrough
	case DB_ERROR:
		fallthrough
	case DB_AFFECTED_ZERO_ERROR:
		return http.StatusInternalServerError

	case TOKEN_EXPIRED_ERROR:
		fallthrough
	case TOKEN_INVALID_ERROR:
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func (err *CustomError) ToJSON() *ErrorResp {
	return &ErrorResp{
		Code:    err.StatusCode(),
		Message: err.message,
	}
}
