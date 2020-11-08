package io

import (
	"fmt"
	"github.com/gorilla/websocket"
	"time"
	"ws_server/io/types"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 1024 * 1024
)

type Conn struct {
	Ws	 *websocket.Conn
	Send chan []byte
}

func (c *Conn) Receiver(unreg chan *Conn, broadcast chan *types.Message)  {
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error {
		c.Ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for{
		_, msg, err := c.Ws.ReadMessage()
		msgReader := types.Message{}
		if err != nil{
			break
		}

		fmt.Println(c, msg)
		broadcast <- msgReader.UnMarshal(msg)
	}

	unreg <- c
	c.Ws.Close()
}

func (c *Conn) Sender()  {
	ticker := time.NewTicker(pingPeriod)

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}

	ticker.Stop()
	c.Ws.Close()
}

func (c *Conn) write(mt int, message []byte) error{
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(mt, message)
}

