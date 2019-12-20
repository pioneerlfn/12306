package main

import (
	"fmt"
	"log"
	"time"

	"github.com/pioneerlfn/12306/login"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/pioneerlfn/12306/config"
)

var (
	cfg = pflag.StringP("config", "c", "", "ticket assistant config file path.")
)


func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	err := login.Login()

	// cookies, err := login.GetCookies()
	if err != nil {
		log.Fatal(err)
	}
	/*for _, cookie := range cookies {
		if cookie.Name == "RAIL_EXPIRATION" || cookie.Name == "RAIL_DEVICEID" {
			log.Println(cookie.Name, ":", cookie.Value)
		}
	}

	*/

	// 测试热更新
	for {
		fmt.Println(viper.GetString("runmode"))
		time.Sleep(4 * time.Second)
	}
}
