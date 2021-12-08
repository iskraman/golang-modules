package redislib

import (
	"github.com/go-redis/redis/v8"
	"github.com/iskraman/golang-modules/utils/syslog"
)

type errorMsg struct {
	msg string
}

func (e errorMsg) Error() string {
	return e.msg
}

func Set(rdb *redis.Client, key string, val string) error {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		syslog.WARLN(err)
	}
	return err
}

func Get(rdb *redis.Client, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return "", e
	} else if err != nil {
		syslog.WARLN(err)
	}
	return val, err
}
