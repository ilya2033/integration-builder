package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAllRequiredFilled_False(t *testing.T) {
	config := &Config{}

	assert.Panics(t, config.CheckAllRequiredFilled, "Blank config doesn't panic")

	config.Name = "Name"
	config.Modifiers = []string{"h", "m"}

	assert.Panics(t, config.CheckAllRequiredFilled, "Half filled config doesn't panic")
}

func TestCheckAllRequiredFilled_True(t *testing.T) {
	config := &Config{}

	config.Name = "Name"
	config.Modifiers = []string{"h", "m"}
	config.TargetPath = "target"

	assert.NotPanics(t, config.CheckAllRequiredFilled, "Full config panics")
}
