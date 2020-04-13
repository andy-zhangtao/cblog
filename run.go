package main

import (
	"fmt"
	"html/template"
	"os"
)

func run() error {
	if *wholeDir {

	}

	if *markdownFile == "" {
		return fmt.Errorf("Dir or makedown file must exist. ")
	}

	return buildMarkdownFile(*markdownFile)
}

func buildWholeDir() error {
	return nil
}

func buildMarkdownFile(name string) error {

	md, err := parseMetadata(name)
	if err != nil {
		return err
	}

	return generateIndex(md)
}

func generateIndex(md metadata) error {

	t, err := template.New("index").Parse(indexTPL)
	if err != nil {
		return err
	}
	f, err := os.Create("index.html")
	if err != nil {
		return err
	}

	defer f.Close()

	return t.Execute(f, &md)
}
