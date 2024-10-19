package main

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"sync"
	"text/template"
)

type File struct {
	Name         string   `json:"name"`
	Modifiers    []string `json:"modifiers"`
	TemplatePath string   `json:"templatePath"`
	TargetPath   string   `json:"targetPath"`
}

type Files struct {
	Files []File `json:"files"`
}

func fillFileWithContent(file File, config Config) {
	content := renderFileTemplate(file, config)
	saveFile(file, content, config)
}

func renderFileTemplate(file File, config Config) []byte {
	log("Start render - " + file.Name)

	var fileContent bytes.Buffer
	path := config.TemplatePath + file.TemplatePath
	name := filepath.Base(path)
	tmpl, err := template.New(name).ParseFiles(path)
	check(err)

	err = tmpl.Execute(&fileContent, config)
	check(err)

	log("Finish render - " + file.Name)
	return fileContent.Bytes()
}

func filterFiles(files Files, config Config) Files {
	log("Start file filtering")
	resultFiles := &Files{}

	for _, file := range files.Files {
		if file.hasOneOfModifiers(config.Modifiers) {
			resultFiles.Files = append(resultFiles.Files, file)
		}
	}

	log("Finish file filtering")
	return *resultFiles
}

func (this *File) HasModifier(modifier string) bool {
	for _, fileMod := range this.Modifiers {
		if fileMod == modifier {
			return true
		}
	}

	return false
}

func (this *File) hasOneOfModifiers(mods []string) bool {
	for _, fileMod := range this.Modifiers {
		for _, modifier := range mods {
			if fileMod == modifier {
				return true
			}
		}
	}

	return false
}

func parseFilesFromJson(config Config) Files {
	var files Files
	var buffer bytes.Buffer

	log("Start json parsing")
	jsonFile, err := os.Open(config.JsonPath)
	check(err)

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	check(err)
	tmpl, err := template.New("").Parse(string(byteValue))
	tmpl.Execute(&buffer, config)
	check(err)

	json.Unmarshal(buffer.Bytes(), &files)

	log("Finish json parsing")
	return files
}

func placeFiles(files Files, config Config) {
	log("Start file placement")
	wg := sync.WaitGroup{}

	for _, file := range files.Files {
		wg.Add(1)
		go func(file File, config Config) {
			defer wg.Done()
			fillFileWithContent(file, config)

		}(file, config)
	}

	wg.Wait()
	log("Finish file placement")
}
