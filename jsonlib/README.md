# golang-modules/jsonlib
Golang Json 모듈화

### func Encoding
Go data to JSON message encoding<br/>
```
func Encoding(v interface{}) ([]byte, error)

(example)  
type User struct {  
	string `json:"name"`  
	Age  int    `json:"age,omitempty"`  
}  

var u1 = User{Name: "iskra", Age: 10}  
enc, _ := jsonlib.Encoding(u1)  
fmt.Println(string(enc))  
```

### func EncodingIndent
Go data to JSON cute message encoding
```
func EncodingIndent(v interface{}) ([]byte, error)

(example)
enc, _ := jsonlib.EncodingIndent(u1)
```

### func EncodingStream
Go data to JSON stream encoding
```
func EncodingStream(w io.Writer, v interface{}) error

(example)
jsonlib.EncodingStream(os.Stdout, u1)

(output : stdout)
{"name":"iskra","age":10}
```

### func EncodingIndentStream
Go data to JSON cute stream encoding
```
func EncodingIndentStream(w io.Writer, v interface{}) error

(example)
wfd, _ := os.Create("out.txt")
jsonlib.EncodingIndentStream(wfd, u1)
wfd.Close()

(output : "out.txt")
{
	"name": "iskra",
	"age": 10
}
```

### func Decoding
JSON to Go data decoding
```
func Decoding(data []byte, v interface{}) error

(example)
	u2 := User{}
	jsonlib.Decoding(enc, &u2)
	fmt.Printf(%+v\n", u2)
```

### func DecodingStream
JSON to Go data stream decoding
```
func DecodingStream(r io.Reader, v interface{}) error

(example)
	u3 := User{}
	rfd, _ := os.Open("out.txt")
	jsonlib.DecodingStream(rfd, &u3)
	fmt.Printf("DecodingStream: %+v\n", u3)
	rfd.Close()
```
