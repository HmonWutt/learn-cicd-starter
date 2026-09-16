package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKEY(t *testing.T) {
	API_KEY := "12345hjkl"
	want := API_KEY
	header := http.Header{}
	header.Add("Authorization", "ApiKey "+API_KEY)

	api, err := GetAPIKey(header)
	if want != api || err != nil {
		t.Errorf(`%s does not match %s, error: %v`, api, want, err)
	}
}
