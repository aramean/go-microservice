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

func (a *Assigments) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch method := r.Method; method {
	case http.MethodGet:
		a.GetAssigments(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (a *Assigments) GetAssigments(rw http.ResponseWriter, r *http.Request) {
	a.l.Println("Hello Assigments!")

	la := data.GetAssigments()
	err := la.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
