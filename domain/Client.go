package domain

import (
	"net"
	"encoding/json"
)

type Client struct {
	ClientName string
	ClientAddr string
}
type ClientOperations interface {
	MakeDiscover()
	Renew()
}

func (client *Client) MakeDiscover() {
	conn, err := net.Dial("tcp", "47.89.21.243:8888")
	defer conn.Close()
	result, _ := json.Marshal(client)
	if nil != err {
		panic(err)
	}
	content := Community{len(result), result}.ToByte()
	conn.Write(content)
}
func (client *Client) Renew() {
	client.MakeDiscover()
}
