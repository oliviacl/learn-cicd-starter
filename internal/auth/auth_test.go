package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	testCases := []struct {
		name          string
		authHeader    string
		expectedKey   string
		expectedError bool
	}{
		{
			name:          "No auth header",
			authHeader:    "",
			expectedKey:   "",
			expectedError: true,
		},
		{
			name:          "Valid auth header",
			authHeader:    "ApiKey 1234a",
			expectedKey:   "1234a",
			expectedError: false,
		},
		{
			name:          "Malformed auth header",
			authHeader:    "apiKey 1234a",
			expectedKey:   "",
			expectedError: true,
		},
		{
			name:          "Missing token",
			authHeader:    "ApiKey",
			expectedKey:   "",
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			headers := make(http.Header)
			if tc.authHeader != "" {
				headers.Set("Authorization", tc.authHeader)
			}

			key, err := GetAPIKey(headers)
			if !tc.expectedError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}

			if tc.expectedError && err == nil {
				t.Errorf("expected an error, got nil")
			}

			if key != tc.expectedKey {
				t.Errorf("expected %s; got %s instead", tc.expectedKey, key)
			}
		})
	}
}
