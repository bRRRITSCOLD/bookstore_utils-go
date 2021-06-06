package errors_utils

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBadRequestAPIError(t *testing.T) {
	err := NewBadRequestAPIError("hello world", errors.New("error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "hello world", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, err.Causes[0], "error")
}

func TestNewNotFoundAPIError(t *testing.T) {
	err := NewNotFoundAPIError("hello world", errors.New("error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "hello world", err.Message)
	assert.EqualValues(t, "not_found", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, err.Causes[0], "error")
}

func TestNewInternalServerAPIError(t *testing.T) {
	err := NewInternalServerAPIError("hello world", errors.New("error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "hello world", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "error", err.Causes[0])
}
