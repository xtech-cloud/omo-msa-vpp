package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/xtech-cloud/omo-msa-vpp/gateway"
	"github.com/xtech-cloud/omo-msa-vpp/settings"
)

func main() {
	settings.SetupEnv()
	settings.SetupConfig()

	httpGateway := &gateway.Http{}
	httpGateway.Start()

	wsGateway := &gateway.WebSocket{}
	wsGateway.Start()

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-c

	httpGateway.Stop()
	wsGateway.Stop()

	os.Exit(0)
}
