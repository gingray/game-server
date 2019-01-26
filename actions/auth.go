package actions

import (
	"encoding/json"
	"game-server/entities"
	"github.com/satori/go.uuid"
)

func HandleAuth(evt *entities.Evt) {
	id, _ := uuid.NewV4()
	player := entities.Player{Id: id.String(), Position: entities.Vector{X: 0, Y: 0}}
	entities.Players.Store(id, player)
	resp, _ :=json.Marshal(player)
	ServerConn.WriteToUDP(resp, evt.Addr)

}



