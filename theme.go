package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var themeDir = fmt.Sprintf("%s/theme", restartConfigDirPath)

func useTheme(theme string) error {
	if theme != "default" {
		if err := downTheme(theme); err != nil {
			return err
		}

		rc.Theme = theme
	}
	
	return loadTheme()
}

func downTheme(theme string) error {

	if err := downIndex(theme); err != nil {
		return err
	}

	return downPage(theme)
}

func downIndex(theme string) error {
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/cdn-blog/theme/master/%s/index.html", theme))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	newThemeDir := fmt.Sprintf("%s/%s", themeDir, theme)
	if _, err := os.Stat(newThemeDir); os.IsNotExist(err) {
		os.MkdirAll(newThemeDir, 0700)
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/index.html", newThemeDir), data, 0700)
}

func downPage(theme string) error {
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/cdn-blog/theme/master/%s/page.html", theme))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	newThemeDir := fmt.Sprintf("%s/%s", themeDir, theme)
	if _, err := os.Stat(newThemeDir); os.IsNotExist(err) {
		os.MkdirAll(newThemeDir, 0700)
	}

	return ioutil.WriteFile(fmt.Sprintf("%s/page.html", newThemeDir), data, 0700)
}

func loadTheme() error {
	if _, err := os.Stat(themeDir); os.IsNotExist(err) {
		defaultThemeDir := fmt.Sprintf("%s/default", themeDir)
		os.MkdirAll(defaultThemeDir, 0700)
		err := ioutil.WriteFile(fmt.Sprintf("%s/index.html", defaultThemeDir), []byte(indexTPL), 0700)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(fmt.Sprintf("%s/page.html", defaultThemeDir), []byte(htmlTPL), 0700)
	}

	dir := fmt.Sprintf("%s/%s", themeDir, rc.Theme)
	data, err := ioutil.ReadFile(fmt.Sprintf("%s/index.html", dir))
	if err != nil {
		return err
	}

	indexTPL = string(data)

	data, err = ioutil.ReadFile(fmt.Sprintf("%s/page.html", dir))
	if err != nil {
		return err
	}

	htmlTPL = string(data)

	return nil
}
