package main

import (
	"log"
	"net/http"

	"github.com/roger-king/tasker"
	"github.com/roger-king/tasker/config"
)

func main() {
	t, err := tasker.New(&config.TaskerConfig{
		Migrate:            true,
		Auth:               false,
		DBConnectionURL:    "postgres://appuser:appuser@localhost:5432/tasker?sslmode=disable",
		GithubClientID:     "",
		GithubClientSecret: "",
	})

	if err != nil {
		log.Panic(err)
	}

	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
