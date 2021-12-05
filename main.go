package main

import "github.com/iskraman/golang-modules/syslog"

func main() {
	// slice test
	/*
		items := []int{1, 2, 3, 4, 5, 6}
		syslog.DBG("%v", slice.SliceExists(items, 5))   // returns true
		syslog.DBG("%v", slice.SliceExists(items, "5")) // returns false
	*/

	// syslog test
	syslog.SetLogLevel(syslog.STD_LEVEL)
	syslog.DBG("%s %d", "system ready", 12)
	syslog.STD("%s %d", "system ready", 34)
	syslog.WAR("%s %d", "system ready", 56)
	syslog.ERR("%s %d", "system ready", 78)

	// jsonlib test
	/*
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

	// fs test
	/*
		data, _ := fslib.FileReader("./fslib/readme.txt")
		syslog.DBG("%s", string(data))

		fslib.FileWriter("./test.txt", data, 0644)
	*/
}
