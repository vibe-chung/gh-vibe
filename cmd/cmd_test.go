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

func TestReadyMergeCommandExists(t *testing.T) {
	if readyMergeCmd == nil {
		t.Error("readyMergeCmd should not be nil")
	}
}

func TestReadyMergeCommandUse(t *testing.T) {
	if readyMergeCmd.Use != "ready-merge [PR number]" {
		t.Errorf("Expected readyMergeCmd.Use to be 'ready-merge [PR number]', got '%s'", readyMergeCmd.Use)
	}
}

func TestReadyMergeCommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "ready-merge [PR number]" {
			found = true
			break
		}
	}
	if !found {
		t.Error("ready-merge command should be registered with root command")
	}
}

func TestReadyMergeCommandAcceptsNoArgs(t *testing.T) {
	if readyMergeCmd.Args == nil {
		t.Error("readyMergeCmd.Args should not be nil")
		return
	}

	err := readyMergeCmd.Args(readyMergeCmd, []string{})
	if err != nil {
		t.Errorf("readyMergeCmd should accept no arguments, got error: %v", err)
	}
}

func TestReadyMergeCommandAcceptsOneArg(t *testing.T) {
	if readyMergeCmd.Args == nil {
		t.Error("readyMergeCmd.Args should not be nil")
		return
	}

	err := readyMergeCmd.Args(readyMergeCmd, []string{"123"})
	if err != nil {
		t.Errorf("readyMergeCmd should accept one argument, got error: %v", err)
	}
}

func TestReadyMergeCommandRejectsTwoArgs(t *testing.T) {
	if readyMergeCmd.Args == nil {
		t.Error("readyMergeCmd.Args should not be nil")
		return
	}

	err := readyMergeCmd.Args(readyMergeCmd, []string{"123", "456"})
	if err == nil {
		t.Error("readyMergeCmd should reject more than one argument")
	}
}

func TestSetupAICommandExists(t *testing.T) {
	if setupAICmd == nil {
		t.Error("setupAICmd should not be nil")
	}
}

func TestSetupAICommandUse(t *testing.T) {
	if setupAICmd.Use != "setup-ai" {
		t.Errorf("Expected setupAICmd.Use to be 'setup-ai', got '%s'", setupAICmd.Use)
	}
}

func TestSetupAICommandRegistered(t *testing.T) {
	found := false
	for _, cmd := range rootCmd.Commands() {
		if cmd.Use == "setup-ai" {
			found = true
			break
		}
	}
	if !found {
		t.Error("setup-ai command should be registered with root command")
	}
}

func TestSetupAICommandRejectsArgs(t *testing.T) {
	if setupAICmd.Args == nil {
		t.Error("setupAICmd.Args should not be nil")
		return
	}

	err := setupAICmd.Args(setupAICmd, []string{"unexpected-arg"})
	if err == nil {
		t.Error("setupAICmd should reject unexpected arguments")
	}
}

func TestSetupAICommandAcceptsNoArgs(t *testing.T) {
	if setupAICmd.Args == nil {
		t.Error("setupAICmd.Args should not be nil")
		return
	}

	err := setupAICmd.Args(setupAICmd, []string{})
	if err != nil {
		t.Errorf("setupAICmd should accept no arguments, got error: %v", err)
	}
}
