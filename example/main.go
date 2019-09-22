package main

import (
	"net/http"

	"github.com/roger-king/tasker"
)

func main() {
	t := tasker.New()
	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
