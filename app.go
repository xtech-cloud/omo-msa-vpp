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
	go router.RunTLS(env_https_addr, env_tls_crt, env_tls_key)
	router.Run(env_http_addr)
}
