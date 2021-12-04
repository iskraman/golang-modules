package main

import (
	"fmt"

	"github.com/iskraman/golang-lib/modules/slice"
	"github.com/iskraman/golang-lib/modules/syslog"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice.SliceExists(items, 5))   // returns true
	fmt.Println(slice.SliceExists(items, "5")) // returns false

	syslog.DBG("%s %d", "log", 12)
	syslog.ST("%s %d", "log", 34)
	syslog.WARN("%s %d", "log", 56)
	syslog.ERR("%s %d", "log", 78)
}
