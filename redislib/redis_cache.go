package redislib

import (
	"encoding/json"

	"github.com/go-redis/redis/v8"
	"github.com/iskraman/golang-modules/syslog"
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
		return err
	}
	return err
}

func HSet(rdb *redis.Client, key string, field map[string]interface{}) error {
	err := rdb.HSet(ctx, key, field).Err()
	if err != nil {
		syslog.WARLN(err)
		return err
	}
	return err
}

func SetJson(rdb *redis.Client, key string, val interface{}) error {
	jsonVal, err := json.Marshal(val)
	if err != nil {
		syslog.WARLN(err)
		return err
	}

	err = rdb.Set(ctx, key, jsonVal, 0).Err()
	if err != nil {
		syslog.WARLN(err)
		return err
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
		return val, err
	}

	return val, err
}

func HGet(rdb *redis.Client, key string, field string) (string, error) {
	val, err := rdb.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return val, e
	} else if err != nil {
		syslog.WARLN(err)
		return val, err
	}
	return val, err
}

func GetJson(rdb *redis.Client, key string, data interface{}) error {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return e
	} else if err != nil {
		syslog.WARLN(err)
		return err
	}

	err = json.Unmarshal([]byte(val), data)
	if err != nil {
		syslog.WARLN(err)
		return err
	}

	return err
}

func Keys(rdb *redis.Client, key string) []string {
	var cursor uint64
	var keys []string
	var err error
	for {
		keys, cursor, err = rdb.Scan(ctx, cursor, key, 0).Result()
		if err != nil {
			syslog.WARLN(err)
		}

		/*
			for _, key := range keys {
				fmt.Println("key", key)
			}
		*/

		if cursor == 0 { // no more keys
			break
		}
	}

	return keys
}

func HGetAll(rdb *redis.Client, key string) (map[string]string, error) {
	field, err := rdb.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return field, e
	} else if err != nil {
		syslog.WARLN(err)
		return field, err
	}
	return field, err
}

func Del(rdb *redis.Client, key string) bool {
	rslt, _ := rdb.Del(ctx, key).Result()
	if rslt == 1 {
		return true
	} else {
		return false
	}
}

func HDel(rdb *redis.Client, key string, field string) bool {
	rslt, _ := rdb.HDel(ctx, key, field).Result()
	if rslt == 1 {
		return true
	} else {
		return false
	}
}
