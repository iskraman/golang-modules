# golang-modules/websock
Golang Websocket 모듈화

### func Reader
read websocket message<br/>
```
func Reader(conn *websocket.Conn) (string, error)
```

### func Sender
send websocket message
```
func Sender(conn *websocket.Conn, msg string)
```

### func ServerRun
Websocket server run
```
func ServerRun(addr string, port int, cb func(conn *websocket.Conn))
```

### Web Server Start
Chrome Web browser (localhost:8000) -> F12 (Console debug)
```
$ cd golang-modules/websock
$ python -m SimpleHTTPServer
Serving HTTP on 0.0.0.0 port 8000 ...
```
