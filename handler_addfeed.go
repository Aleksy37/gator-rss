package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aleksy37/gator-rss/internal/database"
	"github.com/google/uuid"
)


func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	if len(cmd.args) > 2 {
		return fmt.Errorf("too many arguments provided")
	}
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
	follow, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID: user.ID,
			FeedID: feed.ID,
		})
	if err != nil {
		return fmt.Errorf("there was an issue following the feed: %w", err)
	}
	fmt.Printf("%s subscribed to a new feed: %s\n", user.Name, follow.FeedName)
	return nil
}