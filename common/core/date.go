package core

import (
	"time"
	"strconv"
)

//时间戳转日期
func TimeToDate(timestamp string) string {
	ts, _ := strconv.ParseInt(timestamp, 10, 64)
	tm := time.Unix(ts, 0)
	return tm.Format("2006-01-02 15:04:05")
}