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
func Pub(client *redis.Client, title string, data string) error

(example)
client := New("localhost:6379", "changeme", 0)
Pub(client, "Project", `{"name":"iskraman", "age":12}`)
```

### func Sub
Make subscriber session
```
func Sub(client *redis.Client, title string) *redis.PubSub
```

### func SubRecvMsg
Receive Subscribe message
```
func SubRecvMsg(subscriber *redis.PubSub) (string, error)

(example)
client := New("localhost:6379", "changeme", 0)
subs := Sub(client, "Project")
for {
	msg, _ := SubRecvMsg(subs)
	syslog.STD(msg)
}
```

