package main

import (
	"fmt"
	"log"

	"github.com/aleksy37/gator-rss/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read Config: %+v\n", cfg)

	err = cfg.SetUser("Aleksy")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read updated config: %+v\n", cfg)

}