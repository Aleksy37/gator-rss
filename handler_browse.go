package main

import (
	"context"
	"fmt"
	"github.com/aleksy37/gator-rss/internal/database"
)


func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("feeds only takes one optional arg for amount of results to return")
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:	user.ID, 
		Limit: 10,
	})
	if err != nil {
		return fmt.Errorf("there was an issue retreiving the posts: %w", err)
	}
	
	if len(posts) == 0 {
		fmt.Println("no post found.")
		return nil
	}

	for _, post := range posts {
		fmt.Println()
		fmt.Printf("%s\n", post.Title)
		fmt.Printf("%s", post.Url)
		fmt.Println()
	}
	return nil
}