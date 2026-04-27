package main

import (
	"fmt"

	"github.com/aleksy37/gator-rss/internal/config"
	"github.com/aleksy37/gator-rss/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	cmds map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	command, ok := c.cmds[cmd.name]
	if !ok {
		return fmt.Errorf("%s is not a registered command", cmd.name)
	}
	err := command(s, cmd)
	if err != nil {
		return fmt.Errorf("error running command: %s", err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmds[name] = f
}
