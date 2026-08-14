package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aleksy37/gator-rss/internal/database"
	"github.com/google/uuid"
)

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
		return fmt.Errorf("usage: register <username>")
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("too many arguments provided")
	}
	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(), 
		Name: cmd.args[0],
	})
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

func handlerListUsers(s *state, _ command) error {
	userList, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting list of users: %s", err)
	}
	for _, user := range userList {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}
	return nil
}
