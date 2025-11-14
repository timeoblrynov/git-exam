package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "route: %q", r.URL.Path)
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {  
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")  
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/color", ColorHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
