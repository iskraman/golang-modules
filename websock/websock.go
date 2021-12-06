package websock

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/iskraman/golang-modules/utils/syslog"
)

/* messageType: https://www.rfc-editor.org/rfc/rfc6455.html#section-11.8
    |Opcode  | Meaning                             | Reference |
   -+--------+-------------------------------------+-----------|
    | 0      | Continuation Frame                  | RFC 6455  |
   -+--------+-------------------------------------+-----------|
    | 1      | Text Frame                          | RFC 6455  |
   -+--------+-------------------------------------+-----------|
    | 2      | Binary Frame                        | RFC 6455  |
   -+--------+-------------------------------------+-----------|
    | 8      | Connection Close Frame              | RFC 6455  |
   -+--------+-------------------------------------+-----------|
    | 9      | Ping Frame                          | RFC 6455  |
   -+--------+-------------------------------------+-----------|
    | 10     | Pong Frame                          | RFC 6455  |
   -+--------+-------------------------------------+-----------|
*/

/* CloseError: https://pkg.go.dev/github.com/gorilla/websocket#section-readme
CloseNormalClosure           = 1000
CloseGoingAway               = 1001
CloseProtocolError           = 1002
CloseUnsupportedData         = 1003
CloseNoStatusReceived        = 1005
CloseAbnormalClosure         = 1006
CloseInvalidFramePayloadData = 1007
ClosePolicyViolation         = 1008
CloseMessageTooBig           = 1009
CloseMandatoryExtension      = 1010
CloseInternalServerErr       = 1011
CloseServiceRestart          = 1012
CloseTryAgainLater           = 1013
CloseTLSHandshake            = 1015
*/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Reader(conn *websocket.Conn) (string, error) {
	_, p, err := conn.ReadMessage()
	if err != nil {
		return "", err
	}
	return string(p), err
}

func Sender(conn *websocket.Conn, msg string) {
	messageType := 1
	if err := conn.WriteMessage(messageType, []byte(msg)); err != nil {
		syslog.ERRLN(err)
		return
	}
}

func ServerRun(addr string, port int, cb func(conn *websocket.Conn)) {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			syslog.ERRLN(err)
		}

		syslog.STD("Client successfully connected...")

		go cb(ws)
	})

	server := fmt.Sprintf("%s:%d", addr, port)
	syslog.STD("WebSocket start... %s", server)

	err := http.ListenAndServe(server, nil)
	if err != nil {
		syslog.ERRLN(err)
		panic(err)
	}
}
