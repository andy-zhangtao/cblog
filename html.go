package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"text/template"
	"time"
)

type html struct {
	Content string
}

func generateHtml(name string) (string, error) {
	origiHtml, err := parseMarkdown(name)
	if err != nil {
		return "", err
	}

	t := time.Now().Format("2006-01-02")
	pwd, _ := os.Getwd()

	dir := fmt.Sprintf("%s/%s", pwd, t)
	os.MkdirAll(dir, 0700)
	_md5 := md5.Sum([]byte(origiHtml))
	path := fmt.Sprintf("%s/%x.html", dir, _md5)
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	_t, _ := template.New("html").Parse(htmlTPL)

	return fmt.Sprintf("%s/%x.html", t, _md5), _t.Execute(f, &html{Content: origiHtml})
}
