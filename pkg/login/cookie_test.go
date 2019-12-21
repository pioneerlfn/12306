package login

import (
	"log"
	"testing"
)

func TestGetCookies(t *testing.T) {
	cookies, err := GetCookies()
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range cookies {
		if cookie.Name == "RAIL_EXPIRATION" || cookie.Name == "RAIL_DEVICEID" {
			t.Log(cookie.Value)
		}
	}
}
