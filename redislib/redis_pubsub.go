package redislib

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/iskraman/golang-modules/utils/syslog"
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

func Pub(client *redis.Client, title string, data string) error {
	err := client.Publish(ctx, title, data).Err()
	if err != nil {
		syslog.WAR("Publish failed: %v", err)
	}
	return err
}

func Sub(client *redis.Client, title string) *redis.PubSub {
	subscriber := client.Subscribe(ctx, title)
	return subscriber
}

func SubRecvMsg(subscriber *redis.PubSub) (string, error) {
	msg, err := subscriber.ReceiveMessage(ctx)
	return msg.Payload, err
}

/*
// Publisher
func main() {
	client := New("localhost:6379", "changeme", 0)
	Pub(client, "Project", `{"name":"iskraman", "age":12}`)
}
*/

/*
// Subscriber
func main() {
	client := New("localhost:6379", "changeme", 0)
	subs := Sub(client, "Project")
	for {
		msg, _ := SubRecvMsg(subs)
		syslog.STD(msg)
	}
}

*/
