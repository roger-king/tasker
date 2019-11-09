package main

import (
	"net/http"

	"github.com/roger-king/tasker"
)

func main() {
	t := tasker.New(&tasker.TaskerConfig{
		MongoConnectionURL: "mongodb://appuser:appuser@localhost:27017/tasker?authSource=admin",
	})
	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
