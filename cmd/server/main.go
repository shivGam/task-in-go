package main

import (
	"github.com/shivGam/task-in-go/internal/config"
	"github.com/shivGam/task-in-go/internal/db"
)

func main() {

	config.LoadEnv()
	db.Connect()
	defer db.Close()

}