package domain

import (
	"net"
	"encoding/json"
	"fmt"
	"io"
)

type ServerOperations interface {
	StartServer()
	AddClient(client Client)
	RemoveClient(client Client)
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
	server.AddClient(client)
	fmt.Println(server.ClientList)
}
func (server *Server) AddClient(client Client) {
	clients := server.ClientList
	for _, client1 := range clients {
		if client1.ClientName == client.ClientName {
			client1.ClientAddr = client.ClientAddr
			return
		}
	}
	server.ClientList = append(clients, client)
}
func (server *Server) RemoveClient(client Client) {
	clients := server.ClientList
	removeList := make([]int, 0)
	for index, value := range clients {
		if value.ClientName == client.ClientName {
			removeList = append(removeList, index)
		}
	}
	for _, value := range removeList {
		clients = removeValueFromArray(clients, value)
	}
	server.ClientList = clients
}
func removeValueFromArray(clients []Client, target int) []Client {
	tem := clients[0:target-1]
	tem2 := clients[target+1:len(clients)-1]
	return append(tem, tem2...)
}
