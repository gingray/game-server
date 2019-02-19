package server

import (
	"net"
	"sync"
)

type UserConnections struct {
	connections map[string]*net.UDPAddr
	mutex       sync.Mutex
}

func (self *UserConnections) AddConnection(conn *net.UDPAddr) {
	self.mutex.Lock()
	self.connections[conn.String()] = conn
	self.mutex.Unlock()
}

func (self *UserConnections) GetAllConnections() (list []*net.UDPAddr) {
	for _, v := range self.connections {
		list = append(list, v)
	}
	return list
}
