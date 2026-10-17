package WebService

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartWebService(tcpLister net.Listener) {
	go func() {
		gin.SetMode(gin.ReleaseMode)
		r := gin.New()
		r.Use(gin.Recovery())
		fmt.Println("Web Service GIN Start From Port ", tcpLister.Addr())
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		if err := r.RunListener(tcpLister); err != nil {
			fmt.Printf("\033[1A\r\nGin web ervice stopped:\n>", err)
		}
	}()

}
