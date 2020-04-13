package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

//config 全局配置参数
type config struct {
	url  string `toml:"url"`
	name string `toml:"name"`
}

type restoreConfig struct {
	Md      metadata   `toml:"-"`
	Docs    []string   `toml:"docs"`
	History []metadata `toml:"history"`
}

type runtime struct {
	docs    map[string]bool
	history map[string]bool
}

func initConfig() (config, error) {
	if os.Getenv("CB_CONF") != "" {
		return initConfigViaFile(os.Getenv("CB_CONF"))
	}

	return initConfigViaEnv()
}

func initConfigViaFile(path string) (config, error) {
	var c config
	_, err := toml.DecodeFile(path, &c)

	return c, err
}

func initConfigViaEnv() (config, error) {
	var c config

	return c, nil
}
