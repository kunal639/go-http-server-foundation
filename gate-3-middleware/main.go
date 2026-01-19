package main

import (
	"io"
	"log"
	"net/http"
	"time"
	"fmt"
)

type myHandler struct {}

func (m myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	if (path == "/health") {
		if(method == http.MethodGet) {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		} else {
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, "Method Not Allowed")
		}
	} else if (path == "/users") {
		if (method == http.MethodGet) {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "OK")
		} else if (method == http.MethodPost) {
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

func logging (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s", r.Method, r.RequestURI)
		next.ServeHTTP(w,r)
	})
}

func timing (next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w,r)
		duration := time.Since(start)
		fmt.Printf("%s", duration)
	})
}

func main () {
	
	router := myHandler{}
	chain := timing(logging(router))

	http.Handle("/", chain)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}