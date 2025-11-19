package apperrors

import "net/http"

type AppError interface {
	Error() string
	Code() string
	Message() string
	Status() int
	Unwrap() error
}

type appError struct {
	code    string
	message string
	status  int
	err     error
}

func (e *appError) Error() string   { return e.message }
func (e *appError) Code() string    { return e.code }
func (e *appError) Message() string { return e.message }
func (e *appError) Status() int     { return e.status }
func (e *appError) Unwrap() error   { return e.err }

func New(code string, err error, status int) *appError {
	return &appError{code: code, message: err.Error(), status: status, err: err}
}

func BadRequest(code string, err error) *appError {
	return New(code, err, http.StatusBadRequest)
}

func Internal(code string, err error) *appError {
	return New(code, err, http.StatusInternalServerError)
}
