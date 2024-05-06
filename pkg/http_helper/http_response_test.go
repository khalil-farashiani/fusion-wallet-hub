package http_helper

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJSON(t *testing.T) {
	recorder := httptest.NewRecorder()

	// Test data
	data := map[string]interface{}{
		"key": "value",
	}

	// Expected JSON
	expectedJSON, _ := json.Marshal(data)

	// Call the function
	JSON(recorder, http.StatusOK, data)

	// Check the status code
	assert.Equal(t, http.StatusOK, recorder.Code, "Status code should be 200")

	// Check the content type
	assert.Equal(t, jsonMIME, recorder.Header().Get(headerContentType), "Content type should be application/json;charset=UTF-8")

	// Check the body
	assert.JSONEq(t, string(expectedJSON), recorder.Body.String(), "Body should match expected JSON")
}

func TestJSONErr(t *testing.T) {
	recorder := httptest.NewRecorder()

	// Test error message
	errorMessage := "error occurred"

	// Expected JSON
	expectedJSON, _ := json.Marshal(simpleErr{Err: errorMessage})

	// Call the function
	JSONErr(recorder, http.StatusInternalServerError, errorMessage)

	// Check the status code
	assert.Equal(t, http.StatusInternalServerError, recorder.Code, "Status code should be 500")

	// Check the content type
	assert.Equal(t, jsonMIME, recorder.Header().Get(headerContentType), "Content type should be application/json;charset=UTF-8")

	// Check the body
	assert.JSONEq(t, string(expectedJSON), recorder.Body.String(), "Body should match expected JSON")
}
