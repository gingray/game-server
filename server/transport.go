package server

import (
	"fmt"
	"game-server/entities"
	"net"
	"sync"
)

type BytesDataReader interface {
	ReadData(msg []byte, addr *net.UDPAddr)
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

type SocketKey struct {
	addr *net.UDPAddr
	key *string
}

type Transport struct {
	game  *entities.Game
	pairAddrIds []*PairAddrId
	addrs map[string]*SocketKey
	mutex sync.RWMutex

}

type PairAddrId struct {
	addr *net.UDPAddr
	id *string
}

func NewTransport(game *entities.Game) *Transport {
	return &Transport{game: game, pairAddrIds: []*PairAddrId{}}
}

func (self *Transport) ReadData(msg []byte, addr *net.UDPAddr) {
	id:=self.game.Fetch(string(msg))
	fmt.Printf("Evt received: %s", msg)
	if self.addrs[addr.String()].key == nil {
		self.addrs[addr.String()].key = id
	}
}

func (self *Transport) ReadState() []byte {
	data := self.game.Broadcast()
	return []byte(data)
}

func (self *Transport) Add(addr *net.UDPAddr) {
	self.mutex.Lock()
	self.addrs[addr.String()] = &SocketKey{addr: addr, key:nil}
	self.mutex.Unlock()
}

func (self *Transport) All() []*net.UDPAddr {
	tmp := make([]*net.UDPAddr, len(self.addrs))
	idx :=0
	self.mutex.RLock()
	fmt.Printf("Broadcast: ")
	for _, item := range self.addrs {
		tmp[idx] = item.addr
		idx+=1
	}
	self.mutex.RUnlock()
	return tmp
}
