/*
@Time : 2019-12-19 17:48
@Author : lfn
@File : login
*/

package session

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pioneerlfn/12306/config"
	"github.com/pioneerlfn/12306/pkg/time"
	// "github.com/pioneerlfn/12306/time"
	"github.com/spf13/viper"
)

type GoLogin struct {
}


type Conf struct {
	Url         string
	Method      string
	Referer     string
	Host        string
	Body        string
	ContentType string
	Retry       int
	ReTime      int
	STime       float64
	IsLogger    bool
	Isjson      bool
}

// TODO:合理的错误处理
func (s *Session) LogIn() error {
	time.SleepIfNeeded()

	// 获取、识别验证码
	var answer []string
	for {
		if NeedCaptcha(nil) { // todo:fix
			// 获取验证码
			rawCaptcha, err := s.GetCaptcha()
			if err != nil {
				panic(err)
			}
			// 获取验证码答案
			answer, err = GetCaptchaAnswer(rawCaptcha)
			if err != nil {
				panic(err)
			}
		}
		// 尝试登陆
		err := s.login(answer)
		if err == nil {
			break
		}
		// TODO:根据错误类型不同选择合适地处理方式
	}
	return nil
}



func (s *Session) login(answer []string) error {
	data := make(url.Values)
	data.Set("answer", strings.Join(answer, ","))

	user := viper.GetString("auth.user")
	pw := viper.GetString("auth.pwd")
	appId := viper.GetString("auth.appid")
	if user == "" || pw == "" || appId == "" {
		return errors.New("auth info needed")
	}
	data.Set("username", user)
	data.Set("password", pw)
	data.Set("appid", appId)

	// 按 urlencoded 组织数据
	body := data.Encode()

	// 创建请求并设置内容类型
	loginConf := config.Urls["login"]
	loginConf.Body = body

	rsp, err := SendRequest(s, &loginConf)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return errors.New("login failed")
	}

	return nil
}

// 是否需要识别验证码
func NeedCaptcha(c *Conf) bool {

	return true
}
