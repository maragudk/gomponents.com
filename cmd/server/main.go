package main

import (
	"errors"
	"log"
	"net/http"
)

func main() {
	if err := start(); err != nil {
		log.Fatalln("Error:", err)
	}
}

func start() error {
	h := http.FileServer(http.Dir("docs"))
	if err := http.ListenAndServe(":8080", h); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
