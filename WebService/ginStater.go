package WebService

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartWebService(tcpLister net.TCPListener) {
	go func() {
		r := gin.Default()
		fmt.Println("Web Service GIN Start From TCP IP ")
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		if err := r.RunListener(tcpLister); err != nil {
			fmt.Printf("\033[1A\r\nGin web service stopped:\n>", err)
		}
	}()

}
