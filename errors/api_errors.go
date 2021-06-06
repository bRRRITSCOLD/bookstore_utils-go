package errors_utils

import "net/http"

type APIError struct {
	Message string        `json:"message"`
	Status  int           `json:"status"`
	Error   string        `json:"error"`
	Causes  []interface{} `json:"causes"`
}

func NewBadRequestAPIError(message string, err error) *APIError {
	return &APIError{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
		Causes:  []interface{}{err.Error()},
	}
}

func NewNotFoundAPIError(message string, err error) *APIError {
	return &APIError{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
		Causes:  []interface{}{err.Error()},
	}
}

func NewInternalServerAPIError(message string, err error) *APIError {
	apiErr := &APIError{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
	if err != nil {
		apiErr.Causes = append(apiErr.Causes, []interface{}{err.Error()})
	}
	return apiErr
}
