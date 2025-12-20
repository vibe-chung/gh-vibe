package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spf13/cobra"
)

var branchFlag string

var initCmd = &cobra.Command{
	Use:   "init [owner/repo]",
	Short: "Initialize repository settings and branch protection",
	Long: `Initialize repository with recommended settings:
- Sets branch protection for the specified branch (enforce admins, no force pushes, no deletions)
- Configures merge settings (squash merge only, allow update branch)

If owner/repo is not provided, uses the current repository.`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInit,
}

func init() {
	initCmd.Flags().StringVarP(&branchFlag, "branch", "b", "main", "Branch name to protect")
}

func runInit(cmd *cobra.Command, args []string) error {
	var owner, repoName string

	if len(args) > 0 {
		// Parse owner/repo from argument
		parts := strings.Split(args[0], "/")
		if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
			return fmt.Errorf("invalid repository format: expected 'owner/repo', got '%s'", args[0])
		}
		owner = parts[0]
		repoName = parts[1]
	} else {
		// Use current repository
		repo, err := repository.Current()
		if err != nil {
			return fmt.Errorf("failed to get current repository: %w", err)
		}
		owner = repo.Owner
		repoName = repo.Name
	}

	fmt.Printf("Initializing repository %s/%s...\n", owner, repoName)

	client, err := api.DefaultRESTClient()
	if err != nil {
		return fmt.Errorf("failed to create REST client: %w", err)
	}

	// Set branch protection for the specified branch
	if err := setBranchProtection(client, owner, repoName, branchFlag); err != nil {
		return fmt.Errorf("failed to set branch protection: %w", err)
	}

	// Update repository settings
	if err := updateRepoSettings(client, owner, repoName); err != nil {
		return fmt.Errorf("failed to update repository settings: %w", err)
	}

	fmt.Println("Repository initialized successfully!")
	return nil
}

func setBranchProtection(client *api.RESTClient, owner, repo, branch string) error {
	fmt.Printf("Setting branch protection for %s branch...\n", branch)

	// Branch protection payload
	payload := map[string]interface{}{
		"required_status_checks":        nil,
		"enforce_admins":                true,
		"required_pull_request_reviews": nil,
		"restrictions":                  nil,
		"allow_force_pushes":            false,
		"allow_deletions":               false,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	endpoint := fmt.Sprintf("repos/%s/%s/branches/%s/protection", owner, repo, branch)
	err = client.Put(endpoint, bytes.NewReader(body), nil)
	if err != nil {
		return err
	}

	fmt.Println("Branch protection configured successfully!")
	return nil
}

func updateRepoSettings(client *api.RESTClient, owner, repo string) error {
	fmt.Println("Updating repository settings...")

	// Repository settings payload
	payload := map[string]interface{}{
		"allow_squash_merge":     true,
		"allow_merge_commit":     false,
		"allow_rebase_merge":     false,
		"allow_update_branch":    true,
		"delete_branch_on_merge": true,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	endpoint := fmt.Sprintf("repos/%s/%s", owner, repo)
	err = client.Patch(endpoint, bytes.NewReader(body), nil)
	if err != nil {
		return err
	}

	fmt.Println("Repository settings updated successfully!")
	return nil
}
