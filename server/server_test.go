package server

import (
	"fmt"
	"game-server/test"
	"net"
	"testing"
)

func init() {

	go func() {
		StartServer(10002)
	}()
}

func TestNETServer_Run(t *testing.T) {
	conn, err := net.Dial("udp", ":10002")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()
}

func TestNETServer_SendAuthEvt(t *testing.T) {
	conn, err := net.Dial("udp", ":10002")
	resp,_ := test.AuthOnServer(conn)
	fmt.Printf("%s", string(resp))
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()
}

func TestNETServer_SendCommand(t *testing.T) {
	conn, err := net.Dial("udp", ":10002")
	_, playerId := test.AuthOnServer(conn)

	bytes :=test.GetAuthEventFixture("command_1.json", &playerId)
	bytes2 :=test.GetAuthEventFixture("command_2.json", &playerId)
	resp:=test.WriteAndReadResponse(conn, bytes)
	resp =test.WriteAndReadResponse(conn, bytes)
	resp =test.WriteAndReadResponse(conn, bytes2)
	fmt.Printf("resp: %s\n", resp)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
}
