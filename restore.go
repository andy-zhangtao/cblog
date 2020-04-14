package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

var restartConfigDirPath = fmt.Sprintf("%s/.cblogs", os.Getenv("HOME"))
var restartConfigPath = fmt.Sprintf("%s/runtime.toml", restartConfigDirPath)
var globalConfigPath = fmt.Sprintf("%s/cb.toml", restartConfigDirPath)

func beforeRun() {
	if *rebuild {
		clearRC()
		*wholeDir = true
	}

	if *preview {
		rc.Conf.Url = fmt.Sprintf("http://localhost:%d", *port)
	}
}

func loadRestoreConfig() (restoreConfig, error) {

	var rc restoreConfig

	_, err := os.Stat(restartConfigPath)
	if os.IsNotExist(err) {
		os.MkdirAll(restartConfigDirPath, 0700)
		return rc, nil
	}
	_, err = toml.DecodeFile(restartConfigPath, &rc)
	rc.Conf = loadGlobalConfig()

	beforeRun()

	return rc, err
}

func loadGlobalConfig() config {

	var cf config
	_, err := os.Stat(globalConfigPath)
	if os.IsNotExist(err) {
		return cf
	}

	toml.DecodeFile(globalConfigPath, &cf)
	return cf
}

func saveRestoreConfig(rc restoreConfig) error {
	var b bytes.Buffer
	encode := toml.NewEncoder(&b)
	err := encode.Encode(rc)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(restartConfigPath, b.Bytes(), 0700)
}

func parseRC(rc restoreConfig) runtime {
	rt = runtime{
		docs:    map[string]bool{},
		history: map[string]bool{},
		upload:  []string{},
	}

	for _, d := range rc.Docs {
		rt.docs[d] = true
	}

	for _, h := range rc.History {
		rt.history[calcMetadataMD5(h)] = true
	}

	return rt
}

func clearRC() {
	rc.Docs = nil
	rc.History = nil
}
