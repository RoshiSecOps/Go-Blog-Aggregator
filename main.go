package main

import (
	"log"
	"os"

	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := &state{cfg: &cfg}
	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
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
