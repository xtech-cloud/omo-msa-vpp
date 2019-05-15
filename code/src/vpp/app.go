package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	setupEnv()
	setupConfig()
	router := gin.Default()
	router.Use(filterMiddleware())
	router.Any("/*action", forward)
	router.Run(env_http_addr)
}
