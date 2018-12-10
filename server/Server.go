package server

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"time"
)

type Operation interface {
	StartServer()
}
type server struct {
	ListenPort     int
	secretKey      string
	liveClientBook LiveClientBookAction
}

func (server *server) StartServer() {
	conn, err := net.Listen("tcp", "127.0.0.1"+":"+strconv.Itoa(server.ListenPort))
	if nil != err {
		panic(err)
	}
	for {
		connection, _ := conn.Accept()
		connection.SetReadDeadline(time.Now().Add(2 * time.Second))
		go handleConnection(connection, server)
	}
}
func handleConnection(connection net.Conn, server *server) {
	defer connection.Close()

	cache := make([]byte, 512)
	content := make([]byte, 0)
	for {
		length, err := connection.Read(cache)
		content = append(content, cache[0:length]...)
		if io.EOF != err {
			break
		}
	}
	code := string(content)
	code = strings.TrimSpace(code)
	if strings.Contains(code, server.secretKey) {
		server.liveClientBook.AddLiveClient(code)
		_, err := connection.Write([]byte(connection.RemoteAddr().String()))
		if nil != err {
			fmt.Println(err.Error())
		}
	}
}
func Server(port int, key string, liveClientBook LiveClientBookAction) server {
	return server{ListenPort: port, secretKey: key, liveClientBook: liveClientBook}
}
