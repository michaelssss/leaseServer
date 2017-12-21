package main

import (
	"leaseServer/server"
	"os"
	"net/http"
)

func main() {
	ss := os.Args
	communityKey := ss[1]
	server1 := server.Server(8888, communityKey)
	http.Handle("/", server.MyHandleHttpStuct(communityKey))
	go http.ListenAndServe("0.0.0.0:8889", nil)
	server1.StartServer()
}
