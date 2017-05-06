package server

import (
    "github.com/googollee/go-socket.io"
    "log"
)

func newWebsocketServer() *socketio.Server {
    server, err := socketio.NewServer(nil)
    if err != nil {
        log.Fatal(err)
    }
    server.On("connection", func(so socketio.Socket) {
        log.Println("on connection")
        so.Join("chat")
        so.On("chat message", func(msg string) {
            log.Println("emit:", so.Emit("chat message", msg))
            so.BroadcastTo("chat", "chat message", msg)
        })
        so.On("disconnection", func() {
            log.Println("on disconnect")
        })
    })
    server.On("error", func(so socketio.Socket, err error) {
        log.Println("error:", err)
    })

    return server

}
