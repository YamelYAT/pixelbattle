package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"ws_server/canvas"
	"ws_server/controller"
	"ws_server/io"
)

func main(){
	canv := canvas.NewCanvas()
	state := controller.NewState(canv)

	go state.Run()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET"{
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		w.Header().Set("content-type", "text/html; charset=utf-8")
		w.Header().Set("content-length", "636401")

		if _, err := w.Write(canv.Points); err != nil{
			log.Println(err)
		}
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
		if _, ok := err.(websocket.HandshakeError); ok {
			http.Error(w, "Not a websocket handshake", 400)
			return
		} else if err != nil {
			return
		}

		conn := &io.Conn{Ws:ws, Send: make(chan []byte)}

		state.Connected <- conn

		go conn.Receiver(state.Disconnected, state.Message)
		conn.Sender()

	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
