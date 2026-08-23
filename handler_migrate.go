package main

import (
	"github.com/Aleksy37/gator-rss/internal/database"
	)	

func handlerMigrate(s *state, cmd command) error {
    return database.RunMigrations(s.rawDB)
}