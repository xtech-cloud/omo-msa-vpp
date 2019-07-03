package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/xtech-cloud/omo-msa-vpp/settings"
)

type Http struct {
}

func (this *Http) Start() {
	router := gin.Default()
	router.Use(filterMiddleware())
	router.Any("/*action", forward)
	httpAddr := settings.GetEnv().HTTP_ADDR
	httpsAddr := settings.GetEnv().HTTPS_ADDR
	crt := settings.GetEnv().TLS_CRT
	key := settings.GetEnv().TLS_KEY
	go router.RunTLS(httpsAddr, crt, key)
	go router.Run(httpAddr)
}

func (this *Http) Stop() {
}
