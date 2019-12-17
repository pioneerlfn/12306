package main

import (
	"fmt"
	"time"

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

	// 测试热更新
	for {
		fmt.Println(viper.GetString("runmode"))
		time.Sleep(4*time.Second)
	}
}
