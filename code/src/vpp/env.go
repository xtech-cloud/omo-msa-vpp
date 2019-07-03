package main

import "os"

var env_http_addr string
var env_https_addr string
var env_config_path string
var env_tls_crt string
var env_tls_key string

func setupEnv() {
	env_http_addr = os.Getenv("VPP_HTTP_ADDR")
	if "" == env_http_addr {
		env_http_addr = ":80"
	}
	env_https_addr = os.Getenv("VPP_HTTPS_ADDR")
	if "" == env_https_addr {
		env_https_addr = ":443"
	}

	env_config_path = os.Getenv("VPP_CONFIG")
	if "" == env_config_path {
		env_config_path = "./conf/vpp.yaml"
	}

	env_tls_crt = os.Getenv("VPP_TLS_CRT")
	if "" == env_tls_crt {
		env_tls_crt = "./conf/tls.crt"
	}

	env_tls_key = os.Getenv("VPP_TLS_KEY")
	if "" == env_tls_key {
		env_tls_key = "./conf/tls.key"
	}
}
