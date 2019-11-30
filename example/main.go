package main

import (
	"net/http"

	"github.com/roger-king/tasker"
	tm "github.com/roger-king/tasker/models"
)

func main() {
	t := tasker.New(tm.TaskerConfig{
		Type:               "mongo",
		MongoConnectionURL: tm.MongoConnectionURL("mongodb://appuser:appuser@localhost:27017/tasker?authSource=admin"),
	})
	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
