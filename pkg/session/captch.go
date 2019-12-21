/*
@Time : 2019-12-20 21:03
@Author : lfn
@File : captch
*/

package session

import (
	"fmt"
	"strconv"
)

type Point struct {
	X int
	Y int
}

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
	base64 string
}

func (s *Session) GetCaptcha() (*Captcha, error) {

	// TODO:
	return &Captcha{}, nil
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