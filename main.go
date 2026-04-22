package main

import (
	"fmt"

	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/config"
)

func main() {
	var gatorfig config.Config
	var err error
	gatorfig, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	gatorfig.SetUser("Lane")
}
