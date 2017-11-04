package main

import (
	"./domain"
	"fmt"
	"os"
	"net"
	"net/http"
)

func main() {
	server := domain.Server{ListenPort: 8888, ClientList: make([]domain.Client, 0)}
	http.HandleFunc("/getClient", server.GetAllClients)
	go http.ListenAndServe(":8080", nil)

	server.StartServer()

	//client := domain.Client{ClientName: "bbbbb", ClientAddr: getAddr().(*net.IPNet).IP.String()}
	//client.MakeDiscover()
}
func getAddr() net.Addr {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return address
			}

		}
	}
	return nil
}
