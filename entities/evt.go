package entities

import (
	"encoding/json"
	"net"
)

type Evt struct {
	Type string `json:"evt"`
	Payload json.RawMessage `json:"payload"`
	Addr *net.UDPAddr `json:"-"`
}

func FetchEvt(msg []byte) (evt *Evt) {
	json.Unmarshal(msg, &evt)
	return evt
}
