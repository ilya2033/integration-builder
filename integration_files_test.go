package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasModifier_True(t *testing.T) {
	file := &File{}
	file.Modifiers = []string{"w", "m"}

	assert.True(t, file.HasModifier("m"))
}

func TestHasModifier_False(t *testing.T) {
	file := &File{}
	file.Modifiers = []string{"w", "h"}

	assert.False(t, file.HasModifier("m"))
}

func TestHasOneOfModifiers_True(t *testing.T) {
	file := &File{}
	file.Modifiers = []string{"w", "m"}

	result := file.hasOneOfModifiers([]string{"w", "h"})
	assert.True(t, result)
}

func TestHasOneOfModifiers_False(t *testing.T) {
	file := &File{}
	file.Modifiers = []string{"w", "h"}

	result := file.hasOneOfModifiers([]string{"c", "m"})
	assert.False(t, result)
}

func TestFilterFiles(t *testing.T) {
	testFiles := &Files{
		Files: []File{
			{Name: "file1", Modifiers: []string{"m", "h"}},
			{Name: "file2", Modifiers: []string{"m", "w"}},
			{Name: "file3", Modifiers: []string{"h", "w"}},
			{Name: "file4", Modifiers: []string{"h", "c"}},
			{Name: "file5", Modifiers: []string{}},
			{Name: "file6", Modifiers: []string{"c"}},
		},
	}

	config := &Config{
		Modifiers: []string{"m", "w"},
	}

	want := &Files{
		Files: []File{
			{Name: "file1", Modifiers: []string{"m", "h"}},
			{Name: "file2", Modifiers: []string{"m", "w"}},
			{Name: "file3", Modifiers: []string{"h", "w"}},
		},
	}

	result := filterFiles(*testFiles, *config)

	assert.Equal(t, *want, result)
}
