package main

import (
	"fmt"
	"net/http"
)

//w: wr,ter,
func index_handler(w http.ResponseWriter, r *http.Request) {
	//Fprint is similar like format?
	fmt.Fprint(w, "Awesome web app!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	//Fprint is similar like format?
	fmt.Fprint(w, "This app is a simple web app!")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)
	http.ListenAndServe(":8000", nil)
}
