package main

import (
	"game-server/server"
	"net"
)

var (
	ServerConn *net.UDPConn
)

//nc -u localhost 10001 < auth.json

func main() {
	server.StartServer(10001)
}
