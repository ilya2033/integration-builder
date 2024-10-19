package main

func main() {
	config := parseConfig()
	files := parseFilesFromJson(config)
	files = filterFiles(files, config)
	placeFiles(files, config)
}
