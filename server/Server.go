package server

import (
	"net"
	"io"
	"strings"
	"fmt"
)

type Operation interface {
	StartServer()
}
type server struct {
	ListenPort int
	secretKey  string
}

func (server *server) StartServer() {
	conn, err := net.Listen("tcp", "127.0.0.1:8888")
	if nil != err {
		panic(err)
	}
	for {
		connection, _ := conn.Accept()
		go handleConnection(connection, server)
	}
}
func handleConnection(connection net.Conn, server *server) {
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
	code := string(content)
	code = strings.TrimSpace(code)
	if strings.Contains(code, server.secretKey) {
		_, err := connection.Write([]byte(connection.RemoteAddr().String()))
		if nil != err {
			fmt.Println(err.Error())
		}
	}
}
func Server(port int, key string) server {
	return server{ListenPort: port, secretKey: key}
}
