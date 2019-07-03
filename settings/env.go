package settings

import "os"

type Env struct {
	HTTP_ADDR   string
	HTTPS_ADDR  string
	WS_ADDR     string
	CONFIG_PATH string
	TLS_CRT     string
	TLS_KEY     string
}

var env Env

func SetupEnv() {
	env.HTTP_ADDR = os.Getenv("VPP_HTTP_ADDR")
	if "" == env.HTTP_ADDR {
		env.HTTP_ADDR = ":80"
	}
	env.HTTPS_ADDR = os.Getenv("VPP_HTTPS_ADDR")
	if "" == env.HTTPS_ADDR {
		env.HTTPS_ADDR = ":443"
	}
	env.WS_ADDR = os.Getenv("VPP_WS_ADDR")
	if "" == env.WS_ADDR {
		env.WS_ADDR = ":8000"
	}

	env.CONFIG_PATH = os.Getenv("VPP_CONFIG")
	if "" == env.CONFIG_PATH {
		env.CONFIG_PATH = "./conf/vpp.yaml"
	}

	env.TLS_CRT = os.Getenv("VPP_TLS_CRT")
	if "" == env.TLS_CRT {
		env.TLS_CRT = "./conf/tls.crt"
	}

	env.TLS_KEY = os.Getenv("VPP_TLS_KEY")
	if "" == env.TLS_KEY {
		env.TLS_KEY = "./conf/tls.key"
	}
}

func GetEnv() *Env {
	return &env
}
