package server

import (
	"fmt"
	"game-server/utils"
	"net"
	. "game-server/entities"
)

var (
	ServerConn *net.UDPConn
)

func StartServer(port int) {
	ServerAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
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
