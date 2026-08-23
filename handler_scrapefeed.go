package main

import (
	"context"
	"errors"
	"fmt"
	"time"
	"database/sql"
	"github.com/google/uuid"
	"github.com/Aleksy37/gator-rss/internal/database"
	"github.com/lib/pq"
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
		var nt sql.NullTime
		t, err := time.Parse(time.RFC1123Z, entry.PubDate)
		nt = sql.NullTime{Time: t, Valid: true}
		if err != nil {
			nt = sql.NullTime{Valid: false}
		}
		_, err = s.db.CreatePost(context.Background(),
		database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: entry.Title,
			Url: entry.Link,
			Description: sql.NullString{String: entry.Description, Valid: entry.Description != ""},
			PublishedAt: nt,
			FeedID: next.ID,
		})
		if err != nil {
   			var pqErr *pq.Error
    		if errors.As(err, &pqErr) && pqErr.Code ==  "23505"{
				continue
			}
			fmt.Printf("Error occurred while creating post record: %s\n%s\n\n", err, entry.Title)
			continue
					}
		fmt.Printf("Added post to database: %s\n", entry.Title)
	}
}