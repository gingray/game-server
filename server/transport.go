package server

import (
	"game-server/entities"
	"net"
)

type BytesDataReader interface {
	ReadData(msg []byte)
}

type BytesStateReader interface {
	ReadState() []byte
}

type SocketStore interface {
	Add(addr *net.UDPAddr)
	All() []*net.UDPAddr
}

type StateReaderSocketStore interface {
	BytesStateReader
	SocketStore
}

type DataStateReaderSocketStore interface {
	BytesDataReader
	BytesStateReader
	SocketStore
}

type DateReaderSocketStore interface {
	BytesDataReader
	SocketStore
}

type Transport struct {
	game  *entities.Game
	addrs []*net.UDPAddr
}

func NewTransport(game *entities.Game) *Transport {
	return &Transport{game: game, addrs: []*net.UDPAddr{}}
}

func (self *Transport) ReadData(msg []byte) {
	self.game.Fetch(string(msg))
}

func (self *Transport) ReadState() []byte {
	data := self.game.Broadcast()
	return []byte(data)
}

func (self *Transport) Add(addr *net.UDPAddr) {
	self.addrs = append(self.addrs, addr)
}

func (self *Transport) All() []*net.UDPAddr {
	tmp := make([]*net.UDPAddr, len(self.addrs))
	return tmp
}
