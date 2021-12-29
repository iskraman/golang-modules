package redislib

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/iskraman/golang-modules/syslog"
)

var ctx = context.Background()

func New(addr string, passwd string, db int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	})
	return redisClient
}

func Pub(rdb *redis.Client, title string, data string) error {
	err := rdb.Publish(ctx, title, data).Err()
	if err != nil {
		syslog.WAR("Publish failed: %v", err)
	}
	return err
}

func Sub(rdb *redis.Client, title string) *redis.PubSub {
	subscriber := rdb.Subscribe(ctx, title)
	return subscriber
}

func SubRecvMsg(subscriber *redis.PubSub) (string, error) {
	msg, err := subscriber.ReceiveMessage(ctx)
	if err != nil {
		syslog.WARLN("Gateway subscriber failed:", err)
		return "", err
	}

	return msg.Payload, err
}

/*
// Publisher
func main() {
	rdb := New("localhost:6379", "changeme", 0)
	Pub(rdb, "Project", `{"name":"iskraman", "age":12}`)
}
*/

/*
// Subscriber
func main() {
	rdb := New("localhost:6379", "changeme", 0)
	subs := Sub(rdb, "Project")
	for {
		msg, _ := SubRecvMsg(subs)
		syslog.STD(msg)
	}
}

*/
