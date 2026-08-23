package main

import (
	"context"
	"fmt"
	"github.com/Aleksy37/gator-rss/internal/database"
)


func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	if len(cmd.args) > 1 {
		return fmt.Errorf("too many arguments provided")
	}
	err := s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		Url: cmd.args[0], UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error fetching feed by url: %w", err)
	}
	fmt.Printf("%s is no longer following: %s\n", user.Name, cmd.args[0])
	return nil
}