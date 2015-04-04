package helpers

import (
	"time"
)

func Unix_milisec_time_now() int64 {
	return time.Now().Unix() * 1000
}

func UnixTimeAddMinFromNow(min int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(min)).Unix()
}
