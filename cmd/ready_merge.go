package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	gh "github.com/cli/go-gh/v2"
	"github.com/spf13/cobra"
)

var readyMergeCmd = &cobra.Command{
	Use:   "ready-merge [PR number]",
	Short: "Mark PR ready for review and merge it",
	Long: `Mark a pull request ready for review and merge it using squash merge.
After merging, the command will checkout the target branch and pull the latest changes.

If the current branch is a PR branch, no PR number is needed.
Otherwise, you must provide the PR number as an argument.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runReadyMerge,
}

func init() {
	rootCmd.AddCommand(readyMergeCmd)
}

func runReadyMerge(_ *cobra.Command, args []string) error {
	var prNumber string
	if len(args) > 0 {
		prNumber = args[0]
	}

	// Get the target branch before marking ready and merging
	targetBranch, err := getPRTargetBranch(prNumber)
	if err != nil {
		return fmt.Errorf("failed to get PR target branch: %w", err)
	}

	// Mark PR ready for review
	if err := markPRReady(prNumber); err != nil {
		return fmt.Errorf("failed to mark PR ready: %w", err)
	}

	// Merge PR with squash
	if err := mergePR(prNumber); err != nil {
		return fmt.Errorf("failed to merge PR: %w", err)
	}

	// Checkout target branch and pull latest changes
	if err := checkoutAndPull(targetBranch); err != nil {
		return fmt.Errorf("failed to checkout and pull: %w", err)
	}

	fmt.Println("PR merged successfully!")
	return nil
}

func markPRReady(prNumber string) error {
	fmt.Println("Marking PR ready for review...")
	args := []string{"pr", "ready"}
	if prNumber != "" {
		args = append(args, prNumber)
	}

	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		return fmt.Errorf("%s: %w", stderr.String(), err)
	}
	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}
	return nil
}

func mergePR(prNumber string) error {
	fmt.Println("Merging PR with squash...")
	args := []string{"pr", "merge", "--squash"}
	if prNumber != "" {
		args = append(args, prNumber)
	}

	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		return fmt.Errorf("%s: %w", stderr.String(), err)
	}
	if stdout.Len() > 0 {
		fmt.Print(stdout.String())
	}
	return nil
}

func getPRTargetBranch(prNumber string) (string, error) {
	args := []string{"pr", "view", "--json", "baseRefName", "--jq", ".baseRefName"}
	if prNumber != "" {
		args = append(args, prNumber)
	}

	stdout, stderr, err := gh.Exec(args...)
	if err != nil {
		return "", fmt.Errorf("%s: %w", stderr.String(), err)
	}
	return strings.TrimSpace(stdout.String()), nil
}

func checkoutAndPull(branch string) error {
	fmt.Printf("Checking out %s and pulling latest changes...\n", branch)

	// git checkout <branch>
	checkoutCmd := exec.Command("git", "checkout", branch)
	if output, err := checkoutCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git checkout failed: %s: %w", string(output), err)
	}

	// git pull
	pullCmd := exec.Command("git", "pull")
	if output, err := pullCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("git pull failed: %s: %w", string(output), err)
	}

	return nil
}
