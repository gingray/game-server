package actions

import (
	"encoding/json"
	"game-server/entities"
	"github.com/satori/go.uuid"
)

func HandleAuth(evt *entities.Evt) {
	id, _ := uuid.NewV4()
	player := entities.Player{Id: id.String(), Position: entities.Vector{X: 0, Y: 0}, Conn: evt.Addr}
	defautPlayersStore := entities.GetDefaultPlayersStorage()
	defautPlayersStore.Players[id.String()]= player
	resp, _ :=json.Marshal(player)
	_,_ = defautPlayersStore.Conn.WriteToUDP(resp, evt.Addr)
}



