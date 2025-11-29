package cmd

import (
	"fmt"

	"github.com/cli/go-gh/v2"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize repository settings and branch protection",
	Long: `Initialize repository with recommended settings:
- Sets branch protection for the main branch (enforce admins, no force pushes, no deletions)
- Configures merge settings (squash merge only, allow update branch)`,
	RunE: runInit,
}

func runInit(cmd *cobra.Command, args []string) error {
	repo, err := repository.Current()
	if err != nil {
		return fmt.Errorf("failed to get current repository: %w", err)
	}

	owner := repo.Owner
	repoName := repo.Name

	fmt.Printf("Initializing repository %s/%s...\n", owner, repoName)

	// Set branch protection for main branch
	if err := setBranchProtection(owner, repoName); err != nil {
		return fmt.Errorf("failed to set branch protection: %w", err)
	}

	// Update repository settings
	if err := updateRepoSettings(owner, repoName); err != nil {
		return fmt.Errorf("failed to update repository settings: %w", err)
	}

	fmt.Println("Repository initialized successfully!")
	return nil
}

func setBranchProtection(owner, repo string) error {
	fmt.Println("Setting branch protection for main branch...")

	// Equivalent to:
	// gh api --method PUT -H "Accept: application/vnd.github+json" \
	//   /repos/{owner}/{repo}/branches/main/protection \
	//   -f required_status_checks=null -f enforce_admins=true \
	//   -f required_pull_request_reviews=null -f restrictions=null \
	//   -f allow_force_pushes=false -f allow_deletions=false
	_, _, err := gh.Exec(
		"api",
		"--method", "PUT",
		"-H", "Accept: application/vnd.github+json",
		fmt.Sprintf("/repos/%s/%s/branches/main/protection", owner, repo),
		"-f", "required_status_checks=null",
		"-F", "enforce_admins=true",
		"-f", "required_pull_request_reviews=null",
		"-f", "restrictions=null",
		"-F", "allow_force_pushes=false",
		"-F", "allow_deletions=false",
	)
	if err != nil {
		return err
	}

	fmt.Println("Branch protection configured successfully!")
	return nil
}

func updateRepoSettings(owner, repo string) error {
	fmt.Println("Updating repository settings...")

	// Equivalent to:
	// gh api --method PATCH -H "Accept: application/vnd.github+json" \
	//   /repos/{owner}/{repo} \
	//   -F allow_squash_merge=true -F allow_merge_commit=false \
	//   -F allow_rebase_merge=false -F allow_update_branch=true
	_, _, err := gh.Exec(
		"api",
		"--method", "PATCH",
		"-H", "Accept: application/vnd.github+json",
		fmt.Sprintf("/repos/%s/%s", owner, repo),
		"-F", "allow_squash_merge=true",
		"-F", "allow_merge_commit=false",
		"-F", "allow_rebase_merge=false",
		"-F", "allow_update_branch=true",
	)
	if err != nil {
		return err
	}

	fmt.Println("Repository settings updated successfully!")
	return nil
}
