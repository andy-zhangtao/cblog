package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

//config 全局配置参数
type config struct {
	url  string `toml:"url"`
	name string `toml:"name"`
}

type runConfig struct {
	docs []string `toml:"docs"`
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

func loadRuntimeConfig() (runConfig, error) {

	var rc runConfig

	if f, err := os.Stat(fmt.Sprintf("%s/.cblogs/runtime.toml", os.Getenv("HOME"))); os.IsExist(err) {
		_, err = toml.DecodeFile(f.Name(), rc)
		return rc, err
	} else {
		return rc, err
	}
}