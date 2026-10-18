package CommandControler

import (
	"encoding/xml"
	"fmt"
	"gameserver/WebService"
	"gameserver/model"
	"os"
	"time"
)

func Shutdown() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\033[1A\r\nRecovered from panic in shutdown command: %s \n>", r)
		}
	}()
	shutdownDuration := time.Duration(*&model.LoadConfig.ShutdownWait) * time.Second
	deadLine := time.Now().Add(shutdownDuration)
	timer := time.NewTimer(shutdownDuration)
	fmt.Println("Server Shutdown In", shutdownDuration.Seconds())
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("\033[1A\r\nRecovered from panic in shutdown GoRoutine Timer: %s \n>", r)
			}
		}()
		for {
			remaining := time.Until(deadLine)
			if remaining <= 5*time.Second && remaining > 0 {
				fmt.Printf("\033[1A\r\nShutting Down In %d\n>", int(remaining.Seconds())+1)
			}
			if remaining <= 0 {
				break
			}
			time.Sleep(1 * time.Second)
		}
		<-timer.C
		//Consider If Have UDP and TCP server For Web And Game Services
		os.Exit(1)
		fmt.Println("Shutdown Server...")
	}()
}

func ReloadConfig(firstLoad bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\033[1A\r\nRecovered from panic in Reload Config: %s \n>", r)
		}
	}()
	file, _err := os.Open("GameConfig.xml")
	if _err != nil {
		fmt.Println(_err.Error())
	}
	decoder := xml.NewDecoder(file)
	if _err := decoder.Decode(&model.LoadConfig); _err != nil {
		fmt.Println(_err.Error())
	}
	if !firstLoad {
		fmt.Printf("\033[1A\r\nConfig Reloaded\n")
	}
}

func RestarWebServices() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("\033[1A\r\nRecovered from panic in Restart Web Service: %s \n>", r)
		}
	}()
	WebService.RestartWebService()
}
