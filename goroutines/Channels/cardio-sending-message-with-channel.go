package main

import (
	"fmt"

)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgCh chan Message
}

func (s *Server) ServeAndListen() {
	for {
		select {
		case msg := <-s.msgCh:
			fmt.Printf("Received a new message from %s with message %s", msg.From, msg.Payload)
			// default:
		}
	}
}

func sendMessageToServer(msgCh chan Message, msg Message) {
	newMsg := Message{
		From:    msg.From,
		Payload: msg.Payload,
	}

	msgCh <- newMsg
}

func main() {
	s := &Server{
		msgCh: make(chan Message),
	}

	go s.ServeAndListen()

	sendMessageToServer(s.msgCh, Message{
		From:    "Joe",
		Payload: "How's it going?",
	})
}
