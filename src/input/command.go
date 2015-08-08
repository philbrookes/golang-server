package input

import "strings"

type Command struct {
	Cmd  string
	Args []string
}

func NewCommand(input string) *Command {
	command := Command{}
	bits := strings.Split(strings.TrimSpace(input), " ")
	for i, val := range bits {
		if i == 0 {
			command.Cmd = strings.ToLower(val)
		} else {
			command.Args = append(command.Args, val)
		}
	}
	return &command
}
