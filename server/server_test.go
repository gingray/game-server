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

func GetAuthEventFixture(filename string, playerId *string) []byte {
	bytes, _:=ioutil.ReadFile(filepath.Join("../test-fixture", filename))
	if playerId != nil {
		var extra map[string]interface{}
		json.Unmarshal(bytes, &extra)
		extra["Payload"].(map[string]interface{})["PlayerId"]= playerId
		payload, _ := json.Marshal(extra)
		return payload
	}
	return bytes
}
func AuthOnServer(conn net.Conn) ([]byte, string) {
	authEvtBytes := GetAuthEventFixture("auth.json", nil)
	conn.Write(authEvtBytes)
	out := make([]byte, 1024)
	size, _:=conn.Read(out)
	var data map[string]interface{}
	json.Unmarshal(out[:size], &data)
	playerId := data["PlayerId"].(string)
	return out[:size], playerId
}

func WriteAndReadResponse(conn net.Conn, bytes []byte) string {
	conn.Write(bytes)
	out := make([]byte, 1024)
	rBytes, _ :=conn.Read(out)
	return string(out[:rBytes])
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
	resp,_ := AuthOnServer(conn)
	fmt.Printf("%s", string(resp))
	if err != nil {
		t.Error(err)
	}

	defer conn.Close()
}

func TestNETServer_SendCommand(t *testing.T) {
	conn, err := net.Dial("udp", ":10002")
	_, playerId := AuthOnServer(conn)

	bytes :=GetAuthEventFixture("command_1.json", &playerId)
	bytes_2:=GetAuthEventFixture("command_2.json", &playerId)
	resp:=WriteAndReadResponse(conn, bytes)
	resp =WriteAndReadResponse(conn, bytes)
	resp =WriteAndReadResponse(conn, bytes_2)
	fmt.Printf("resp: %s\n", resp)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
}
