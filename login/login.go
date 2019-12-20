/*
@Time : 2019-12-19 17:48
@Author : lfn
@File : login
*/

package login

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	// "github.com/pioneerlfn/12306/time"
	"github.com/spf13/viper"
)

func Login() error {
	// time.SleepIfNeeded()

	user := viper.GetString("auth.user")
	pw := viper.GetString("auth.pwd")
	appId := viper.GetString("auth.appid")
	if user == "" || pw == "" {
		return errors.New("auth info needed")
	}
	authInfo := &AuthInfo{
		Username: user,
		Password: pw,
		Appid: appId,
	}
	authInfo.Answer = GetCaptch()

	err := login(authInfo)
	if err != nil {
		return err
	}
	return nil
}

type AuthInfo struct {
	Username string
	Password string
	Appid string
	Answer string
}

type Conf struct {
	Url string
	Method string
	Referer string
	Host string
	ContentType string
	Retry int
	ReTime int
	STime float64
	IsLogger bool
	Isjson bool
}


func login(authInfo *AuthInfo) error {
	data := make(url.Values)
	data.Set("username", authInfo.Username)
	data.Set("password", authInfo.Password)
	data.Set("appid", authInfo.Appid)
	data.Set("answer", authInfo.Answer)

	// 按 urlencoded 组织数据
	body := data.Encode()
	// 创建请求并设置内容类型

	request, _ := http.NewRequest(
		http.MethodPost,
		// TODO: 从Urls获取
		viper.GetString("host") + "/passport/web/login",
		bytes.NewReader([]byte(body)),
	)

	cookie, err := GetCookies()
	if err != nil {
		return fmt.Errorf("login:%w", err)
	}
	cks , err := json.Marshal(cookie)
	if err != nil {
		return fmt.Errorf("login:%w", err)
	}

	request.Header.Set("cookie",string(cks))
	request.Header.Set("content-type","application/x-www-form-urlencoded")
	request.Header.Set("Referer", "https://kyfw.12306.cn/otn/resources/login.html")

	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return errors.New("login failed")
	}

	return nil
}




