package main

type File struct {
	Name         string   `json:"name"`
	Modifiers    []string `json:"modifiers"`
	TemplatePath string   `json:"templatePath"`
	TargetPath   string   `json:"targetPath"`
}

type Files struct {
	Files []File `json:"files"`
}

type Config struct {
	name      string
	modifiers []string
	jsonPath  string
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
