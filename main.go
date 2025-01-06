package main

import (
	"fmt"
	"log"
	"os"

	"github.com/michalronin/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	var s state
	s.cfg = &cfg
	cmds := commands{
		commands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("error: not enough arguments")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}
