package main

import (
	"fmt"

	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/config"
	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
type command struct {
	name      string
	arguments []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("Command does not exist")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
