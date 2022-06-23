package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// %v --> default format
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful!")

	// how to get value from form in html?
	// FormValue

	fmt.Fprintf(w, "Post from Website! %v\n ", r.PostForm)
	name := r.FormValue("name")
	adress := r.FormValue("adress")

	//  %s --> Plain string
	fmt.Fprintf(w, "Name =%s\n", name)
	fmt.Fprintf(w, "Adress = %s\n", adress)
}

func main() {
	// declare fileserver
	fileserver := http.FileServer(http.Dir("./static"))

	//handleRoute
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
