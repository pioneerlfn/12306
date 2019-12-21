/*
@Time : 2019-12-20 22:01
@Author : lfn
@File : login_test
*/

package login

import (
	"testing"

	"github.com/pioneerlfn/12306/config"
)

func TestLogin(t *testing.T) {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	err := Login()
	t.Log(err)
}
