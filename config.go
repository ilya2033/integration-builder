package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/joho/godotenv"
)

const PATH_TO_JSON = "files.json"
const PATH_TO_TEMPLATE = "template/"

type Config struct {
	Name         string
	Modifiers    []string
	JsonPath     string
	TemplatePath string
	TargetPath   string
	WithTests    bool
}

func (this *Config) CheckAllRequiredFilled() {
	if this.Name == "" {
		panic(errors.New("Error: Name required"))
	}

	if this.TargetPath == "" {
		panic(errors.New("Error: Target required"))
	}

	if len(this.Modifiers) == 0 {
		panic(errors.New("Error: Flow types reuqired"))
	}
}

func parseConfig() Config {
	config := &Config{}
	envFile, err := godotenv.Read(".env")

	assignFromFlags(config)

	if err == nil {
		assignFromEnvFile(envFile, config)
	}

	assignFromDefaults(config)

	config.CheckAllRequiredFilled()

	return *config
}

func assignFromEnvFile(envFile map[string]string, config *Config) {
	if config.Name == "" {
		config.Name = envFile["NAME"]
	}

	if len(config.Modifiers) == 0 {
		config.Modifiers = strings.Split(envFile["FLOWS"], "")
	}

	if config.JsonPath == "" {
		config.JsonPath = envFile["JSON_PATH"]
	}

	if config.TemplatePath == "" {
		config.TemplatePath = envFile["TEMPLATE_PATH"]
	}

	if config.TargetPath == "" {
		config.TargetPath = envFile["TARGET_PATH"]
	}
}

func assignFromDefaults(config *Config) {
	if config.JsonPath == "" {
		config.JsonPath = PATH_TO_JSON
	}

	if config.TemplatePath == "" {
		config.TemplatePath = PATH_TO_TEMPLATE
	}
}

func assignFromFlags(config *Config) {
	flows := []string{}

	name := flag.String("name", "", "Integration name")
	flowFlag := flag.String("flow", "", "Integration flow types")
	filesPath := flag.String("files", "", "Json files structure")
	templatePath := flag.String("template", "", "Template folder")
	targetPath := flag.String("target", "", "Target folder")
	withTests := flag.Bool("tests", false, "Should generate tests files")

	flag.Parse()

	flows = strings.Split(*flowFlag, "")

	config.Name = *name
	config.Modifiers = flows
	config.JsonPath = *filesPath
	config.TemplatePath = *templatePath
	config.TargetPath = *targetPath
	config.WithTests = *withTests
}
