package actions

import (
	"encoding/json"
	"fmt"
	"net"
)

type Evt struct {
	Type    string          `json:"Evt"`
	Payload json.RawMessage `json:"Payload"`
	Addr    *net.UDPAddr    `json:"-"`
}

func CreateEvt(addr *net.UDPAddr, msg []byte, size int) (evt *Evt) {
	_ = json.Unmarshal(msg[:size], &evt)
	evt.Addr = addr
	fmt.Println("Connection from: ", addr)
	return evt
}

func (self *Evt) ProcessEvt() {
	fmt.Printf("evt: [%s] from: [%s]\n", self.Type, self.Addr.String())

	switch self.Type {
	case "AUTH":
		HandleAuth(self)
	case "COMMAND":
		HandleCommand(self)
	default:
		println("Unknown evt: [%s]", self.Type)
	}
}
