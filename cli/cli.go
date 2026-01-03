package cli

import (
	"errors"
)

type commandHandler func(*State, Command) error

type commands struct {
	Handlers map[string]commandHandler
}

func (c *commands) Run(s *State, cmd Command) error {
	if fn, ok := c.Handlers[cmd.Name]; ok {
		err := fn(s, cmd)
		return err
	}
	return errors.New("handler " + cmd.Name + " not found")
}

func (c *commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}

func NewCommands() *commands {
	return &commands{
		Handlers: make(map[string]commandHandler),
	}
}
