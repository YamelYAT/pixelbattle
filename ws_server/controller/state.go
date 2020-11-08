package controller

import (
	"fmt"
	"ws_server/canvas"
	"ws_server/io"
	"ws_server/io/types"
)

type State struct {
	Canvas 			*canvas.Canvas
	Players 		map[*io.Conn]bool
	Connected 		chan *io.Conn
	Disconnected	chan *io.Conn
	Message 		chan *types.Message
	Content			*types.Message
}

func NewState(c *canvas.Canvas) *State {
	return &State{
		Canvas: c,
		Players: 		make(map[*io.Conn]bool),
		Connected: 		make(chan *io.Conn),
		Disconnected: 	make(chan *io.Conn),
		Message: 		make(chan *types.Message),
		Content:		&types.Message{},
	}
}

func (s *State)Run()  {
	for{
		select {
		case c := <-s.Connected:
			s.Players[c] = true
			fmt.Println(c, "connected")
			fmt.Println(s.Players)
			break
		case c := <-s.Disconnected:
			_, ok := s.Players[c]
			if ok{
				delete(s.Players, c)
				close(c.Send)
			}
			fmt.Println(c, "disconnected")
			fmt.Println(s.Players)
			break
		case m := <-s.Message:
			s.Content = m
			s.Canvas.SetPoint(m.X, m.Y, m.Color)
			s.SendToAll()
		}
	}
}

func (s *State)SendToAll()  {
	for p := range s.Players{
		select {
			case p.Send <- s.Content.Marhsal():
				break
		default:
			close(p.Send)
			delete(s.Players, p)
		}
	}
}