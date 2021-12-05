package main

import (
	"fmt"

	"github.com/iskraman/golang-modules/fslib"
)

func main() {
	// slice test
	/*
		items := []int{1, 2, 3, 4, 5, 6}
		fmt.Println(slice.SliceExists(items, 5))   // returns true
		fmt.Println(slice.SliceExists(items, "5")) // returns false
	*/

	// log test
	/*
		syslog.DBG("%s %d", "log", 12)
		syslog.STD("%s %d", "log", 34)
		syslog.WAR("%s %d", "log", 56)
		syslog.ERR("%s %d", "log", 78)
	*/

	// jsonlib test
	/*
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age,omitempty"`
		}
		var u1 = User{Name: "iskra", Age: 10}
		enc, _ := jsonlib.Encoding(u1)
		fmt.Println("Encoding:", string(enc))

		u2 := User{}
		jsonlib.Decoding(enc, &u2)
		fmt.Printf("Decoding: %+v\n", u2)

		jsonlib.EncodingStream(os.Stdout, u1)

		wfd, _ := os.Create("out.txt")
		jsonlib.EncodingIndentStream(wfd, u1)
		wfd.Close()

		u3 := User{}
		rfd, _ := os.Open("out.txt")
		jsonlib.DecodingStream(rfd, &u3)
		fmt.Printf("DecodingStream: %+v\n", u3)
		rfd.Close()
	*/

	// fs test
	data, _ := fslib.FileReader("./fslib/readme.txt")
	fmt.Println(string(data))

	fslib.FileWriter("./test.txt", data, 0644)
}
