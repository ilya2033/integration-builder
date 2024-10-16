package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

const PATH_TO_JSON = "files.json"

func main() {
	config, err := parseConfig()

	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := parseJson(config.jsonPath)

	if err != nil {
		fmt.Println(err)
		return
	}

	files = filterFiles(files, config)

	fmt.Println(files)
}

func filterFiles(files Files, config Config) Files {
	resultFiles := &Files{}

	for _, file := range files.Files {
		if file.hasOnOfModifiers(config.modifiers) {
			resultFiles.Files = append(resultFiles.Files, file)
		}
	}

	return *resultFiles
}

func parseJson(path string) (Files, error) {
	var files Files

	jsonFile, err := os.Open(path)

	if err != nil {
		return files, err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return files, err
	}

	json.Unmarshal(byteValue, &files)

	return files, nil
}

func parseConfig() (Config, error) {
	config := &Config{}
	flows := []string{}

	name := flag.String("name", "", "Integration name")
	flowFlag := flag.String("flow", "", "Integration flow types")
	filesPath := flag.String("files", PATH_TO_JSON, "Json files structure")

	flag.Parse()

	if *name == "" {
		return Config{}, errors.New("Error: Name required")
	}

	if *flowFlag != "" {
		flows = strings.Split(*flowFlag, "")
	} else {
		return Config{}, errors.New("Error: Flow types reuqired")
	}

	config.name = *name
	config.modifiers = flows
	config.jsonPath = *filesPath

	return *config, nil
}
