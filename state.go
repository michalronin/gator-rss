package main

import (
	"github.com/michalronin/gator/internal/config"
	"github.com/michalronin/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
