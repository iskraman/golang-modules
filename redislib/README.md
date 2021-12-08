# golang-modules/redislib
Golang Redis Module

### Redis Server
Redis server start
```
$ cd docker
$ docker-compose up -d
```

### func New
Make redis session
```
func New(addr string, passwd string, db int) *redis.Client
```

### func Pub
Send Publish message
```
func Pub(rdb *redis.Client, title string, data string) error

(example)
rdb := New("localhost:6379", "changeme", 0)
Pub(rdb, "Project", `{"name":"iskraman", "age":12}`)
```

### func Sub
Make subscriber session
```
func Sub(rdb *redis.Client, title string) *redis.PubSub
```

### func SubRecvMsg
Receive Subscribe message
```
func SubRecvMsg(subscriber *redis.PubSub) (string, error)

(example)
rdb := New("localhost:6379", "changeme", 0)
subs := Sub(rdb, "Project")
for {
	msg, _ := SubRecvMsg(subs)
	syslog.STD(msg)
}
```

### func Set
Redis set key, value
```
func Set(rdb *redis.Client, key string, val string) error

(example)
err := Set(rdb, "key", "value")
```

### func Get
Redis get key
```
func Get(rdb *redis.Client, key string) (string, error)

(example)
val, err := Get(rdb, "key")
```
