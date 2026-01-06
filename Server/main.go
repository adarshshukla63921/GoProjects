package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello world"))
}

func aboutUs(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Tihs page is created by adarsh"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/aboutUs", aboutUs)
	log.Print("Server starting on 8080")
	err := http.ListenAndServe(":8080",mux)
	log.Fatal(err)
}