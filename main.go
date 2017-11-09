package main

import (
	"net/http"
)

func main() {
	server := Server{ListenPort: 8888, ClientList: make([]Client, 0)}
	http.HandleFunc("/getClient", server.GetAllClients)
	go http.ListenAndServe(":8080", nil)
	server.StartServer()
}
