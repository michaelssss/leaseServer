package main

import (
	"leaseServer/server"
	"os"
)

func main() {
	ss := os.Args
	communityKey := ss[1]
	server1 := server.Server(8888, communityKey)
	server1.StartServer()
}
