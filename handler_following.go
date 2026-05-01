package main

import (
	"context"
	"fmt"
	"github.com/aleksy37/gator-rss/internal/database"
)


func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("too many arguments provided")
	}
	feeds , err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %w", err)
	}
	fmt.Printf("%s is following these feeds:\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("%s\n", feed.FeedName)
	}
	return nil
}