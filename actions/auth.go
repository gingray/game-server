package actions

import (
	"game-server/entities"
	"github.com/satori/go.uuid"
)

func HandleAuth(evt *Evt) {
	id, _ := uuid.NewV4()
	player := entities.Player{Id: id.String(), Position: entities.Vector{X: 0, Y: 0}}
	playersStore := entities.GetOrInitGlobalPlayerStore()
	playersStore.AddPlayer(id.String(), &player)
}
