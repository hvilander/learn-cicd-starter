package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		wantErr    bool
		apiKey     string
		wantAPIKey string
	}{
		{
			name:       "emtpy auth token",
			apiKey:     "",
			wantErr:    true,
			wantAPIKey: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "myURL", nil)
			if err != nil {
				t.Error("test setup failure")
			}
			req.Header.Set("Authorization", tt.apiKey)
			apiKeyOut, err := GetAPIKey(req.Header)
			if (err != nil) != tt.wantErr {
				t.Errorf("somthing about got error unexpectedly")
			}
			if apiKeyOut != tt.wantAPIKey {
				t.Errorf("API keys don't match. \nWanted: %s \nGot: %s", tt.wantAPIKey, apiKeyOut)
			}

		})
	}

}
