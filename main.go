package main

import (
	"log"

	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	s := &state{cfg: &cfg}
}
