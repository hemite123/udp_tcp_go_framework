package WebService

import (
	"context"
	"fmt"
	"gameserver/model"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	mu          sync.Mutex
	webServer   *http.Server
	oldListener net.Listener
)

func StartWebService(tcpLister net.Listener) {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Recovery())
		fmt.Printf("\033[1A\r\nWeb Service GIN Start From Port%s\n", tcpLister.Addr())
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		mu.Lock()
		webServer = &http.Server{
			Handler: r,
		}
		oldListener = tcpLister
		mu.Unlock()
		if err := webServer.Serve(tcpLister); err != nil && err != http.ErrServerClosed {
			fmt.Printf("\033[1A\r\nGin web ervice stopped:\n>", err)
		}
	}()

}

func RestartWebService() {
	mu.Lock()
	defer mu.Unlock()
	tcpAddr, _err := net.ResolveTCPAddr("tcp", model.LoadConfig.TCPPort)
	oldListener.Close()
	if _err != nil {
		fmt.Println(_err.Error())
	}
	tcpLister, _err := net.ListenTCP("tcp", tcpAddr)
	if _err != nil {
		fmt.Println(_err.Error())
	}
	if webServer != nil {
		fmt.Println("\n[WebService] Shutting down current web server...")
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if err := webServer.Shutdown(ctx); err != nil {
			fmt.Printf("[WebService] Forced shutdown: %v\n", err)
		}
	}

	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("\033[1A\r\nRecovered from panic in Go Routing Restart Web Service: %s \n>", r)
			}
		}()
		router := gin.New()
		router.Use(gin.Recovery())

		router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Server Is Restarted " + time.Now().String()})
		})

		mu.Lock()
		webServer = &http.Server{Handler: router}
		oldListener = tcpLister
		mu.Unlock()

		fmt.Printf("\033[1A\r\nWeb Service GIN Restarted on Port %s\n> ", tcpLister.Addr())
		if err := webServer.Serve(tcpLister); err != nil && err != http.ErrServerClosed {
			fmt.Printf("\033[1A\r\nGin web service error: %v ", err)
		}
	}()
}
