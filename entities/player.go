package entities

import "sync"

type Player struct {
	Id       string       `json:"PlayerId"`
	Position Vector       `json:"Position"`
}

var (
	globalPlayerStore *PlayerStore
	mutex sync.Mutex
)


type PlayerStore struct {
	players map[string]*Player
	mutex   sync.Mutex
}

func CreatePlayerStore() *PlayerStore {
	return &PlayerStore{players: map[string]*Player{}}
}

func GetOrInitGlobalPlayerStore() *PlayerStore {
	if globalPlayerStore == nil {
		mutex.Lock()
		if globalPlayerStore == nil {
			globalPlayerStore = CreatePlayerStore()
		}
		mutex.Unlock()
	}
	return globalPlayerStore
}

func (store *PlayerStore) AddPlayer(key string, player *Player) {
	store.players[key] = player
}

func (self *PlayerStore) GetPlayer(key string) *Player {
	return self.players[key]
}

func (store *PlayerStore) RemovePlayer(key string) {
	delete(store.players, key)
}

func (self *PlayerStore) GetAllPlayers() (list []*Player) {
	self.mutex.Lock()
	for _, v := range self.players {
		list = append(list, v)
	}
	self.mutex.Unlock()
	return list
}


