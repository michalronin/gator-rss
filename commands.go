package main

import (
	"errors"
)

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	cmdToRun, ok := c.commands[cmd.name]
	if ok {
		if err := cmdToRun(s, cmd); err != nil {
			return err
		}
	} else {
		return errors.New("command not found")
	}
	return nil
}
