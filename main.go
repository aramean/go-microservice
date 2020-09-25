package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"./handlers"
)

func main() {

	// References
	l := log.New(os.Stdout, "api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	s := &http.Server{
		Addr:         ":888",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	s.ListenAndServe()
}
