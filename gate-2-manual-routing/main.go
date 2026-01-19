package main

import (
	"io"
	"log"
	"net/http"
)

type myHandler struct {}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method
	if (path == "/health") {
		if (method == "GET") {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "Method Not Allowed")
		}
	} else if(path == "/users") {
		if (method == "GET") {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		} else if (method == "POST") {
			w.WriteHeader(http.StatusCreated)
			io.WriteString(w, "Created")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "Method Not Allowed")
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, "Not Found")
	}
}

func main() {
	http.Handle("/", myHandler{})

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}