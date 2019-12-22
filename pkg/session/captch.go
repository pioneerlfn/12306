/*
@Time : 2019-12-20 21:03
@Author : lfn
@File : captch
*/

package session

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Point struct {
	X int
	Y int
}

// TODO:取个更适合的名字
var Offsets [8]Point

func init() {
	Offsets[0] = Point{40, 77}
	Offsets[1] = Point{112, 77}
	Offsets[0] = Point{184, 77}
	Offsets[0] = Point{256, 77}
	Offsets[0] = Point{40, 149}
	Offsets[0] = Point{112, 149}
	Offsets[0] = Point{184, 149}
	Offsets[0] = Point{256, 149}
}

type Captcha struct {
	base64 []byte
}

func (s *Session) GetCaptcha() (*Captcha, error) {
	log.Println("下载验证码...")
	imgPath := "tkcode.png"
	//conf := config.Urls["getCaptcha"]
	resp, err := http.Get("https://kyfw.12306.cn/passport/captcha/captcha-image64?login_site=E&module=login&rand=sjrand&callback=jQuery19108016482864806321_1554298927290&_=1554298927293")
	// body, err := SendRequest(s, &conf)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}

	start := bytes.IndexByte(body, '{')
	end := bytes.IndexByte(body, '}')
	imgInfo := body[start:end+1]

	captchaRes := new(CaptchaRes)
	err = json.Unmarshal(imgInfo, captchaRes)
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}

	if captchaRes.ResultCode != "0" {
		return nil, fmt.Errorf("result error: %v", captchaRes.ResultMessage)
	}

	file, err := os.Create(imgPath)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %w", err)
	}
	defer file.Close()

	w, err := file.Write(captchaRes.Image)
	if err != nil {
		return nil, fmt.Errorf("write file: %w", err)
	}
	if w != len(captchaRes.Image) {
		return nil, fmt.Errorf("paritially write")
	}

	return &Captcha{body}, nil
}

func GetCaptchaAnswer(captcha *Captcha) ([]string, error) {
	// needle是一个数组，比如[]string{"篮球", "网球拍"}或者[]string{"中国结"}
	needle, err := DecodeNeedle(captcha)
	if err != nil {
		return nil, fmt.Errorf("DecodeNeedle: %w", err)
	}
	// hay是一个数组，共8个元素
	hay, err := DecodeHayStack(captcha)
	if err != nil {
		return nil, fmt.Errorf("DecodeHayStack: %w", err)
	}

	var hits []int
	for i, item := range hay {
		for _, nl := range needle {
			if nl == item {
				hits = append(hits, i)
			}
		}
	}

	answer := convertToCoordinate(hits)
	return answer, nil
}

func DecodeHayStack(c *Captcha) ([]string, error) {
	return nil, nil
}

func DecodeNeedle(c *Captcha) ([]string, error) {
	return nil, nil
}

func convertToCoordinate(items []int) []string {
	var res []string
	for it := range items {
		point := Offsets[it]
		res = append(res, strconv.Itoa(point.X), strconv.Itoa(point.Y))
	}
	return res
}

func DecodeCaptcha(captcha *Captcha) (string, error) {
	return "", nil
}

func (s *Session) VerifyCaptcha() error {
	return nil
}
