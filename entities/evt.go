package entities

import (
	"encoding/json"
	"net"
)

type Evt struct {
	Type string `json:"Evt"`
	Payload json.RawMessage `json:"Payload"`
	Addr *net.UDPAddr `json:"-"`
}

func FetchEvt(msg []byte) (evt *Evt) {
	_ = json.Unmarshal(msg, &evt)
	return evt
}
