package main

import (
	"fmt"
	"gameserver/CommandControler"
	"gameserver/WebService"
	"gameserver/model"
	"net"
)

func main() {
	var Command string
	//load Configuration From XMl
	CommandControler.ReloadConfig(true)
	//End Load Configuration

	//Listener Start
	addr, _err := net.ResolveUDPAddr("udp", *&model.LoadConfig.Port)
	if _err != nil {
		fmt.Println(_err.Error())
	}
	lister, _err := net.ListenUDP("udp", addr)
	if _err != nil {
		fmt.Println(_err.Error())
	}
	defer lister.Close()
	fmt.Println("Start Game Server From Port ", *&model.LoadConfig.Port)
	//End Listener

	//Listener TCP
	tcpAddr, _err := net.ResolveTCPAddr("tcp", *&model.LoadConfig.TCPPort)
	if _err != nil {
		fmt.Println(_err.Error())
	}
	liternerTCP, _err := net.ListenTCP("tcp", tcpAddr)
	if _err != nil {
		fmt.Println(_err.Error())
	}
	defer liternerTCP.Close()
	fmt.Println("Start TCP Server On Port ", *&model.LoadConfig.TCPPort)
	WebService.StartWebService(*liternerTCP)
	//END Listerner TCP

	//Package Capture
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, ClientAddr, _err := lister.ReadFromUDP(buffer)
			if _err != nil {
				fmt.Println(_err.Error())
				continue
			}
			fmt.Println("Got Message from %s: %s\n", ClientAddr, string(buffer[:n]))

			_, _err = lister.WriteToUDP(buffer[:n], ClientAddr)
			if _err != nil {
				fmt.Println(_err.Error())
			}
		}
	}()

	//End Package Capture

	//Server Command Execution
	for {
		fmt.Printf(">")
		fmt.Scan(&Command)
		result, _err := MessageCapture(Command)
		if _err != nil {
			fmt.Println(_err.Error())
			continue
		}
		if result == "" {
			continue
		}
		fmt.Println(result)
	}

	//End Server Command Execution

}
