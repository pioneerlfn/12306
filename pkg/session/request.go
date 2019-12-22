/*
@Time : 2019-12-21 18:17
@Author : lfn
@File : request
*/

package session

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/pioneerlfn/12306/config"
)

const userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36"

func SendRequest(s *Session, c *config.Conf) (body []byte, err error) {
	request, _ := http.NewRequest(
		c.Method,
		"https://" + c.Host+c.Url,
		bytes.NewReader([]byte(c.Body)),
	)

	if len(c.ContentType) > 0 {
		request.Header.Set("Content-type", c.ContentType)
	}
	request.Header.Set("Cookie", s.Cookies)
	request.Header.Set("Referer", c.Referer)
	request.Header.Set("Host", c.Host)
	request.Header.Set("User-Agent", userAgent)

	fmt.Printf("%#v\n", *request.URL)
	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}

	defer rsp.Body.Close()

	fmt.Println("status code:", rsp.StatusCode)
	fmt.Println("response header:", rsp.Header)

	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d", rsp.StatusCode)
	}

	body, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	return body, nil
}


func SendRequest2(s *Session, c *config.Conf) (res *http.Response, err error) {
	request, err := http.NewRequest(
		c.Method,
		"http://" + c.Host+c.Url,
		bytes.NewReader([]byte(c.Body)),
	)
	if err != nil {
		panic(err)
	}
	request.Header.Set("Host", c.Host)
	request.Header.Set("User-Agent", userAgent)


	log.Println(request.Host, "\n",request.Body, "\n",request.Header)
	log.Println(*request)



	rsp, err := s.Client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("do: %w", err)
	}
	defer rsp.Body.Close()

	io.Copy(os.Stdout, rsp.Body)
	if rsp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: %d", rsp.StatusCode)
	}

	return res, nil
}
