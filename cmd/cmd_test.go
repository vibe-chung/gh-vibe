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

func TestInitCommandRejectsArgs(t *testing.T) {
	if initCmd.Args == nil {
		t.Error("initCmd.Args should not be nil")
		return
	}

	err := initCmd.Args(initCmd, []string{"unexpected-arg"})
	if err == nil {
		t.Error("initCmd should reject unexpected arguments")
	}
}

func TestInitCommandAcceptsNoArgs(t *testing.T) {
	if initCmd.Args == nil {
		t.Error("initCmd.Args should not be nil")
		return
	}

	err := initCmd.Args(initCmd, []string{})
	if err != nil {
		t.Errorf("initCmd should accept no arguments, got error: %v", err)
	}
}

func TestInitCommandHasBranchFlag(t *testing.T) {
	flag := initCmd.Flags().Lookup("branch")
	if flag == nil {
		t.Error("initCmd should have a 'branch' flag")
		return
	}
	if flag.DefValue != "main" {
		t.Errorf("Expected branch flag default to be 'main', got '%s'", flag.DefValue)
	}
	if flag.Shorthand != "b" {
		t.Errorf("Expected branch flag shorthand to be 'b', got '%s'", flag.Shorthand)
	}
}
