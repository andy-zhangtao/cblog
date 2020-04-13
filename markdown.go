package main

import (
	"io/ioutil"

	"gopkg.in/russross/blackfriday.v2"
)

func parseMarkdown(path string) (html string, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	//output HTML
	out := blackfriday.Run(data)
	html = string(out)
	return
}
