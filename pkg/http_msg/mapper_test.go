package http_msg

import (
	"github.com/khalil-farashiani/fusion-wallet-hub/pkg/richerror"
	"net/http"
	"testing"
)

func TestMapKindToHTTPStatusCode(t *testing.T) {
	tests := []struct {
		name     string
		kind     richerror.Kind
		expected int
	}{
		{
			name:     "Testing Invalid case",
			kind:     richerror.Invalid,
			expected: http.StatusUnprocessableEntity,
		},
		{
			name:     "Testing NotFound case",
			kind:     richerror.NotFound,
			expected: http.StatusNotFound,
		},
		{
			name:     "Testing Forbidden case",
			kind:     richerror.Forbidden,
			expected: http.StatusForbidden,
		},
		{
			name:     "Testing Unexpected case",
			kind:     richerror.Unexpected,
			expected: http.StatusInternalServerError,
		},
		{
			name:     "Testing default case",
			kind:     8,
			expected: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mapKindToHTTPStatusCode(tc.kind)
			if result != tc.expected {
				t.Errorf("Expected %d, but got %d", tc.expected, result)
			}
		})
	}
}
