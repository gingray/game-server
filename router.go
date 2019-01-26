package main

import (
	"game-server/actions"
	"game-server/entities"
)

func Route(evt *entities.Evt) {
	switch evt.Type{
	case "AUTH":
		actions.HandleAuth(evt)
	case "LIST":
		actions.HandleList(evt)
	case "STATE":
		actions.HandleState(evt)
	default:
		println("Unknown evt: [%s]", evt.Type)
	}
}
