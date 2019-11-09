package main

import (
	"net/http"

	"github.com/roger-king/tasker"
	"github.com/roger-king/tasker/pkg"
)

func main() {
	t := tasker.New(&tasker.TaskerConfig{
		Details: &pkg.ConnectionDetails{
			Host:     "localhost:27017",
			DBName:   "tasker",
			User:     "appuser",
			Password: "appuser",
		},
	})
	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
