package main

import (
	"log"
	"net/http"
	"os"

	"./handlers"
)

func main() {

	// References
	l := log.New(os.Stdout, "api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":888", sm)
}
