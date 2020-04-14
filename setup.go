package main

import "io/ioutil"

const configure = `url=""
title=""
summary1=""
summary2=""
favicon=""
[cdn]
	access=""
	secret=""`

func setupUserProfile() error {
	return ioutil.WriteFile(globalConfigPath, []byte(configure), 0700)
}
