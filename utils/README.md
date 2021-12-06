# golang-modules/utils
Golang Utils Module

### func SliceExists
슬라이스내에 특정값이 존재하는지 확인
```
func SliceExists(slice interface{}, item interface{}) bool 

(example)
items := []int{1, 2, 3, 4, 5, 6}
fmt.Println(SliceExists(items, 5))   // returns true
fmt.Println(SliceExists(items, "5")) // returns false
```
