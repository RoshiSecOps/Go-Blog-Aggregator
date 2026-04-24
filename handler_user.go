package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/RoshiSecOps/Go-Blog-Aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No arguments provided...")
	}
	userName := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		os.Exit(1)
	}
	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}
	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("No arguments provided...")
	}
	userName := cmd.arguments[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err == nil {
		os.Exit(1)
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	userId := uuid.New()
	user, err := s.db.CreateUser(context.Background(),
		database.CreateUserParams{ID: userId, CreatedAt: createdAt, UpdatedAt: updatedAt, Name: userName})
	if err != nil {
		return err
	}
	s.cfg.SetUser(userName)
	fmt.Println("User: ", user.Name, " was created successfully")
	fmt.Println("Created at: ", user.CreatedAt)
	fmt.Println("Updated at: ", user.UpdatedAt)
	fmt.Println("UUID: ", user.ID)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetDb(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
