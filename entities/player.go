package entities

import (
	"net"
)

type Player struct {
	Id string `json:"PlayerId"`
	Position Vector `json:"Position"`
	Conn *net.UDPAddr `json:"-"`
}

var defaultPlayersStorage *PlayerStorage

type PlayerStorage struct {
	Players map[string]*Player
	Conn *net.UDPConn
}

func (store *PlayerStorage) AddPlayer(key string, player Player) {
	store.Players[key] = &player
}

func (store *PlayerStorage) RemovePlayer(key string) {
	delete(store.Players, key)
}

func CreatePlayerStorage() *PlayerStorage {
	return &PlayerStorage{Players: map[string]*Player{}}
}
//TODO: Make thread safe
func GetDefaultPlayersStorage() *PlayerStorage {
	if defaultPlayersStorage == nil {
		defaultPlayersStorage = CreatePlayerStorage()
	}
	return defaultPlayersStorage
}

