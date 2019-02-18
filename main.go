package main

import (
	"game-server/entities"
	"game-server/server"
)


//nc -u localhost 10001 < auth.json

func main() {
	game:=entities.NewGame()
	transport:=server.NewTransport(game)
	server.NewServer(transport, 10001)
}
