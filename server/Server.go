package server

import (
	"net"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ServerOperations interface {
	StartServer()
	AddClient(client Client)
	RemoveClient(client Client)
	GetAllClients(w http.ResponseWriter, r *http.Request)
}
type Server struct {
	ListenPort int
	ClientList []Client
}

func (server *Server) StartServer() {
	conn, err := net.Listen("tcp", "0.0.0.0:8888")
	if nil != err {
		panic(err)
	}
	for {
		connection, _ := conn.Accept()
		go handleConnection(connection, err, server)
	}
}
func handleConnection(connection net.Conn, err error, server *Server) {
	defer connection.Close()
	cache := make([]byte, 512)
	content := make([]byte, 0)
	for {
		length, err := connection.Read(cache)
		content = append(content, cache[0: length]...)
		if io.EOF != err {
			break
		}
	}
	community := Community{int(content[0]), content[1:]}
	client := Client{}
	err = json.Unmarshal(community.content, &client)
	if nil != err {
		fmt.Println(err)
	}
	client.ClientAddr = connection.RemoteAddr().String()
	server.AddClient(client)
	fmt.Println(server.ClientList)
}
func (server *Server) AddClient(client Client) {
	clients := server.ClientList
	for index, client1 := range clients {
		if client1.ClientName == client.ClientName {
			clients[index].ClientAddr = client.ClientAddr
			return
		}
	}
	server.ClientList = append(clients, client)
}
func (server *Server) RemoveClient(client Client) {
	clients := &server.ClientList
	removeList := make([]int, 0)
	for index, value := range *clients {
		if value.ClientName == client.ClientName {
			removeList = append(removeList, index)
		}
	}
}
func (server *Server) GetAllClients(w http.ResponseWriter, r *http.Request) {
	content, _ := json.Marshal(server.ClientList)
	w.Write(content)
}
