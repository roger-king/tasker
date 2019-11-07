package main

import (
	"net/http"

	"github.com/roger-king/tasker"
	"github.com/roger-king/tasker/pkg"
)

func main() {
	t := tasker.New(&tasker.TaskerCongfig{
		ConnectionType: "redis",
		Details: &pkg.ConnectionDetails{
			Host: "localhost:6379",
			DB:   0,
		},
	})
	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
