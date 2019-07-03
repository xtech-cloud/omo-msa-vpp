package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

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
	Http HttpSection `yaml:"http"`
}

var config Config

func setupConfig() {
	content, err := ioutil.ReadFile(env_config_path)
	if nil != err {
		panic(err)
	}

	err = yaml.Unmarshal(content, &config)
	if nil != err {
		panic(err)
	}
}
