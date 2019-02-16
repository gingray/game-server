package actions

import (
	"encoding/json"
	"fmt"
	"game-server/entities"
)

const RIGHT = "RIGHT"
const LEFT = "LEFT"
const UP = "UP"
const DOWN = "DOWN"

type Command struct {
	PlayerId string `json:"PlayerId"`
	Command  string `json:"Command"`
}

func HandleCommand(evt *Evt) {
	playersStore := entities.GetOrInitGlobalPlayerStore()
	var command *Command
	_ = json.Unmarshal(evt.Payload, &command)
	player:= playersStore.GetPlayer(command.PlayerId)
	switch command.Command {
	case RIGHT:
		player.Position.X += 1
	case LEFT:
		player.Position.X -= 1
	case UP:
		player.Position.Y += 1
	case DOWN:
		player.Position.Y -= 1
	default:
		fmt.Println("Unknown command: %s", command.Command)
	}
}
