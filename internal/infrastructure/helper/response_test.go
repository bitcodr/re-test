package helper_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitcodr/re-test/internal/infrastructure/helper"
)

func TestResponseError(t *testing.T) {
	t.Run("ResponseError - Successful Response", func(t *testing.T) {
		// Create a test HTTP ResponseRecorder
		w := httptest.NewRecorder()

		// Simulate an error
		err := fmt.Errorf(" %s", "test Error")

		// Call ResponseError with the error
		helper.ResponseError(w, "Test Error Message", err)

		// Check the response status code
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}

		// Check the response body for the message
		expectedBody := "Test Error Message"
		actualBody := w.Body.String()
		if actualBody != expectedBody {
			t.Errorf("Expected body '%s', got '%s'", expectedBody, actualBody)
		}

		// Check if the error message is logged,
		// You may need to capture the log output for this test if it's being logged
	})

	t.Run("ResponseError - No Error", func(t *testing.T) {
		// Create a test HTTP ResponseRecorder
		w := httptest.NewRecorder()

		// Call ResponseError without an error
		helper.ResponseError(w, "Test Message", nil)

		// Check the response status code
		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}

		// Check the response body for the message
		expectedBody := "Test Message"
		actualBody := w.Body.String()
		if actualBody != expectedBody {
			t.Errorf("Expected body '%s', got '%s'", expectedBody, actualBody)
		}

		// Check if the error message is not logged,
		// You may need to capture the log output for this test if it's being logged
	})
}

func TestResponseSuccess(t *testing.T) {
	t.Run("ResponseSuccess - Successful Response", func(t *testing.T) {
		// Create a test HTTP ResponseRecorder
		w := httptest.NewRecorder()

		// Create a sample message
		message := "message"

		// Call ResponseSuccess with the message
		helper.ResponseSuccess[string](w, message)

		// Check the response status code
		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}

		// Check the response body for the JSON message
		expectedJSON, _ := json.Marshal(message)
		actualBody := w.Body.String()
		if actualBody != string(expectedJSON) {
			t.Errorf("Expected JSON body '%s', got '%s'", string(expectedJSON), actualBody)
		}
	})
}
