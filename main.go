package main

import (
	"net/http"
	"leaseServer/server"
)

func main() {
	server1 := server.Server{ListenPort: 8888, ClientList: make([]server.Client, 0)}
	http.HandleFunc("/getClient", server1.GetAllClients)
	go http.ListenAndServe(":8080", nil)
	server1.StartServer()
}
