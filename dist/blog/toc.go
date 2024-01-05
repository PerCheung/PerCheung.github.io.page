package main

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	var fileNames []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		if name != "toc" {
			fileNames = append(fileNames, name)
		}
	}

	jsonData, err := json.Marshal(fileNames)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("toc.json", jsonData, 0644)
	if err != nil {
		panic(err)
	}
}
