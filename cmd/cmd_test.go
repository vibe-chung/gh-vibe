package cmd

import (
	"testing"
)

func TestRootCommandExists(t *testing.T) {
	if rootCmd == nil {
		t.Error("rootCmd should not be nil")
	}
}

func TestRootCommandUse(t *testing.T) {
	if rootCmd.Use != "gh-vibe" {
		t.Errorf("Expected rootCmd.Use to be 'gh-vibe', got '%s'", rootCmd.Use)
	}
}

func TestInitCommandExists(t *testing.T) {
	if initCmd == nil {
		t.Error("initCmd should not be nil")
	}
}

func TestInitCommandUse(t *testing.T) {
	if initCmd.Use != "init" {
		t.Errorf("Expected initCmd.Use to be 'init', got '%s'", initCmd.Use)
	}
}

func TestInitCommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "init" {
			found = true
			break
		}
	}
	if !found {
		t.Error("init command should be registered with root command")
	}
}
