package utils

import (
	"fmt"
	"time"
)

func GetTimeStr() string {
	timestamp := time.Unix(time.Now().Unix(), 0)
	date := fmt.Sprintf("%v", timestamp)
	return date
}

func GetUnixTimeSec() int64 {
	now := time.Now()
	secs := now.Unix()
	return secs
}

func GetUnixTimeMil() int64 {
	now := time.Now()
	nanos := now.UnixNano()
	mils := nanos / 1000000
	return mils
}

func GetUnixTimeNano() int64 {
	now := time.Now()
	nanos := now.UnixNano()
	return nanos
}
