package test

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"os"
	"path"
	"strings"
)
const PACKAGE_NAME = "game-server"

func GetAuthEventFixture(filename string, playerId *string) []byte {
	dir, _ := os.Getwd()
	fixturePath := path.Join(strings.Split(dir,PACKAGE_NAME)[0], "game-server/test/fixtures/", filename)
	bytes, _:=ioutil.ReadFile(fixturePath)
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