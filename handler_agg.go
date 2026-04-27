package main

import (
	"fmt"
	"context"
)

func handlerAgg(s *state, cmd command) error {
	res, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error fetching rss feed: %v", err)
	}
	fmt.Printf("succesfully fetched rss feed from %s : %s", "https://www.wagslane.dev/index.xml", res)
	return nil
}