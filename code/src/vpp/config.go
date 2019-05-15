package main

import (
	"encoding/json"
	"io/ioutil"
)

type FilterSection struct {
	Mode string   `json:"mode"`
	URL  []string `json:"url"`
}

type ForwardSection struct {
	Endpoint string `json:"endpoint"`
	Remote   string `json:"remote"`
}

type Config struct {
	Filter  FilterSection    `json:"filter"`
	Forward []ForwardSection `json:"forward"`
}

var config Config

func setupConfig() {
	content, err := ioutil.ReadFile(env_config_path)
	if nil != err {
		panic(err)
	}

	err = json.Unmarshal(content, &config)
	if nil != err {
		panic(err)
	}
}
