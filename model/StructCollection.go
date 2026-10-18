package model

import (
	"encoding/xml"
	"net"
	"sync"
)

var ListCommand = make(map[string]func())
var LoadConfig Config
var ConnectionList = &ClientConnection{
	client: make(map[string]*net.UDPAddr),
}

type Config struct {
	XMLName      xml.Name `xml:"Config"`
	Port         string   `xml:"Port"`
	ShutdownWait int      `xml:"ShutdownWait"`
	TCPPort      string   `xml:"TCPPort"`
}

type ClientConnection struct {
	mu     sync.RWMutex
	client map[string]*net.UDPAddr
}

func CommandSetter(commandName string, act func()) {
	ListCommand[commandName] = act
}

func AddNewConnection(udpAddres net.UDPAddr) bool {
	ConnectionList.mu.Lock()
	defer ConnectionList.mu.Unlock()
	if _, exists := ConnectionList.client[udpAddres.String()]; !exists {
		ConnectionList.client[udpAddres.String()] = &udpAddres
		return true
	}
	return false
}
