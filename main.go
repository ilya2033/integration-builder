package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const PATH_TO_JSON = "files.json"
const PATH_TO_TEMPLATE = "template/"

func main() {
	config := parseConfig()
	files := parseJson(config.JsonPath)
	files = filterFiles(files, config)
	placeFiles(files, config)
}

func placeFiles(files Files, config Config) {
	for _, file := range files.Files {
		content := renderFileTemplate(file, config)
		saveFile(file, content, config)
	}
}

func saveFile(file File, content []byte, config Config) {
	path := config.TargetPath + file.TargetPath
	err := os.MkdirAll(filepath.Dir(path), fs.ModePerm)
	check(err)

	newFile, err := os.Create(path)
	check(err)

	defer newFile.Close()

	newFile.Write(content)
	newFile.Sync()
}

func renderFileTemplate(file File, config Config) []byte {
	var fileContent bytes.Buffer
	path := config.TemplatePath + file.TemplatePath
	name := filepath.Base(path)
	tmpl, err := template.New(name).ParseFiles(path)
	check(err)

	err = tmpl.Execute(&fileContent, config)
	check(err)

	return fileContent.Bytes()
}

func filterFiles(files Files, config Config) Files {
	resultFiles := &Files{}

	for _, file := range files.Files {
		if file.hasOnOfModifiers(config.Modifiers) {
			resultFiles.Files = append(resultFiles.Files, file)
		}
	}

	return *resultFiles
}

func parseJson(path string) Files {
	var files Files

	jsonFile, err := os.Open(path)
	check(err)

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	check(err)

	json.Unmarshal(byteValue, &files)

	return files
}

func parseConfig() Config {
	config := &Config{}
	flows := []string{}

	name := flag.String("name", "", "Integration name")
	flowFlag := flag.String("flow", "", "Integration flow types")
	filesPath := flag.String("files", PATH_TO_JSON, "Json files structure")
	templatePath := flag.String("template", PATH_TO_TEMPLATE, "Template folder")
	targetPath := flag.String("target", "", "Target folder")

	flag.Parse()

	if *name == "" {
		panic(errors.New("Error: Name required"))
	}

	if *targetPath == "" {
		panic(errors.New("Error: Target required"))
	}

	if *flowFlag != "" {
		flows = strings.Split(*flowFlag, "")
	} else {
		panic(errors.New("Error: Flow types reuqired"))
	}

	config.Name = *name
	config.Modifiers = flows
	config.JsonPath = *filesPath
	config.TemplatePath = *templatePath
	config.TargetPath = *targetPath

	return *config
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
