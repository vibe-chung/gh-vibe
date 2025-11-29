package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gh-vibe",
	Short: "A GitHub CLI extension with commonly used commands for vibe coding",
	Long: `gh-vibe is a GitHub CLI extension that contains a collection of
commonly used gh commands for vibe coding workflows.

It includes commands for:
- Setting up branch protection and repository settings
- Creating issues for AI instructions
- Managing pull requests`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(initCmd)
}
