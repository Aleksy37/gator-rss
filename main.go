package main

import (
	"database/sql"
	"log"
	"os"
	"github.com/aleksy37/gator-rss/internal/config"
	"github.com/aleksy37/gator-rss/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error opening db connection: %s", err)
	}
	dbQueries := database.New(db)

	s := state{db : dbQueries, cfg: &cfg}
	c := commands{cmds : make(map[string]func(*state, command) error)}

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("users", handlerListUsers)


	if len(os.Args) <2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	userCommand := os.Args[1]
	userArgs := os.Args[2:]
	
	err = c.run(&s, command{userCommand, userArgs})
	if err != nil {
		log.Fatalf("error executing command: %v error: %v", userCommand, err)
	}

}