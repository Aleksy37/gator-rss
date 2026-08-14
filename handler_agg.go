package main

import (
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("please provide a duration string for scraping frequency eg. 5m15s")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing duration string: %v", err)
	}
	fmt.Printf("Starting scrapper with a cooldown of %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
for ; ; <-ticker.C {
	scrapeFeeds(s)
}
}