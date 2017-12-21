package server

import (
	"net/http"
	"encoding/json"
)

type jsonResp struct {
	IP string
}
type myHandleHttpStuct struct {
	key string
}
type MyHandleHttp interface {
	HandleHttp(w http.ResponseWriter, req *http.Request)
}

func (myHandleHttp myHandleHttpStuct) HandleHttp(w http.ResponseWriter, req *http.Request) {
	accessKey := req.Header.Get("accessKey")
	if accessKey == myHandleHttp.key {
		respString := jsonResp{IP: req.RemoteAddr}
		jsonbyte, _ := json.Marshal(respString)
		header:=w.Header()
		header.Add("hello","world")
		w.Write(jsonbyte)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
func (myHandleHttp myHandleHttpStuct) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	myHandleHttp.HandleHttp(w, req)
}
func MyHandleHttpStuct(key1 string) http.Handler {
	return myHandleHttpStuct{key: key1}
}
