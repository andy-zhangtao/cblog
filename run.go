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

func run() error {

	var err error

	rc, err = loadRestoreConfig()
	if err != nil {
		return err
	}

	rt = parseRC(rc)

	if err := useTheme(*theme); err != nil {
		return err
	}

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
				//当需要重建整个目录时，直接读取当前所有文件
				buildFile = append(buildFile, info.Name())
			} else {
				//挑选出从来没有构建过的markdown文件
				if _, exist := rt.docs[info.Name()]; !exist {
					buildFile = append(buildFile, info.Name())
				}
			}
		}
		return nil
	})

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

	for i, d := range rc.Docs {
		if hrer, err := generateHtml(d, rc); err != nil {
			return err
		} else {
			rc.History[i].Href = hrer
			rt.upload = append(rt.upload, hrer)
		}

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

	href, err := generateHtml(name, rc)
	if err != nil {
		return err
	}
	md.Href = href
	rc.Md = md
	rt.upload = append(rt.upload, href)

	//过滤是否存在重复内容。如果存在重复内容，从History剔除掉。否则Index会出现重复数据
	filterRestoreConfig(name)
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

func filterRestoreConfig(name string) {
	_, exist := rt.docs[name]
	if !exist {
		return
	}
	delete(rt.docs, name)

	for i, d := range rc.History {
		if d.Doc == name {
			if i == len(rc.History)-1 {
				rc.History = rc.History[:i]
				rc.Docs = rc.Docs[:i]
				return
			}

			rc.History = append(rc.History[:i], rc.History[i+1:]...)
			rc.Docs = append(rc.Docs[:i], rc.Docs[i+1:]...)
			return
		}
	}
}

//generateRestoreConfig 备份当前配置
func generateRestoreConfig(name string, md metadata) {
	//如果当前文件从来没有构建过，直接添加到History。
	//如果已经构建过，需要替换掉History中的记录
	_, exist := rt.docs[name]
	if !exist {
		rc.Docs = append(rc.Docs, name)
		rc.History = append(rc.History, md)
		return
	}

	for i, d := range rc.History {
		if d.Doc == name {
			rc.History[i] = md
		}
	}
}
