package main

import "os"

var env_http_addr string
var env_config_path string

func setupEnv() {
	env_http_addr = os.Getenv("VPP_HTTP_ADDR")
	if "" == env_http_addr {
		env_http_addr = ":16000"
	}

	env_config_path = os.Getenv("VPP_CONFIG")
	if "" == env_config_path {
		env_config_path = "./conf/vpp.cfg"
	}
}
