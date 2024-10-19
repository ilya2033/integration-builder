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
