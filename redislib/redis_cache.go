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

func Set(rdb interface{}, key string, val string) error {
	var err error
	switch rdb.(type) {
	case *redis.Client:
		err = rdb.(*redis.Client).Set(ctx, key, val, 0).Err()
		if err != nil {
			syslog.WARLN(err)
			return err
		}
	case *redis.ClusterClient:
		err = rdb.(*redis.ClusterClient).Set(ctx, key, val, 0).Err()
		if err != nil {
			syslog.WARLN(err)
			return err
		}
	}
	return err
}

func HSet(rdb interface{}, key string, field map[string]interface{}) error {
	var err error
	switch rdb.(type) {
	case *redis.Client:
		err := rdb.(*redis.Client).HSet(ctx, key, field).Err()
		if err != nil {
			syslog.WARLN(err)
			return err
		}
	case *redis.ClusterClient:
		err := rdb.(*redis.ClusterClient).HSet(ctx, key, field).Err()
		if err != nil {
			syslog.WARLN(err)
			return err
		}
	}
	return err
}

func SetJson(rdb interface{}, key string, val interface{}) error {
	var err error
	jsonVal, err := json.Marshal(val)
	if err != nil {
		syslog.WARLN(err)
		return err
	}

	switch rdb.(type) {
	case *redis.Client:
		err = rdb.(*redis.Client).Set(ctx, key, jsonVal, 0).Err()
		if err != nil {
			syslog.WARLN(err)
			return err
		}
	case *redis.ClusterClient:
		err = rdb.(*redis.ClusterClient).Set(ctx, key, jsonVal, 0).Err()
		if err != nil {
			syslog.WARLN(err)
			return err
		}
	}
	return err
}

func Get(rdb interface{}, key string) (string, error) {
	var val string
	var err error

	switch rdb.(type) {
	case *redis.Client:
		val, err = rdb.(*redis.Client).Get(ctx, key).Result()
	case *redis.ClusterClient:
		val, err = rdb.(*redis.ClusterClient).Get(ctx, key).Result()
	}

	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return val, e
	} else if err != nil {
		syslog.WARLN(err)
		return val, err
	}

	return val, err
}

func HGet(rdb interface{}, key string, field string) (string, error) {
	var val string
	var err error

	switch rdb.(type) {
	case *redis.Client:
		val, err = rdb.(*redis.Client).HGet(ctx, key, field).Result()
	case *redis.ClusterClient:
		val, err = rdb.(*redis.ClusterClient).HGet(ctx, key, field).Result()
	}

	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return val, e
	} else if err != nil {
		syslog.WARLN(err)
		return val, err
	}
	return val, err
}

func GetJson(rdb interface{}, key string, data interface{}) error {
	var val string
	var err error

	switch rdb.(type) {
	case *redis.Client:
		val, err = rdb.(*redis.Client).Get(ctx, key).Result()
	case *redis.ClusterClient:
		val, err = rdb.(*redis.ClusterClient).Get(ctx, key).Result()
	}

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

func Keys(rdb interface{}, key string) []string {
	var cursor uint64
	var keys []string
	var err error
	for {
		switch rdb.(type) {
		case *redis.Client:
			keys, cursor, err = rdb.(*redis.Client).Scan(ctx, cursor, key, 0).Result()
			if err != nil {
				syslog.WARLN(err)
			}

			if cursor == 0 { // no more keys
				//break
				return keys
			}
		case *redis.ClusterClient:
			keys, cursor, err = rdb.(*redis.Client).Scan(ctx, cursor, key, 0).Result()
			if err != nil {
				syslog.WARLN(err)
			}

			if cursor == 0 { // no more keys
				//break
				return keys
			}
		}
	}
	//return keys
}

func HGetAll(rdb interface{}, key string) (map[string]string, error) {
	var field map[string]string
	var err error

	switch rdb.(type) {
	case *redis.Client:
		field, err = rdb.(*redis.Client).HGetAll(ctx, key).Result()
	case *redis.ClusterClient:
		field, err = rdb.(*redis.ClusterClient).HGetAll(ctx, key).Result()
	}

	if err == redis.Nil {
		e := errorMsg{msg: "Does not exist key!"}
		return field, e
	} else if err != nil {
		syslog.WARLN(err)
		return field, err
	}
	return field, err
}

func Del(rdb interface{}, key string) bool {
	var rslt int64

	switch rdb.(type) {
	case *redis.Client:
		rslt, _ = rdb.(*redis.Client).Del(ctx, key).Result()
	case *redis.ClusterClient:
		rslt, _ = rdb.(*redis.ClusterClient).Del(ctx, key).Result()
	}

	if rslt == 1 {
		return true
	} else {
		return false
	}
}

func HDel(rdb interface{}, key string, field string) bool {
	var rslt int64

	switch rdb.(type) {
	case *redis.Client:
		rslt, _ = rdb.(*redis.Client).HDel(ctx, key, field).Result()
	case *redis.ClusterClient:
		rslt, _ = rdb.(*redis.ClusterClient).HDel(ctx, key, field).Result()
	}

	if rslt == 1 {
		return true
	} else {
		return false
	}
}

/*
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
*/
