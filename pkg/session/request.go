/*
@Time : 2019-12-21 18:17
@Author : lfn
@File : request
*/

package session

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const userAgent  = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"

func SendRequest(s *Session, c *Conf) (res *http.Response, err error) {
	request, _ := http.NewRequest(
		c.Method,
		c.Host+c.Url,
		bytes.NewReader([]byte(c.Body)),
	)
	request.Header.Set("Cookie", s.Cookies)
	request.Header.Set("Content-type", c.ContentType)
	request.Header.Set("Referer", c.Referer)
	request.Header.Set("Host", c.Host)
	request.Header.Set("User-Agent", userAgent)

	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("login: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed")
	}

	return res, nil
}
