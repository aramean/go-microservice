package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"../data"
)

type Assigments struct {
	l *log.Logger
}

func NewAssigments(l *log.Logger) *Assigments {
	return &Assigments{l}
}

func (a *Assigments) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	a.l.Println("Hello Assigments!")

	la := data.GetAssigments()
	d, err := json.Marshal(la)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	rw.Write(d)
}
