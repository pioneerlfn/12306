package cookie

import (
	"fmt"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	railDeviceId   = "RAIL_DEVICEID"
	railExpiration = "RAIL_EXPIRATION"
)

func getCookies(index string) ([]selenium.Cookie, error) {
	// TODO: 从配置文件读取
	const (
		// These paths will be different on your system.
		//seleniumPath = "vendor/selenium-server-standalone-3.4.jar"
		//geckoDriverPath = "vendor/geckodriver-v0.18.0-linux64"
		seleniumPath = "chromedriver"
		port         = 9999
	)

	//selenium.SetDebug(true)
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	service, err := selenium.NewChromeDriverService(seleniumPath, port, []selenium.ServiceOption{}...)
	if err != nil {
		fmt.Println("NewChromeDriverService")
		panic(err) // panic is used only as an example and is not otherwise recommended.
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

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return nil, err
	}

	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get(index); err != nil {
		return nil, err
	}
	cookies, _ := wd.GetCookies()
	// deviceId, _ := wd.GetCookie(railDeviceId)
	//expiration, _ := wd.GetCookie(railExpiration)

	return cookies, nil
}
