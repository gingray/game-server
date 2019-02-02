package entities

import (
	"net"
	"sync"
)

type Player struct {
	Id string `json:"id"`
	Position Vector `json:"position"`
	Conn *net.UDPAddr `json:"-"`
}

var defaultPlayersStorage *PlayerStorage

type PlayerStorage struct {
	Players *sync.Map
	Conn *net.UDPConn
}

func (store *PlayerStorage) AddPlayer(key string, player Player) {
	store.Players.Store(key, player)
}

func (store *PlayerStorage) RemovePlayer(key string) {
	store.Players.Delete(key)
}

func CreatePlayerStorage() *PlayerStorage {
	return &PlayerStorage{Players: &sync.Map{}}
}
//TODO: Make thread safe
func GetDefaultPlayersStorage() *PlayerStorage {
	if defaultPlayersStorage == nil {
		defaultPlayersStorage = CreatePlayerStorage()
	}
	return defaultPlayersStorage
}

