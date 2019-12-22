/*
@Time : 2019-12-19 17:48
@Author : lfn
@File : login
*/

package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/pioneerlfn/12306/config"
	. "github.com/pioneerlfn/12306/pkg/time"

	"github.com/spf13/viper"
)

// TODO:合理的错误处理
func (s *Session) Login() error {
	if RunMode(viper.GetString("mode")) != Test {
		SleepIfNeeded()
	}

	need, err := s.NeedCaptcha()
	if err != nil {
		return fmt.Errorf("NeeedCaptcha: %w", err)
	}
	if need {
		// 获取、识别验证码
		var answer []string
		for {
			time.Sleep(2 * time.Second)
			// 获取验证码
			rawCaptcha, err := s.GetCaptcha()
			if err != nil {
				log.Printf("获取验证码失败:%v\n", err)
				log.Println("重试...")
				continue
			}
			log.Println("下载验证码成功...")
			// 获取验证码答案
			answer, err = GetCaptchaAnswer(rawCaptcha)
			if err != nil {
				log.Printf("获取验证码答案失败:%v\n", err)
				log.Println("重试...")
				continue
			}
			log.Println("验证码解码成功...")
			// 尝试登陆
			err = s.login(answer)
			if err != nil {
				log.Printf("登录失败:%v\n", err)
				log.Println("重试...")
				continue
			}
			return nil
		}
	}
	// TODO:不需要验证码的时候？似乎没有
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

	res, err := SendRequest(s, &loginConf)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	fmt.Printf("%#v\n", string(res))

	return nil
}

// NeedCaptcha判断是否需要识别验证码
// 只有在error为nil的情况下才有意义.
func (s *Session) NeedCaptcha() (bool, error) {
	conf := config.Urls["loginConf"]
	body, err := SendRequest(s, &conf)

	cf := new(LoginConf)
	err = json.Unmarshal(body, cf)
	if err != nil {
		return false, fmt.Errorf("json.Unmarshal: %w", err)
	}

	if cf.Data.IsLoginPassCode == "N" {
		log.Println("不需要验证码...")
		return false, nil
	}
	log.Println("需要验证码:", cf.Data.IsLoginPassCode)
	return true, nil
}
