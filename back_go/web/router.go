package web

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	Port = 8081
)

func RunHttp() error {
	r := gin.Default()
	r.Use(CorsConfig())
	router := r.Group("/student")
	{
		router.POST("/save", Save)
		router.GET("/findAll", FindAll)
		router.GET("/findById/:id", FindById)
		router.POST("/update", Update)
		router.DELETE("/remove/:id", Remove)
	}
	return r.Run(fmt.Sprintf("127.0.0.1:%d", Port))
}

func CorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
