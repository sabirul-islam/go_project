package main

import (
	"fmt"
	"net/http"
)

/*
type myType string
func (myType) ServeHTTP(w http.ResponseWriter, r *http.Request){
	// fmt.Fprint(w, `welcome`, r.URL.Path, r.URL.RawPath, r.Method, r.RemoteAddr)

	if r.URL.Path=="/portfolio"{
		fmt.Fprint(w, "welcome to portfolio page.", r.URL.Path)
	}else if r.URL.Path=="/contact"{
		fmt.Fprint(w, "welcome to contact page.", r.URL.Path)

	}else if r.URL.Path=="/about"{
		fmt.Fprint(w, "welcome to about page.", r.URL.Path)

	}else {
		fmt.Fprint(w, "welcome to home page.", r.URL.Path)
	}
}
*/

func main() {
	// var handler myType
	// http.ListenAndServe(":8888", handler)

	http.HandleFunc("/", home)
	http.HandleFunc("/portfolio", portfolio)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)

	http.ListenAndServe(":8888", nil)
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `welcome to home page.`,r.URL.Path)
}

func portfolio(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `welcome to portfolio page.`,r.URL.Path)
}

func about(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `welcome to about page.`,r.URL.Path)
}

func contact(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `welcome to contact page.`,r.URL.Path)
}