package main

import (
	"errors"
	"flag"
	"strings"
)

type Config struct {
	Name         string
	Modifiers    []string
	JsonPath     string
	TemplatePath string
	TargetPath   string
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

func (this *File) HasModifier(modifier string) bool {
	for _, fileMod := range this.Modifiers {
		if fileMod == modifier {
			return true
		}
	}

	return false
}

func (this *File) hasOnOfModifiers(mods []string) bool {
	for _, fileMod := range this.Modifiers {
		for _, modifier := range mods {
			if fileMod == modifier {
				return true
			}
		}
	}

	return false
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

	flows = strings.Split(*flowFlag, "")

	config.Name = *name
	config.Modifiers = flows
	config.JsonPath = *filesPath
	config.TemplatePath = *templatePath
	config.TargetPath = *targetPath
	config.CheckAllRequiredFilled()

	return *config
}
