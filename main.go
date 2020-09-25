package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(rw, "Hello World! %s", d)
	})

	http.HandleFunc("/post", func(rw http.ResponseWriter, r *http.Request) {
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Something went wrong", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.ListenAndServe(":888", nil)
}
