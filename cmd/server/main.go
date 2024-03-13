package main

import (
	"errors"
	"fmt"
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
	fmt.Println("Serving at http://localhost:8080")
	if err := http.ListenAndServe(":8080", h); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
