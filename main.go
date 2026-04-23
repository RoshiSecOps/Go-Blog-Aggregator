package main

import (
	"fmt"
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
		fmt.Println("Not enough arguments... Exiting!")
		os.Exit(1)
	}
}
