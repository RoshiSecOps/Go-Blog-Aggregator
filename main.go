package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/config"
	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/database"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	s := &state{cfg: &cfg, db: dbQueries}
	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerFetchFeed)
	cmds.register("addfeed", handlerAddFeed)
	args := os.Args
	if len(args) < 2 {
		log.Fatal("Not enough arguments... Exiting!")
	}
	cmd := command{name: args[1], arguments: args[2:]}
	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
