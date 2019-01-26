package entities

import "sync"

type Player struct {
	Id string `json:"id"`
	Position Vector `json:"position"`
}

var Players = &sync.Map{}
