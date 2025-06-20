package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter,r *http.Request) {
	if err := r.ParseForm(); err!=nil {
		fmt.Fprintf(w,"ParseForm() error %v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")


	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)

	fmt.Fprintf(w, "hello!")
}


func helloHandler(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w,"404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET"{
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello", helloHandler)


	fmt.Printf("Starting server at 8080 \n")
	if err := http.ListenAndServe(":8080",nil); err!=nil {
		log.Fatal(err)
	}
}