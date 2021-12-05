# golang-modules/jsonlib
Golang Json 모듈화

### func Encoding
```
func Encoding(v interface{}) ([]byte, error)
```
Go data to JSON message encoding
(example)
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
	}

	var u1 = User{Name: "iskra", Age: 10}
	enc, _ := jsonlib.Encoding(u1)
	fmt.Println(string(enc))

### func EncodingIndent
```
func EncodingIndent(v interface{}) ([]byte, error)
```
Go data to JSON cute message encoding
(example)
	enc, _ := jsonlib.EncodingIndent(u1)

### func EncodingStream
```
func EncodingStream(w io.Writer, v interface{}) error
```
Go data to JSON stream encoding
(example)
	jsonlib.EncodingStream(os.Stdout, u1)

(output : stdout)
	{"name":"iskra","age":10}

### func EncodingIndentStream
```
func EncodingIndentStream(w io.Writer, v interface{}) error
```
Go data to JSON cute stream encoding
(example)
	wfd, _ := os.Create("out.txt")
	jsonlib.EncodingIndentStream(wfd, u1)
	wfd.Close()

(output : "out.txt")
	{
  		"name": "iskra",
  		"age": 10
	}

### func Decoding
```
func Decoding(data []byte, v interface{}) error
```
JSON to Go data decoding
(example)
	u2 := User{}
	jsonlib.Decoding(enc, &u2)
	fmt.Printf(%+v\n", u2)

### func DecodingStream
```
func DecodingStream(r io.Reader, v interface{}) error
```
JSON to Go data stream decoding
(example)
	u3 := User{}
	rfd, _ := os.Open("out.txt")
	jsonlib.DecodingStream(rfd, &u3)
	fmt.Printf("DecodingStream: %+v\n", u3)
	rfd.Close()
