package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username required")
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Println("Username has been set.")
	return nil
}
