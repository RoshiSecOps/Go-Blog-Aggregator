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

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if s.cfg.CurrentUserName == user.Name {
			fmt.Println("*", user.Name, "(current)")
		} else {
			fmt.Println("*", user.Name)
		}
	}
	return nil
}

func handlerFetchFeed(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
	return nil
}

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("Not enough arguments, provide Feed name and URL")
	}
	feedName := cmd.arguments[0]
	feedurl := cmd.arguments[1]
	createdAt := time.Now()
	updatedAt := time.Now()
	feedId := uuid.New()
	userId := user.ID
	followId := uuid.New()
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID: feedId, CreatedAt: createdAt, UpdatedAt: updatedAt, Name: feedName, Url: feedurl, UserID: userId,
	})
	if err != nil {
		return err
	}
	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID: followId, CreatedAt: createdAt, UpdatedAt: updatedAt, UserID: userId, FeedID: feedId,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed.Name)
	fmt.Println(feed.ID)
	fmt.Println(feed.Url)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.UserID)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
		fmt.Println(feed.Username)
	}
	return nil
}

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 1 {
		return fmt.Errorf("Not enough arguments, please provide url")
	}
	feedUrl := cmd.arguments[0]
	feed, err := s.db.GetFeed(context.Background(), feedUrl)
	if err != nil {
		return err
	}
	createdAt := time.Now()
	updatedAt := time.Now()
	followId := uuid.New()
	follow, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{ID: followId, CreatedAt: createdAt, UpdatedAt: updatedAt, UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return err
	}
	fmt.Println(follow.FeedName)
	fmt.Println(follow.UserName)
	return nil
}

func handlerFollows(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
	}
	fmt.Println("Followed by: ", user.Name)
	return nil
}
