package main

import (
	"fmt"
	"net/http"   //its module for creating web servers and handling HTTP requests and responses
	                // if it doesnt exist we need to build it manually using the command go build -o main nov13/main.go nov13/shapes.go
)
func main() {
	mux:= http.NewServeMux() //creates a new instance of ServeMux, which is an HTTP request multiplexer that matches incoming requests to their respective handler functions based on the URL path and HTTP method.

	// mux.Handler("Get"/header", health) //syntax for registering a handler function for a specific HTTP method and URL path. In this case, it registers the health handler function to handle GET requests to the /header path.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rect := Rectangle{Length: 10, Width: 5}
		circ := Circle{Radius: 5}

		fmt.Fprintf(w, "Rectangle Area: %.2f\n", rect.Area())
		fmt.Fprintf(w, "Circle Area: %.2f\n", circ.Area())
	})
	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}


type handler interface{
	Handle(w http.ResponseWriter, r *http.Request)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))


	http.ListenAndServe(":8080", nil)
} 

//mux is a request router and dispatcher that matches incoming requests to their respective handler functions based on the URL path and HTTP method. It allows you to define multiple routes and associate them with specific handler functions, enabling you to build more complex web applications with different endpoints.


// HTTP

//1. Request --- have

// url
// path
// origin
// header

//2. Response --- have

// status
// body
// header
// origin
// 