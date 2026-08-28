package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		authHeader string
		wantKey    string
		wantErr    error
	}{
		{
			name:       "valid API key",
			authHeader: "ApiKey my-secret-key",
			wantKey:    "my-secret-key",
			wantErr:    nil,
		},
		{
			name:       "missing authorization header",
			authHeader: "",
			wantKey:    "",
			wantErr:    ErrNoAuthHeaderIncluded,
		},
		{
			name:       "missing API key",
			authHeader: "ApiKey",
			wantKey:    "",
			wantErr:    errors.New("malformed authorization header"),
		},
		{
			name:       "wrong authorization scheme",
			authHeader: "Bearer my-secret-key",
			wantKey:    "",
			wantErr:    errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}

			if tt.authHeader != "" {
				headers.Set("Authorization", tt.authHeader)
			}

			gotKey, err := GetAPIKey(headers)

			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() key = %q, want %q", gotKey, tt.wantKey)
			}

			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("GetAPIKey() unexpected error = %v", err)
				}
				return
			}

			if err == nil {
				t.Fatalf("GetAPIKey() expected error %q, got nil", tt.wantErr)
			}

			if err.Error() != tt.wantErr.Error() {
				t.Errorf("GetAPIKey() error = %q, want %q", err, tt.wantErr)
			}
		})
	}
}
