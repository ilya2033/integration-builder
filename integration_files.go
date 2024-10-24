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
	Files     []File `json:"files"`
	TestFiles []File `json:"testFiles"`
}

func fillFileWithContent(file File, path string, config Config) {
	content := renderFileTemplate(file, config)
	saveFile(file, content, path)
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

func filterFiles(files []File, config Config) []File {
	log("Start file filtering")
	resultFiles := []File{}

	for _, file := range files {
		if file.hasOneOfModifiers(config.Modifiers) {
			resultFiles = append(resultFiles, file)
		}
	}

	log("Finish file filtering")
	return resultFiles
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

func generateFiles(files []File, basePath string, config Config) {
	log("Start file placement")
	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func(file File, basePath string, config Config) {
			defer wg.Done()
			path := basePath + file.TargetPath
			fillFileWithContent(file, path, config)

		}(file, basePath, config)
	}

	wg.Wait()
	log("Finish file placement")

}

func placeFiles(files []File, config Config) {
	generateFiles(files, config.TargetPath, config)
}
