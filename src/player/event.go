package player

type Event struct {
	Player *Player
	Msg    string
}

func NewEvent(player *Player, msg string) *Event {
	event := Event{}
	event.Player = player
	event.Msg = msg
	return &event
}
