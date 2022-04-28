package initialize

import (
	"github.com/gin-gonic/gin"
	"liuqichat/servers/wb"
	"net/http"
	"runtime"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/sys", func(c *gin.Context) {
		data := map[string]interface{}{
			"当前用户": len(wb.GetClientManager().Clients),
			"goroutineNum": runtime.NumGoroutine(),
			"当前节点CPU": runtime.NumCPU(),
			"集群数量": 0,
		}
		c.JSON(http.StatusOK, data)
	})
	return r
}
