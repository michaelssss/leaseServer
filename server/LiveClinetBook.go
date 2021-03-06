package server

import (
	"encoding/json"
	"time"
)

type liveClient struct {
	ExpireTime     time.Time `json:"expireTime"`
	ClientIndetify string `json:"clientIndetify"`
}

type liveClientBook struct {
	liveClients []liveClient
}

type LiveClientBookAction interface {
	GetAllClientJSON() string
	AddLiveClient(clientIndetify string)
	IsAlive(clientIndetify string) bool
}

func (liveClientBook *liveClientBook) GetAllClientJSON() string {
	jsonbyte, _ := json.Marshal(liveClientBook.liveClients)
	return string(jsonbyte)
}

func (liveClientBook *liveClientBook) IsAlive(clientIndetify string) bool {
	for _, liveClient := range liveClientBook.liveClients {
		if liveClient.ClientIndetify == clientIndetify && time.Now().Before(liveClient.ExpireTime) {
			return true
		}
	}
	return false
}
func (liveClientBook *liveClientBook) AddLiveClient(clientIndetify string) {
	clean(liveClientBook)
	if !liveClientBook.IsAlive(clientIndetify) {
		liveClient := liveClient{ClientIndetify: clientIndetify, ExpireTime: time.Now().Add(time.Second * 15)}
		liveClientBook.liveClients = append(liveClientBook.liveClients, liveClient)
	}
}
func clean(book *liveClientBook) {
	client := book.liveClients
	indexs := []int{}
	for index, value := range client {
		if time.Now().After(value.ExpireTime) {
			indexs = append(indexs, index)
		}
	}
	newClient := []liveClient{}
	for index, value := range client {
		if !isElementContain(&indexs, index) {
			newClient = append(newClient, value)
		}
	}
	book.liveClients = newClient
}
func isElementContain(array *[]int, element int) bool {
	for _, value := range *array {
		if value == element {
			return true
		}
	}
	return false
}
func NewLiveClientBook() LiveClientBookAction {
	liveClients := []liveClient{}
	liveClientBook := liveClientBook{liveClients: liveClients}
	return &liveClientBook
}
