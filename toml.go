package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

type metadata struct {
	Title     string   `toml:"title"`
	Date      string   `toml:"date"`
	Thumbnail []string `toml:"thumbnail"`
	Summary   string   `toml:"summary"`
	Category  string   `toml:"category"`
	Tags      []string `toml:"tags"`
}

//parseMetadata 从指定的markdown文件中解析元数据
func parseMetadata(path string) (md metadata, err error) {
	f, err := os.Open(path)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(f)

	var tomls []byte
	save := false
	for scanner.Scan() {

		if strings.TrimSpace(scanner.Text()) == "<!--" {
			save = true
			continue
		}

		if strings.TrimSpace(scanner.Text()) == "-->" {
			break
		}

		if save {
			tomls = append(tomls, scanner.Bytes()...)
			tomls = append(tomls, '\n')
		}
	}

	err = toml.Unmarshal(tomls, &md)

	return
}
