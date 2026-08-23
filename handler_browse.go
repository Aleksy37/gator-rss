package main

import (
	"context"
	"fmt"
	"github.com/Aleksy37/gator-rss/internal/database"
	"strconv"
)


func handlerBrowse(s *state, cmd command, user database.User) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("browse only takes one optional arg for amount of results to return")
	}
	i64, err := strconv.ParseInt(cmd.args[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid argument for browse: %w", err)
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID:	user.ID, 
		Limit: int32(i64),
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