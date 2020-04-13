package main

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var defaultMetadataMD5 = []byte{
	'1',
	'2',
	'3',
	'4',
	'5',
}

type metadata struct {
	Title     string   `toml:"title"`
	Date      string   `toml:"date"`
	Thumbnail []string `toml:"thumbnail"`
	Summary   string   `toml:"summary"`
	Category  string   `toml:"category"`
	Tags      []string `toml:"tags"`
	Href      string   `toml:"href"`
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

func calcMetadataMD5(md metadata) string {
	var b bytes.Buffer

	encode := toml.NewEncoder(&b)
	err := encode.Encode(md)
	if err != nil {
		fmt.Printf("Calc %d md5 error. Use default md5.")
		return fmt.Sprintf("%x", md5.Sum(defaultMetadataMD5))
	}

	return fmt.Sprintf("%x", md5.Sum(b.Bytes()))
}
