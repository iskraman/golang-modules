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

/*
https://pkg.go.dev/github.com/go-redis/redis/v8#section-readme

- HSet("myhash", "key1", "value1", "key2", "value2")
- HSet("myhash", []string{"key1", "value1", "key2", "value2"})
- HSet("myhash", map[string]interface{}{"key1": "value1", "key2": "value2"})
*/

func Set(rdb *redis.Client, key string, val string) error {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		syslog.WARLN(err)
	}
	return err
}

func HSet(rdb *redis.Client, key string, val map[string]interface{}) error {
	err := rdb.HSet(ctx, key, val).Err()
	if err != nil {
		syslog.WARLN(err)
	}
	return err
}

func Get(rdb *redis.Client, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return val, e
	} else if err != nil {
		syslog.WARLN(err)
	}
	return val, err
}

func HGet(rdb *redis.Client, key string, val string) (string, error) {
	val, err := rdb.HGet(ctx, key, val).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return val, e
	} else if err != nil {
		syslog.WARLN(err)
	}
	return val, err
}

func HGetAll(rdb *redis.Client, key string) (map[string]string, error) {
	val, err := rdb.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return val, e
	} else if err != nil {
		syslog.WARLN(err)
	}
	return val, err
}
