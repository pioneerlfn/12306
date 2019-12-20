/*
@Time : 2019-12-19 15:56
@Author : lfn
@File : sleep
*/

package time

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func SleepIfNeeded() {
	start := viper.GetString("sleep.start")
	stop := viper.GetString("sleep.stop")

	startHour, startMin := parseTime(start)
	stopHour, stopMin := parseTime(stop)
	now := time.Now()
	year, month, day := now.Year(), now.Month(), now.Day()
	startTs := time.Date(year, month, day, startHour, startMin, 0, 0, time.Local).Unix()
	stopTs := time.Date(year, month, day, stopHour, stopMin, 0, 0, time.Local).Unix()
	ts := now.Unix()

	if ts >= startTs && ts < stopTs {
		return
	}

	// 不在12306开放时间段内，sleep.
	var needSleep time.Duration
	if ts < startTs {
		needSleep = time.Duration(startTs-ts) * time.Second
	} else { // ts >= stopTs
		needSleep = time.Duration(startTs +
			int64(time.Hour*time.Duration(startHour+24-stopHour)+
				time.Minute*time.Duration(startMin-stopMin)) - ts)
	}
	log.Println("need sleep...")
	time.Sleep(needSleep)
	log.Println("awake from sleep...")
}

func parseTime(t string) (int, int) {
	it := strings.Split(t, ":")
	if len(it) < 2 {
		panic("time config error")
	}
	hour, _ := strconv.ParseInt(it[0], 10, 64)
	min, _ := strconv.ParseInt(it[1], 10, 64)

	return int(hour), int(min)
}
