package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

var rc restoreConfig
var rt runtime

func beforeRun() {
	if *rebuild {
		clearRC()
		*wholeDir = true
	}

}
func run() error {

	beforeRun()

	var err error

	rc, err = loadRestoreConfig()
	if err != nil {
		return err
	}

	rt = parseRC(rc)

	if *wholeDir {
		return buildWholeDir()
	}

	if *markdownFile == "" {
		return fmt.Errorf("Dir or makedown file must exist. ")
	}

	return buildMarkdownFile(*markdownFile)
}

func buildWholeDir() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	var buildFile []string
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			if *rebuild {
				buildFile = append(buildFile, info.Name())
			} else {
				if _, exist := rt.docs[info.Name()]; !exist {
					buildFile = append(buildFile, info.Name())
				}
			}
		}
		return nil
	})

	fmt.Printf("%#v\n", buildFile)
	if len(buildFile) > 0 {
		return batchBuildMarkdownFile(buildFile, *rebuild)
	}

	return nil
}

func batchBuildMarkdownFile(fileList []string, updateHistory bool) error {
	if updateHistory {
		rc.History = make([]metadata, len(fileList))

		for i, f := range fileList {
			md, err := parseMetadata(f)
			if err != nil {
				return err
			}
			rc.History[i] = md
		}

		rc.Docs = fileList
	} else {
		for _, f := range fileList {
			md, err := parseMetadata(f)
			if err != nil {
				return err
			}
			rc.History = append(rc.History, md)
		}
		rc.Docs = append(rc.Docs, fileList...)
	}

	err := generateIndex(rc)
	if err != nil {
		return err
	}

	return saveRestoreConfig(rc)
}

func buildMarkdownFile(name string) error {

	md, err := parseMetadata(name)
	if err != nil {
		return err
	}

	rc.Md = md

	err = generateIndex(rc)
	if err != nil {
		return err
	}

	generateRestoreConfig(name, md)
	return saveRestoreConfig(rc)
}

func generateIndex(rc restoreConfig) error {

	t, err := template.New("index").Parse(indexTPL)
	if err != nil {
		return err
	}
	f, err := os.Create("index.html")
	if err != nil {
		return err
	}

	defer f.Close()

	return t.Execute(f, &rc)
}

func generateRestoreConfig(name string, md metadata) {
	if _, exist := rt.docs[name]; !exist {
		rc.Docs = append(rc.Docs, name)
	}

	if _, exist := rt.history[calcMetadataMD5(md)]; !exist {
		rc.History = append(rc.History, md)
	}
}
