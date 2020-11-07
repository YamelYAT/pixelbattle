package types

import (
	"encoding/json"
	"log"
)

type Message struct {
	X,
	Y,
	Color int
}

func (m *Message) Marhsal() []byte{
	res,err := json.Marshal(m)
	if err != nil{
		log.Print(err)
	}
	return res
}

func (m *Message)UnMarshal(b []byte) *Message {
	err := json.Unmarshal(b, m)
	if err != nil{
		log.Print(err)
	}
	return m
}