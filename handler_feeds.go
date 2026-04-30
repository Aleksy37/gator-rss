package main

import (
	"context"
	"fmt"
)


func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		return fmt.Errorf("feeds does not take any arguments")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("there was an issue retreiving the feeds: %w", err)
	}
	
	if len(feeds) == 0 {
		fmt.Println("no feeds found.")
		return nil
	}

	for _, feed := range feeds {
		fmt.Printf("Name: %s | URL: %s | User: %s\n", feed.Name, feed.Url, feed.UserName)
	}
	return nil
}