package main

import (
	"log"
	"net/http"
)


func main() {
	server := http.NewServeMux()
	log.Println("Server start at port 8030")
	log.Fatal(http.ListenAndServe(":8030", server))
}