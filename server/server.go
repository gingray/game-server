package server

import (
	"encoding/json"
	"fmt"
	"game-server/actions"
	"game-server/entities"
	"net"
	"sync"
	"time"
)

func listener(sc *net.UDPConn, userConnections *UserConnections, wg *sync.WaitGroup) {
	var buf = make([]byte, 1024)

	defer wg.Done()
	for {
		size, addr, err := sc.ReadFromUDP(buf)
		userConnections.AddConnection(addr)
		evt := actions.CreateEvt(addr, buf, size)
		evt.ProcessEvt()

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func broadcastState(sc  *net.UDPConn, userConnections *UserConnections, wg *sync.WaitGroup) {
	defer wg.Done()
	var nanoFrame int64 = 50000000
	for {
		start := time.Now()
		conns := userConnections.GetAllConnections()
		playerStore:= entities.GetOrInitGlobalPlayerStore()
		players:= playerStore.GetAllPlayers()
		data,_ := json.Marshal(&players)
		for _, addr := range conns {
			if len(players) > 0 {
				fmt.Printf("Reponse: %s\n", string(data))
				sc.WriteToUDP(data, addr)
			}
		}
		elapsed :=time.Since(start).Nanoseconds()

		if elapsed > nanoFrame  {
			toSleep:= elapsed - nanoFrame
			time.Sleep(time.Duration(toSleep))
		}
	}
}

func StartServer(port int) {
	ServerAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	serverConn, _ := net.ListenUDP("udp", ServerAddr)
	fmt.Println("Server start")

	defer func() {
		_ = serverConn.Close()
	}()

	var wg sync.WaitGroup
	userConnections := CreateUserConnectionsStore()
	entities.GetOrInitGlobalPlayerStore()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go listener(serverConn, userConnections, &wg)
	}
	wg.Add(1)
	go broadcastState(serverConn, userConnections, &wg)


	wg.Wait()
}
