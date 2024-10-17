package main

const PATH_TO_JSON = "files.json"
const PATH_TO_TEMPLATE = "template/"

func main() {
	config := parseConfig()
	files := parseFilesFromJson(config)
	files = filterFiles(files, config)
	placeFiles(files, config)
}
