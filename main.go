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

	// create serve mux and register handlers
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/assigments", ah)

	// create a new server
	s := &http.Server{
		Addr:         ":888",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// graceful shutdown server
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	signal.Notify(sc, os.Kill)

	sig := <-sc
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
