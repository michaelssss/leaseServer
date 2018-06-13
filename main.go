package main

import (
	"leaseServer/server"
	"net/http"
	_ "net/http/pprof"
	"os"
)

func main() {
	liveClientBook := server.NewLiveClientBook()
	ss := os.Args
	communityKey := ss[1]
	server1 := server.Server(8888, communityKey, liveClientBook)
	http.Handle("/", server.MyHandleHttpStuct(communityKey, liveClientBook))
	http.HandleFunc("/getLiveClient", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(liveClientBook.GetAllClientJSON()));
	})
	go http.ListenAndServe("0.0.0.0:8889", nil)
	server1.StartServer()
}
