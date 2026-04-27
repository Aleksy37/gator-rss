package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aleksy37/gator-rss/internal/config"
	"github.com/aleksy37/gator-rss/internal/database"
	"github.com/google/uuid"
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: login <username>")
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("too many arguments provided")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error looking up user in the databse: %s", err)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error logging in: %s", err)
	}
	fmt.Printf("%s is now logged in\n", s.cfg.CurrentUserName)
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: login <username>")
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("too many arguments provided")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: cmd.args[0]})
	if err != nil {
		return fmt.Errorf("error registering user in the databse: %s", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error loggin in: %s", err)
	}
	fmt.Printf("user created successfully: %v\n", user)
	return nil
}
