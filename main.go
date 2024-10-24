package main

func main() {
	config := parseConfig()
	allFiles := parseFilesFromJson(config)
	allFiles.Files = filterFiles(allFiles.Files, config)
	placeFiles(allFiles.Files, config)

	if config.WithTests {
		allFiles.TestFiles = filterFiles(allFiles.TestFiles, config)
		placeFiles(allFiles.TestFiles, config)
	}
}
