package server

import (
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

func broadCastState(sc  *net.UDPConn, userConnections *UserConnections, wg *sync.WaitGroup) {
	defer wg.Done()
	var nanoFrame int64 = 50000000
	for {
		start := time.Now()
		conns := userConnections.GetAllConnections()
		for _, addr := range conns {
			fmt.Printf("I am here")
			sc.WriteToUDP([]byte("1231"), addr)
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
	defaultPlayersStore := entities.GetDefaultPlayersStorage()
	defaultPlayersStore.Conn = serverConn

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go listener(serverConn, userConnections, &wg)
	}
	wg.Add(1)
	go broadCastState(serverConn, userConnections, &wg)


	wg.Wait()

	//buf := make([]byte, 1024)
	//
	//for {
	//	n, addr, err := serverConn.ReadFromUDP(buf)
	//	evt := FetchEvt(buf[:n])
	//	evt.Addr = addr
	//	fmt.Println("Connection from: ",addr)
	//	fmt.Println("Data received: ",string(buf[:n]))
	//	Route(evt)
	//
	//	if err != nil {
	//		fmt.Println("Error: ",err)
	//	}
	//}
}
