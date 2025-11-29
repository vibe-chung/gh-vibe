package cmd

import (
	"fmt"

	gh "github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"
)

var setupAICmd = &cobra.Command{
	Use:   "setup-ai",
	Short: "Create an issue to set up Copilot instructions",
	Long: `Create a GitHub issue to configure Copilot instructions for this repository.

The issue is created with:
- Title: "✨ Set up Copilot instructions"
- Body: Instructions for configuring Copilot coding agent
- Assignee: @copilot

This follows the best practices documented at https://gh.io/copilot-coding-agent-tips`,
	Args: cobra.NoArgs,
	RunE: runSetupAI,
}

func init() {
	rootCmd.AddCommand(setupAICmd)
}

func runSetupAI(_ *cobra.Command, _ []string) error {
	fmt.Println("Creating issue to set up Copilot instructions...")

	title := "✨ Set up Copilot instructions"
	body := `Configure instructions for this repository as documented in [Best practices for Copilot coding agent in your repository](https://gh.io/copilot-coding-agent-tips).

<Onboard this repo>`

	args := []string{
		"issue", "create",
		"--title", title,
		"--body", body,
		"--assignee", "@copilot",
	}

	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		return fmt.Errorf("%s: %w", stderr.String(), err)
	}
	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}

	fmt.Println("Issue created successfully!")
	return nil
}
