package player

import (
	"bufio"
	"fmt"
	"input"
	"net"
	"position"
	"strconv"
)

type Player struct {
	Id        int
	Name      string
	conn      net.Conn
	Connected bool
	out       *bufio.Writer
	in        *bufio.Reader
	Pos       *position.Position
	bus       chan *Event
}

func NewPlayer(name string, id int, bus chan *Event) *Player {
	player := Player{}
	player.Name = name
	player.Id = id
	player.Pos = position.NewPosition(0, 0, 0)
	player.bus = bus
	return &player
}

func (this Player) Send(msg string) {
	this.out.WriteString(msg + "\n")
	this.out.Flush()
}

func (this Player) Connect(connection net.Conn) {
	this.conn = connection
	this.Connected = true
	fmt.Println("player connected")
	this.out = bufio.NewWriter(this.conn)
	this.in = bufio.NewReader(this.conn)
	this.init()
	go this.listen()
}

func (this Player) processCommand(msg string) {
	command := input.NewCommand(msg)

	switch command.Cmd {
	case "setpos":
		origin := position.Position{}
		if len(command.Args) != 3 {
			this.Send("incorrect arguments, we expected setpos <x> <y> <z>")
		} else {
			this.Pos.X, _ = strconv.ParseFloat(command.Args[0], 64)
			this.Pos.Y, _ = strconv.ParseFloat(command.Args[1], 64)
			this.Pos.Z, _ = strconv.ParseFloat(command.Args[2], 64)
			this.Send("Distance from origin: " + strconv.FormatFloat(this.Pos.Distance(origin), 'f', -1, 64))
		}

		break
	case "say":
		fmt.Println("sending event to bus")
		event := NewEvent(&this, "message")
		this.bus <- event
	}
}

func (this Player) init() {
	this.Send("Welcome to go-server, " + this.Name)
}

func (this Player) listen() {
	for this.Connected {
		message, err := this.in.ReadString('\n')
		if err != nil {
			this.Connected = false
			fmt.Println("Player: " + strconv.Itoa(this.Id) + " has disconnected!")
		} else {
			this.processCommand(message)
		}
	}
}
