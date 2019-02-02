package main

import (
	"fmt"
	. "game-server/entities"
	"game-server/utils"
	"net"
)

var (
	ServerConn *net.UDPConn
)

//nc -u localhost 10001 < auth.json

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp",":10001")
	utils.CheckError(err)

	ServerConn, _ = net.ListenUDP("udp", ServerAddr)
	defaultPlayersStore := GetDefaultPlayersStorage()
	defaultPlayersStore.Conn = ServerConn
	fmt.Println("Server start")

	defer func() {
		_= ServerConn.Close()
	}()

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		evt := FetchEvt(buf[0:n])
		evt.Addr = addr
		Route(evt)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)

		if err != nil {
			fmt.Println("Error: ",err)
		}
	}
}
