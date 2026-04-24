package main

import (
	"log"
	"os"

	"github.com/aleksy37/gator-rss/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	s := state{cfg: &cfg}
	c := commands{cmds : make(map[string]func(*state, command) error)}
	c.register("login", handlerLogin)
	userCommand := os.Args[1]
	userArgs := os.Args[2:]
	if len(userArgs) < 1 {
		log.Fatalf("too few arguments provided, please provide at least 1")
	}
	err = c.run(&s, command{userCommand, userArgs})
	if err != nil {
		log.Fatalf("error executing command: %v error: %v", userCommand, err)
	}

}