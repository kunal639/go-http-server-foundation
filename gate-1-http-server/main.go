// A handler is a thing that knows how to respond to a request.
// In a backend system: Request comes in; Something handles it; Response goes out. The "something" is the handler

// Handler is a contract between our code and HTTP server because HTTP server needs a generic way to call user code. It only knows call: ServeHTTP
// More formally: Handler is a single agreed entrypoint through which the HTTP server transfrs control to user code.


/*
	What happens when i run this program:
	1. OS starts the program
	2. GO runtime starts
	3. main() is called
	4. We register a handler
	5. We start an HTTP server
	6. The program blocks
	7. Main goroutine remains blocked, handler executes in a new goroutine
	8. We write a rsponse
	9. Server keeps running
*/

package main
// There's exactly only one main package 
// It is required by the GO runtime to create an executable. It must be prsent. Without it no compilation

import (	// packages we are currently using. Acts as an abstraction by removing the complexity of reating everything from scrach
	"net/http"	// This knows TCP listening, HTTP parsing, concurrency, request lifecycle. Called by our code. Uess OS networking internally
	"io"	// We want to write bytes to our rsponse. It is used by our handler 
	"log"	// To print fatal errors cleanly. Used only when server fails to start.
)

func main() {	// Called by GO runtime immediately when program starts

	// This is a function. Does not do anything unless called. It will be called by HTTP server
	// w http.ResponseWriter is an interface provided by net/http. It reprsents the response being consructed. 
	// w http.ResponseWriter is created by net/http server. Owned by the server. We can write to it and control what is written, but we do not control when it is sent or how the connection is managed.
	// r *http.Request represents the incoming http request. Created by net/http.
	// r *http.Request contains: Method, URL, Headers, Body, Context 
	helloHandler := func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK) // sets http status code to 200. sent by the server when response headers are flushed.
		io.WriteString(w, "Hello World") // Writes response body in bytes. The bytes ar sent by the server. We only provide data 
	}

	// It registers heeloHandler inside http.DefaultServeMux. Simply means if reuest comes to /health, call this function.
	// It is used by the http server.
	http.HandleFunc("/health", helloHandler)
	
	/*
		http.ListenAndServer(":8000",nil). What it does:
		1. Creates an http server
		2. Binds to port 8000
		3. Starts listening to TCP connections
		4. Accepts connections
		5. Parse HTTP requests
		6. Find matching handler
		7. Calls handler in a goroutine
		8. Calls ServeHTTP
		9. Writes response
		10. Repeats

		It is a blocking function as it blocksthe main goroutine to keep the server alive and running
	*/
	err := http.ListenAndServe(":8000",nil)

	if err != nil {	// runs only if server fails to start
		log.Fatal(err)
	}
}