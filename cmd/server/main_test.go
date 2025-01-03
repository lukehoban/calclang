package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gotest.tools/assert"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		body           CalculationRequest
		expectedStatus int
		expectedBody   CalculationResponse
	}{
		{
			name:           "valid addition",
			method:         http.MethodPost,
			body:           CalculationRequest{Expression: "ADD 1 2"},
			expectedStatus: http.StatusOK,
			expectedBody:   CalculationResponse{Result: 3},
		},
		{
			name:           "valid subtraction",
			method:         http.MethodPost,
			body:           CalculationRequest{Expression: "SUB 5 3"},
			expectedStatus: http.StatusOK,
			expectedBody:   CalculationResponse{Result: 2},
		},
		{
			name:           "invalid expression",
			method:         http.MethodPost,
			body:           CalculationRequest{Expression: "MUL 1 2"},
			expectedStatus: http.StatusOK,
			expectedBody:   CalculationResponse{Error: "Invalid expression"},
		},
		{
			name:           "invalid method",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var bodyBytes []byte
			if tc.method == http.MethodPost {
				var err error
				bodyBytes, err = json.Marshal(tc.body)
				assert.NilError(t, err)
			}

			req := httptest.NewRequest(tc.method, "/calculate", bytes.NewBuffer(bodyBytes))
			w := httptest.NewRecorder()

			calculateHandler(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)

			if tc.expectedStatus == http.StatusOK {
				var response CalculationResponse
				err := json.NewDecoder(w.Body).Decode(&response)
				assert.NilError(t, err)
				assert.Equal(t, tc.expectedBody, response)
			}
		})
	}
}