package server

import (
	"fmt"
	"net"
	"player"
	"strconv"
)

type Server struct {
	Port    int
	Players []*player.Player
}

func eventHandler(plyr *player.Player, bus chan *player.Event) {
	var event *player.Event
	for !plyr.Connected {
		fmt.Println("Waiting for player to connect")
	}
	for plyr.Connected {
		event = <-bus
		fmt.Println(event)
		plyr.Send("received event.")
	}
}

func NewServer(port int) *Server {
	server := Server{}
	server.Port = port
	return &server
}

func (this Server) Start() {
	fmt.Println("Server start up initiated, opening socket on: " + strconv.Itoa(this.Port))
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(this.Port))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Server listening on " + strconv.Itoa(this.Port))

	for {
		bus := make(chan *player.Event)
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		plyr := player.NewPlayer("---UNNAMED---", len(this.Players)+1, bus)
		this.Players = append(this.Players, plyr)

		go plyr.Connect(conn)
		go eventHandler(plyr, bus)
	}
}
