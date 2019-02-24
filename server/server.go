package server

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func NewServer(transport DataStateReaderSocketStore, port int) {
	ServerAddr, _ := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	serverConn, _ := net.ListenUDP("udp", ServerAddr)
	fmt.Println("Server start")
	var wg sync.WaitGroup

	defer func() {
		_ = serverConn.Close()
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go listener(transport, serverConn, &wg)
	}
	wg.Add(1)
	go broadcastState(transport, serverConn, &wg)

	wg.Wait()
}

func listener(dataReaderSocketStore DateReaderSocketStore, sc *net.UDPConn, wg *sync.WaitGroup) {
	var buf = make([]byte, 1024)

	defer wg.Done()
	for {
		size, addr, err := sc.ReadFromUDP(buf)
		dataReaderSocketStore.ReadData(buf[:size], addr)
		dataReaderSocketStore.Add(addr)
		fmt.Println("Msg received: %s", string(buf[:size]))

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

func broadcastState(writerStore StateReaderSocketStore, sc *net.UDPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	var nanoFrame int64 = 50000000
	for {
		start := time.Now()
		for _, addr := range writerStore.All() {
			data:= writerStore.ReadState()
			fmt.Printf("Reponse: %s\n", string(data))
			sc.WriteToUDP(data, addr)
		}
		elapsed := time.Since(start).Nanoseconds()

		if elapsed > nanoFrame {
			toSleep := elapsed - nanoFrame
			time.Sleep(time.Duration(toSleep))
		}
	}
}
