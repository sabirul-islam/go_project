package main

import (
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/", home)
	http.Handle("/resources", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, `welcome`)
}