package main

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func saveFile(file File, content []byte, config Config) {
	path := config.TargetPath + file.TargetPath
	fileExists := checkIfFileExists(path)

	if fileExists {
		log(path + " already exists. Skip")
		return
	}

	createDirAllPath(path)

	newFile := createNewFile(path)
	defer newFile.Close()

	writeToFile(newFile, content)

	log("File saved - " + file.Name)
}

func checkIfFileExists(path string) bool {
	_, err := os.Stat(path)

	if err == nil {
		return true
	}

	return false
}

func createNewFile(path string) *os.File {
	newFile, err := os.Create(path)
	check(err)
	return newFile
}

func writeToFile(file *os.File, content []byte) {
	file.Write(content)
	file.Sync()
}

func createDirAllPath(path string) {
	err := os.MkdirAll(filepath.Dir(path), fs.ModePerm)
	check(err)
}

func readFromJson(path string) []byte {
	jsonFile, err := os.Open(path)
	check(err)

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	check(err)

	return byteValue
}
