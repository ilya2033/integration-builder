package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
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
	files := parseJson(config)
	files = filterFiles(files, config)
	placeFiles(files, config)
}

func placeFiles(files Files, config Config) {
	fmt.Println("Start file placement")
	for _, file := range files.Files {
		content := renderFileTemplate(file, config)
		saveFile(file, content, config)
	}
	fmt.Println("Finish file placement")
}

func saveFile(file File, content []byte, config Config) {
	path := config.TargetPath + file.TargetPath
	_, err := os.Stat(path)

	if err == nil {
		fmt.Println(path + " already exists. Skip")
		return
	}

	err = os.MkdirAll(filepath.Dir(path), fs.ModePerm)
	check(err)

	newFile, err := os.Create(path)
	check(err)

	defer newFile.Close()

	newFile.Write(content)
	newFile.Sync()

	fmt.Println("File saved - " + file.Name)
}

func renderFileTemplate(file File, config Config) []byte {
	fmt.Println("Start render - " + file.Name)

	var fileContent bytes.Buffer
	path := config.TemplatePath + file.TemplatePath
	name := filepath.Base(path)
	tmpl, err := template.New(name).ParseFiles(path)
	check(err)

	err = tmpl.Execute(&fileContent, config)
	check(err)

	fmt.Println("Finish render - " + file.Name)
	return fileContent.Bytes()
}

func filterFiles(files Files, config Config) Files {
	fmt.Println("Start file filtering")
	resultFiles := &Files{}

	for _, file := range files.Files {
		if file.hasOnOfModifiers(config.Modifiers) {
			resultFiles.Files = append(resultFiles.Files, file)
		}
	}

	fmt.Println("Finish file filtering")
	return *resultFiles
}

func parseJson(config Config) Files {
	var files Files
	var buffer bytes.Buffer

	fmt.Println("Start json parsing")
	jsonFile, err := os.Open(config.JsonPath)
	check(err)

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	check(err)
	tmpl, err := template.New("").Parse(string(byteValue))
	tmpl.Execute(&buffer, config)
	check(err)

	json.Unmarshal(buffer.Bytes(), &files)

	fmt.Println("Finish json parsing")
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
