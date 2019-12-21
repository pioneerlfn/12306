/*
@Time : 2019-12-19 17:48
@Author : lfn
@File : login
*/

package login

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strings"

	"net/http"
	"net/url"


	// "github.com/pioneerlfn/12306/time"
	"github.com/spf13/viper"
)

type GoLogin struct {

}





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
	log.Println("len(cookies):",len(cookie))
	if err != nil {
		return fmt.Errorf("login:%w", err)
	}
	var cookieVal []string
	cookieVal = append(cookieVal, "BIGipServerpool_statistics=635503114.44582.0000")
	for _, ck := range cookie {
		fmt.Println(ck.Name, ck.Value, ck.Expiry, ck.Domain)
		if ck.Name == railExpiration {
			cookieVal = append(cookieVal, strings.Join([]string{ck.Name,ck.Value}, "="))
		}
	}
	cookieVal = append(cookieVal, "RAIL_DEVICEID=FMKAiW_-bbqIc1q9UdMLwI0kBLV7hk1AIctjmhOclmG0Ymn3MXdXocVC92xn-5UgE-Xc123nHAo4iAcasT844vKj9gf2pVHwmK3fAzm2tUaqy0jtl1e7K3j2Hld-xd5lvC8dCXwkKR3ujaumpqJ_2iTtKDCIvMX1")

	request.Header.Set("Cookie",strings.Join(cookieVal, "; "))
	request.Header.Set("Content-type","application/x-www-form-urlencoded")
	request.Header.Set("Referer", "https://kyfw.12306.cn/otn/resources/login.html")
	request.Header.Set("User-Agent", "Paw/3.1.10 (Macintosh; OS X/10.12.6)")
	request.Header.Set("Host","kyfw.12306.cn")

	fmt.Println("req:",request)
	return nil
	rsp, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		return errors.New("login failed")
	}

	for k, v := range rsp.Header {
		fmt.Println(k,":", v)
	}
	return nil
}




