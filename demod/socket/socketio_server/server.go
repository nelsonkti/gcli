package socketio_server

import (
	"demod/lib/logger"
	"fmt"
	"github.com/googollee/go-socket.io"
	engineio "github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	_ "github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
	"log"
	"net/http"
	_ "net/http/pprof"
)

const NameSpace = "/"

func Start(port int) {

	wt := websocket.Default
	pt := polling.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			wt,
			pt,
		},
	})

	server.OnConnect(NameSpace, func(s socketio.Conn) error {
		return nil
	})

	server.OnError(NameSpace, func(s socketio.Conn, e error) {
	})

	//断开连接
	server.OnDisconnect(NameSpace, func(s socketio.Conn, reason string) {

	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Println(fmt.Sprintf("Serving at localhost:%d...", port))
	//http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func StopDevice() {

	logger.Sugar.Info("stop device")
}
