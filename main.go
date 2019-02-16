package main

import (
	"game-server/server"
)


//nc -u localhost 10001 < auth.json

func main() {
	server.StartServer(10001)
}
