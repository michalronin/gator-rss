package main

import (
	"fmt"
	"github.com/michalronin/gator/internal/config"
	"log"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	cfg.SetUser("Ronin")
	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(updatedCfg)
}
