package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	// if we define a variable, and not use it, then that is a error
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s", name)
}

// response and request
// *  -->  r is pointing to a pointer
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprint(w, "hello!")
}

func main() {
	// :=   --> short form assignment
	// we tell it to look into static folder, and it auto choses index.html file like any other server in the market
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	// err will either be nil if no error, or it stores the error if any
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
