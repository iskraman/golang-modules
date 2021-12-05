package slice

import (
	"reflect"

	"github.com/iskraman/golang-modules/syslog"
)

func SliceExists(slice interface{}, item interface{}) bool {
	s := reflect.ValueOf(slice)

	if s.Kind() != reflect.Slice {
		syslog.WAR("SliceExists() given a non-slice type")
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if s.Index(i).Interface() == item {
			return true
		}
	}

	return false
}
