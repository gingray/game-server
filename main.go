package main

import (
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	. "game-server/entities"
)

func auth(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	fmt.Println(message)
	w.Write([]byte(message))
}

func MovePlayer(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var player Player
	json.Unmarshal(body, &player)
}

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	id, _ := uuid.NewV4()

	player := Player{Id: id.String(), Position: Vector{X: 0, Y: 0}}
	players.Store(id, player)
	resp, err :=json.Marshal(player)
	if err !=nil {
		panic(err)
	}
	w.Write(resp)
}

func GameState(w http.ResponseWriter, r *http.Request) {
	var listPlayers []Player
	players.Range(func(_, value interface{}) bool {
		listPlayers = append(listPlayers,value.(Player))

		return true
	})
	resp, _ := json.Marshal(listPlayers)
	w.Write(resp)
}

var players = &sync.Map{}
var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/Auth", auth)
	http.HandleFunc("/MovePlayer", MovePlayer)
	http.HandleFunc("/AddPlayer", AddPlayer)
	http.HandleFunc("/GameState", GameState)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
