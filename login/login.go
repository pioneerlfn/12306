/*
@Time : 2019-12-19 17:48
@Author : lfn
@File : login
*/

package login

import (
	"errors"

	"github.com/pioneerlfn/12306/time"
	"github.com/spf13/viper"
)

func Login(auth bool) error {
	if auth {
		Auth()
	}
	time.SleepIfNeeded()
	user := viper.GetString("auth.USER")
	pw := viper.GetString("auth.PWD")
	if user == "" || pw == "" {
		return errors.New("auth info needed")
	}

	return nil
}

func Auth() {

}
