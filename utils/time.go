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
