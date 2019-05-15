package main

import "os"

var env_addr string
var env_config_path string

func setupEnv() {
	env_addr = os.Getenv("VPP_ADDR")
	if "" == env_addr {
		env_addr = ":16000"
	}

	env_config_path = os.Getenv("VPP_CONFIG")
	if "" == env_config_path {
		env_config_path = "./conf/vpp.cfg"
	}
}
