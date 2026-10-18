package main

import (
	"errors"
	"gameserver/CommandControler"
	"gameserver/model"
	"strings"
)

func MessageCapture(message string) (string, error) {
	//Check if no have Trail "/" Send Chat
	if strings.Contains(message, "/") {
		runable, exist := model.ListCommand[strings.TrimPrefix(message, "/")]
		if !exist {
			return "Command " + message + " Not Found ", errors.New("Not Found Command")
		}
		runable()
		return "", nil
	}
	return "Message Return", errors.New("Error Code 100")
}

func init() {
	//Add Command Name + Function At Here
	model.CommandSetter("shutdown", CommandControler.Shutdown)
	model.CommandSetter("reload", func() { CommandControler.ReloadConfig(false) })
	model.CommandSetter("webreload", CommandControler.RestarWebServices)
}
