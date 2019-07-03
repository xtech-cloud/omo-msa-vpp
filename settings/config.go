package settings

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type WebSocketSection struct {
	ReadTimeoutMs  int `yaml:"read_timeout"`
	WriteTimeoutMs int `yaml:"write_timeout"`
	IdleTimeoutMs  int `yaml:"idle_timeout"`
}

type HttpSection struct {
	Filter  FilterSection    `yaml:"filter"`
	Forward []ForwardSection `yaml:"forward"`
}

type FilterSection struct {
	Mode string   `yaml:"mode"`
	URL  []string `yaml:"url"`
}

type ForwardSection struct {
	Endpoint string `yaml:"endpoint"`
	Remote   string `yaml:"remote"`
}

type Config struct {
	Http      HttpSection      `yaml:"http"`
	WebSocket WebSocketSection `yaml:"websocket"`
}

var config Config

func SetupConfig() {
	content, err := ioutil.ReadFile(GetEnv().CONFIG_PATH)
	if nil != err {
		panic(err)
	}

	err = yaml.Unmarshal(content, &config)
	if nil != err {
		panic(err)
	}
}

func GetConfig() *Config {
	return &config
}
