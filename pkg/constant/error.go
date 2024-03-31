package constant

import (
	"errors"
	"net/http"
)

const (
	ErrSQLUniqueViolation = "23505"
)

var (
	ErrEmailAlreadyRegistered    = &ErrWithCode{HTTPStatusCode: http.StatusConflict, Message: "email already registered"}
	ErrUsernameOrPasswordInvalid = &ErrWithCode{HTTPStatusCode: http.StatusBadRequest, Message: "username or password invalid"}
	ErrUserNotFound              = &ErrWithCode{HTTPStatusCode: http.StatusNotFound, Message: "user not found"}
	ErrInvalidUUID               = errors.New("invalid uuid length or format")
	ErrAccessForbidden           = &ErrWithCode{HTTPStatusCode: http.StatusForbidden, Message: "access forbidden"}
	ErrBalanceInsufficient       = &ErrWithCode{HTTPStatusCode: http.StatusBadRequest, Message: "balance insufficient"}
	ErrFileRequired              = &ErrWithCode{HTTPStatusCode: http.StatusBadRequest, Message: "file required"}
)

type ErrWithCode struct {
	HTTPStatusCode int
	Message        string
}

func (e *ErrWithCode) Error() string {
	return e.Message
}

type ErrValidation struct {
	Message string
}

func (e *ErrValidation) Error() string {
	return e.Message
}
