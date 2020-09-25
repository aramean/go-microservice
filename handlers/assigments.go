package handlers

import (
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

}
