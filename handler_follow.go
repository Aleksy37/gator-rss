package main

import (
	"context"
	"fmt"
	"time"
	"github.com/aleksy37/gator-rss/internal/database"
	"github.com/google/uuid"
)


func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: follow <url>")
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("too many arguments provided")
	}
	feed , err := s.db.GetFeedByURL(context.Background(), cmd.args[0])
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %w", err)
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
	fmt.Printf("%s subscribed to a new feed: %s\n", follow.UserName, follow.FeedName)
	return nil
}