package cookie

import (
	"log"
	"testing"
)

func TestGetCookies(t *testing.T) {
	index := "https://www.12306.cn/index/index.html"
	cookies, err := GetCookies(index)
	if err != nil {
		log.Fatal(err)
	}
	for _, cookie := range cookies {
		if cookie.Name == "RAIL_EXPIRATION" || cookie.Name == "RAIL_DEVICEID" {
			t.Log(cookie.Value)
		}
	}
}
