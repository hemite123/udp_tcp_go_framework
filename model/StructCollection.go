package model

import "encoding/xml"

var ListCommand = make(map[string]func())
var LoadConfig Config

type Config struct {
	XMLName      xml.Name `xml:"Config"`
	Port         string   `xml:"Port"`
	ShutdownWait int      `xml:"ShutdownWait"`
}

func CommandSetter(commandName string, act func()) {
	ListCommand[commandName] = act

}
