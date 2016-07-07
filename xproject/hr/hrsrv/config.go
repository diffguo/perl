package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ServiceConfig struct {
	Custom *CustomConfig `json:"custom"`
}

type CustomConfig struct {
	HttpListenAddr string `json:"listen"`
}

func loadConfig(filename string) (*ServiceConfig, error) {

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cfg := &ServiceConfig{}
	err = json.Unmarshal(d, cfg)
	if err != nil {
		return nil, err
	}

	fmt.Println("Load CustomConfig:")
	if cfg.Custom != nil {
		fmt.Println(*cfg.Custom)
	}
	fmt.Println()

	return cfg, nil
}
