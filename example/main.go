package main

import (
	"log"
	"net/http"

	"github.com/roger-king/tasker"
)

func main() {
	t, err := tasker.New()

	if err != nil {
		log.Panic(err)
	}

	router := t.Start()

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
