package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/michalronin/gator/internal/database"
)

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username required")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Println("user not found")
		os.Exit(1)
	}
	s.cfg.SetUser(cmd.args[0])
	fmt.Printf("user %v logged in\n", cmd.args[0])
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username required")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err == nil {
		fmt.Println("user with that name already exists")
		os.Exit(1)
	}
	if err == sql.ErrNoRows {
		user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      cmd.args[0],
		})
		if err != nil {
			return err
		}
		s.cfg.SetUser(user.Name)
		fmt.Printf("user has been created: %v", user)
	}
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.Reset(context.Background()); err != nil {
		fmt.Println("database reset failed", err)
		os.Exit(1)
	}
	fmt.Println("database reset successful")
	os.Exit(0)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %v (current)\n", user.Name)
		} else {
			fmt.Printf("* %v\n", user.Name)
		}
	}
	return nil
}

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) < 2 {
		fmt.Println("two parameters required: feed name and feed url")
		os.Exit(1)
	}
	currentUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    currentUser.ID,
	})
	if err != nil {
		return err
	}
	fmt.Println(feed)
	os.Exit(0)
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}
	for _, feed := range feeds {
		fmt.Printf("* Name: %v, URL: %v, created by: %v\n", feed.Name, feed.Url, feed.Username)
	}
	return nil
}
