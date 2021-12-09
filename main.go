package main

import (
	"github.com/iskraman/golang-modules/redislib"
	"github.com/iskraman/golang-modules/utils/syslog"
)

func main() {
	/*
		// slice test
		items := []int{1, 2, 3, 4, 5, 6}
		syslog.DBG("%v", slice.SliceExists(items, 5))   // returns true
		syslog.DBG("%v", slice.SliceExists(items, "5")) // returns false
	*/

	// syslog test
	syslog.SetLogLevel(syslog.DBG_LEVEL)
	syslog.DBG("%s %d", "system ready", 12)
	syslog.STD("%s %d", "system ready", 34)
	syslog.WAR("%s %d", "system ready", 56)
	syslog.ERR("%s %d", "system ready", 78)

	/*
		// jsonlib test
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age,omitempty"`
		}
		var u1 = User{Name: "iskra", Age: 10}
		enc, _ := jsonlib.Encoding(u1)
		syslog.DBG("Encoding: %s", string(enc))

		u2 := User{}
		jsonlib.Decoding(enc, &u2)
		syslog.DBG("Decoding: %+v\n", u2)

		var reading map[string]interface{}
		jsonlib.DecodingMap(enc, &reading)
		syslog.DBG("Map Decoding: %+v\n", reading)
		syslog.DBG("%v, %v", reading["name"], reading["age"])

		jsonlib.EncodingStream(os.Stdout, u1)

		wfd, _ := os.Create("out.txt")
		jsonlib.EncodingIndentStream(wfd, u1)
		wfd.Close()

		u3 := User{}
		rfd, _ := os.Open("out.txt")
		jsonlib.DecodingStream(rfd, &u3)
		syslog.DBG("DecodingStream: %+v", u3)
		rfd.Close()
	*/

	/*
		// fs test
		data, _ := fslib.FileReader("./fslib/readme.txt")
		syslog.DBG("%s", string(data))

		fslib.FileWriter("./test.txt", data, 0644)
	*/

	/*
		// websocket test
		recvCallBack := func(conn *websocket.Conn) {
			for {
				msg, err := websock.Reader(conn)
				if err != nil {
					syslog.DBGLN(err)
					return
				}

				// TODO : Echo test
				syslog.DBGLN("Recv:", msg)
				websock.Sender(conn, msg)
			}
		}
		websock.ServerRun("localhost", 8080, recvCallBack)
	*/

	/*
		// Publisher
		rdb := redis.New("localhost:6379", "changeme", 0)
		redis.Pub(rdb, "Project", `{"name":"iskraman", "age":12}`)
	*/

	/*
		// Subscriber
		rdb := redislib.New("localhost:6379", "changeme", 0)
		subs := redislib.Sub(rdb, "Project")
		for {
			msg, _ := redislib.SubRecvMsg(subs)
			syslog.STD(msg)
		}
	*/

	// Redis Set/Get cache
	rdb := redislib.New("localhost:6379", "changeme", 0)
	redislib.Set(rdb, "key1", "my value")
	val, err := redislib.Get(rdb, "key1") // Exist key
	syslog.STDLN(val, err)

	val2, err := redislib.Get(rdb, "key2") // Not exist key
	if err != nil {
		syslog.WARLN(err)
	} else {
		syslog.STDLN(val2)
	}

	// Redis HSet/HGet cache
	server := map[string]interface{}{"cpu": 25.0, "mem": 10.5, "hdd": "40"}
	redislib.HSet(rdb, "media-server-0", server)
	redislib.HSet(rdb, "media-server-1", map[string]interface{}{"cpu": 10, "mem": 15, "hdd": "20"})

	cpu, _ := redislib.HGet(rdb, "media-server-0", "cpu")
	mem, _ := redislib.HGet(rdb, "media-server-0", "mem")
	hdd, _ := redislib.HGet(rdb, "media-server-0", "hdd")
	syslog.STD("cpu:%v, mem:%v, hdd:%v", cpu, mem, hdd)

	// Redis HGetAll
	all, _ := redislib.HGetAll(rdb, "media-server-1")
	for k, v := range all {
		syslog.STD("%v(%T):%v(%T)", k, k, v, v)
	}

	// key update
	redislib.HSet(rdb, "media-server-1", map[string]interface{}{"hdd": "100"})
	all, _ = redislib.HGetAll(rdb, "media-server-1")
	for k, v := range all {
		syslog.STD("Update: %v(%T):%v(%T)", k, k, v, v)
	}

	// Redis HDel (default field)
	rslt := redislib.HDel(rdb, "media-server-0", "mem")
	mem, _ = redislib.HGet(rdb, "media-server-0", "mem")
	syslog.STD("HDel:%v, cpu:%v, mem:%v, hdd:%v", rslt, cpu, mem, hdd)

	// Redis Del (default key)
	rslt = redislib.Del(rdb, "media-server-1")
	all, _ = redislib.HGetAll(rdb, "media-server-1")
	syslog.STD("Del:%v", rslt)
}
