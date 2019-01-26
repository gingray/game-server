package main

import (
	"fmt"
	"game-server/actions"
	. "game-server/entities"
	"net"
	"os"
	"sync"
)

//func auth(w http.ResponseWriter, r *http.Request) {
//	message := r.URL.Path
//	message = strings.TrimPrefix(message, "/")
//	message = "Hello " + message
//	fmt.Println(message)
//	w.Write([]byte(message))
//}
//
//func MovePlayer(w http.ResponseWriter, r *http.Request) {
//	body, _ := ioutil.ReadAll(r.Body)
//	defer r.Body.Close()
//	var player Player
//	json.Unmarshal(body, &player)
//}
//
//func AddPlayer(w http.ResponseWriter, r *http.Request) {
//	id, _ := uuid.NewV4()
//
//	player := Player{Id: id.String(), Position: Vector{X: 0, Y: 0}}
//	players.Store(id, player)
//	resp, err :=json.Marshal(player)
//	if err !=nil {
//		panic(err)
//	}
//	w.Write(resp)
//}
//
//func GameState(w http.ResponseWriter, r *http.Request) {
//	var listPlayers []Player
//	players.Range(func(_, value interface{}) bool {
//		listPlayers = append(listPlayers,value.(Player))
//
//		return true
//	})
//	resp, _ := json.Marshal(listPlayers)
//	w.Write(resp)
//}

var mutex = &sync.Mutex{}

func CheckError(err error) {
	if err  != nil {
		fmt.Println("Error: " , err)
		os.Exit(0)
	}
}

func Route(evt *Evt) {
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



//nc -u localhost 10001 < auth.json

func main() {
	ServerAddr, err := net.ResolveUDPAddr("udp",":10001")
	CheckError(err)

	actions.ServerConn, _ = net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer actions.ServerConn.Close()
	buf := make([]byte, 1024)

	for {
		n,addr,err := actions.ServerConn.ReadFromUDP(buf)
		evt := FetchEvt(buf[0:n])
		evt.Addr = addr
		Route(evt)
		fmt.Println("Received ",string(buf[0:n]), " from ",addr)

		if err != nil {
			fmt.Println("Error: ",err)
		}
	}
}
