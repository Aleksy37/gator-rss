package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aleksy37/gator-rss/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	if len(cmd.args) > 2 {
		return fmt.Errorf("too many arguments provided")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	feed, err := s.db.CreateFeed(
		context.Background(),
		database.CreateFeedParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name: cmd.args[0],
			Url: cmd.args[1],
			UserID: user.ID,
		})
	if err != nil {
		return fmt.Errorf("there was an issue creating the feed: %w", err)
	}
	fmt.Printf("%s subscribed to a new feed: %s", user.Name, feed)
	return nil
}