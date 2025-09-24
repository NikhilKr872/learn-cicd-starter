package auth

import (
	"net/http"
	"testing"
)

func TestSplit(t *testing.T) {
	headers := make(http.Header)
	apiKey := "sk-12i9319"
	headers.Add("Authorization", "Api "+apiKey)
	resApiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Error encountered %v", err)
	} else if resApiKey != apiKey {
		t.Fatalf("Api key mismatch")
	}

}
