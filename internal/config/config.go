package config

import (
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func ReadConfig() (Config, error) {
	var config Config
	dir, err := os.UserHomeDir()
	if err != nil {
		return config, err
	}
	fullpath := dir + "/.gatorconfig.json"
	fmt.Println(fullpath)
	return config, fmt.Errorf("nice")
}
