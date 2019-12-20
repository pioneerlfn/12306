package login

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	railDeviceId   = "RAIL_DEVICEID"
	railExpiration = "RAIL_EXPIRATION"
)

type RunMode string

// 似乎有点别扭，找机会改进
const (
	Debug   RunMode = "DEBUG"
	Release RunMode = "RELEASE"
	Test    RunMode = "TEST"
)

func GetCookies(index string) ([]selenium.Cookie, error) {

	selenium.SetDebug(viper.Get("runmode") == Debug)
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	path := viper.GetString("selenium.path")
	port := viper.GetInt("selenium.port")
	service, err := selenium.NewChromeDriverService(path, port, []selenium.ServiceOption{}...)
	if err != nil {
		return nil, fmt.Errorf("NewChromeDriverService%w", err)
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}

	// 不开启chrome浏览器界面
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"--headless",
		},
	}
	caps.AddChrome(chromeCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", viper.GetInt("selenium.port")))
	if err != nil {
		return nil, fmt.Errorf("NewRemote%w", err)
	}
	defer wd.Quit()

	if err := wd.Get(index); err != nil {
		return nil, fmt.Errorf("wd.Get%w", err)
	}
	// 给wd一点时间，获取cookies.
	time.Sleep(time.Second * 3)
	cookies, _ := wd.GetCookies()

	return cookies, nil
}
