package server

import (
	"encoding/json"
	"net/http"
)

type jsonResp struct {
	IP string
}
type myHandleHttpStuct struct {
	key            string
	liveClientBook LiveClientBookAction
}
type MyHandleHttp interface {
	HandleHttp(w http.ResponseWriter, req *http.Request)
}

func (myHandleHttp myHandleHttpStuct) HandleHttp(w http.ResponseWriter, req *http.Request) {
	accessKey := req.Header.Get("accessKey")
	if "" == accessKey {
		accessKey = req.URL.Query().Get("accessKey")
	}
	if accessKey == myHandleHttp.key {
		ip1 := req.Header.Get("X-real-ip")
		if "" == ip1 {
			ip1 = req.RemoteAddr
		}
		respString := jsonResp{IP: ip1}
		jsonbyte, _ := json.Marshal(respString)
		myHandleHttp.liveClientBook.AddLiveClient(accessKey)
		w.Write(jsonbyte)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func (myHandleHttp myHandleHttpStuct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	myHandleHttp.HandleHttp(w, req)
}
func MyHandleHttpStuct(key1 string, liveClientBook LiveClientBookAction) http.Handler {
	return myHandleHttpStuct{key: key1, liveClientBook: liveClientBook}
}
