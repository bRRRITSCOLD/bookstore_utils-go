package errors_utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type APIError interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type apiError struct {
	message string        `json:"message"`
	status  int           `json:"status"`
	error   string        `json:"error"`
	causes  []interface{} `json:"causes"`
}

func (e apiError) Message() string {
	return e.message
}

func (e apiError) Status() int {
	return e.status
}

func (e apiError) Error() string {
	return fmt.Sprintf(
		"message: %s - status: %d - error: %s - causes: %v",
		e.message,
		e.status,
		e.error,
		e.causes,
	)
}

func (e apiError) Causes() []interface{} {
	return e.causes
}

func NewAPIErrorFromBytes(bytes []byte) (APIError, error) {
	var apiErr apiError
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewAPIError(message string, status int, err string, causes []interface{}) APIError {
	return apiError{
		message: message,
		status:  status,
		error:   err,
		causes:  causes,
	}
}

func NewBadRequestAPIError(message string, err error) APIError {
	var causes []interface{}
	if err != nil {
		causes = append(causes, err.Error())
	}
	return NewAPIError(
		message,
		http.StatusBadRequest,
		"bad_request",
		causes,
	)
}

func NewUnauthorizedAPIError(message string, err error) APIError {
	var causes []interface{}
	if err != nil {
		causes = append(causes, err.Error())
	}
	return NewAPIError(
		message,
		http.StatusUnauthorized,
		"unauthorized",
		causes,
	)
}

func NewNotFoundAPIError(message string, err error) APIError {
	var causes []interface{}
	if err != nil {
		causes = append(causes, err.Error())
	}
	return NewAPIError(
		message,
		http.StatusNotFound,
		"not_found",
		causes,
	)
}

func NewInternalServerAPIError(message string, err error) APIError {
	var causes []interface{}
	if err != nil {
		causes = append(causes, err.Error())
	}
	return NewAPIError(
		message,
		http.StatusInternalServerError,
		"internal_server_error",
		causes,
	)
}
