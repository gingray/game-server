package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"path/filepath"
	"testing"
)

func init() {

	go func() {
		StartServer(10002)
	}()
}

func GetAuthEventFixture(filename string) []byte {
	bytes, _:=ioutil.ReadFile(filepath.Join("../test-fixture", filename))
	return bytes
}
func AuthOnServer(conn net.Conn) []byte {
	authEvtBytes := GetAuthEventFixture("auth.json")
	conn.Write(authEvtBytes)
	out := make([]byte, 1024)
	size, _:=conn.Read(out)
	return out[:size]
}

func TestNETServer_Run(t *testing.T) {
	// Simply check that the server is up and can
	// accept connections.
	conn, err := net.Dial("udp", ":10002")
	if err != nil {
		t.Error("could not connect to server: ", err)
	}
	defer conn.Close()
}

func TestNETServer_SendAuthEvt(t *testing.T) {
	conn, err := net.Dial("udp", ":10002")
	resp := AuthOnServer(conn)
	fmt.Printf("%s", string(resp))
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()
}

func TestNETServer_SendCommand(t *testing.T) {
	conn, err := net.Dial("udp", ":10002")
	resp := AuthOnServer(conn)
	var data map[string]interface{}
	json.Unmarshal(resp, &data)
	playerId := data["PlayerId"].(string)
	var evt map[string]interface{}

	json.Unmarshal(GetAuthEventFixture("command.json"), &evt)
	evt["Payload"].(map[string]interface{})["PlayerId"]= playerId
	bytes, err := json.Marshal(evt)
	conn.Write(bytes)
	out := make([]byte, 1024)
	rBytes, _ :=conn.Read(out)
	fmt.Printf("resp: %s\n", string(out[:rBytes]))
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
}
