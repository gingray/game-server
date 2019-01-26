package actions

import (
	"encoding/json"
	"game-server/entities"
)


func HandleList(evt *entities.Evt) {
		var listPlayers []entities.Player
		entities.Players.Range(func(_, value interface{}) bool {
			listPlayers = append(listPlayers,value.(entities.Player))

			return true
		})
		resp, _ := json.Marshal(listPlayers)
	ServerConn.WriteToUDP(resp, evt.Addr)

}
