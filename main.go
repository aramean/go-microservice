package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"./handlers"
)

func main() {

	l := log.New(os.Stdout, "api", log.LstdFlags)

	// create handlers
	hh := handlers.NewHello(l)
	ah := handlers.NewAssigments(l)

	// create server and register handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/assigments", ah)

	// start a new server
	s := &http.Server{
		Addr:         ":888",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	signal.Notify(sc, os.Kill)

	sig := <-sc
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
