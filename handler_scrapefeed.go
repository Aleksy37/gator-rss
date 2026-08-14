package main

import (
	"fmt"
	"context"
)

func scrapeFeeds(s *state) {
	next, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Printf("error fetching rss feed details: %v", err)
		return 
	}
	err = s.db.MarkFeedFetched(context.Background(), next.ID)
	if err != nil {
		fmt.Printf("Error updating timestamp: %v", err)
		return
	}
	res, err := fetchFeed(context.Background(), next.Url)
	if err != nil {
		fmt.Printf("Error fetching feed by url: %v", err)
		return 
	}
	fmt.Printf("Fetched RSS Feed from %s\n", next.Url)
	fmt.Printf("Found %d entries:\n", len(res.Channel.Item))
	for _, entry := range res.Channel.Item {
		fmt.Printf("%s\n", entry.Title)
	}
}