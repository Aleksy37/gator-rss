package main

import (
	"github.com/aleksy37/gator-rss/internal/config"
	"fmt"
)

type state struct {
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: login <username>")
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("too many arguments provided")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error logging in: %s", err)
	}
	fmt.Printf("%s is now logged in\n", s.cfg.CurrentUserName)
	return nil
}
