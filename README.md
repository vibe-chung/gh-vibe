# gh-vibe

A GitHub CLI extension that contains a collection of commonly used commands for vibe coding workflows.

## Installation

```bash
gh extension install vibe-chung/gh-vibe
```

## Commands

### `gh vibe init`

Initialize repository with recommended settings:
- Sets branch protection for the specified branch (enforce admins, no force pushes, no deletions)
- Configures merge settings (squash merge only, allow update branch)

```bash
gh vibe init                  # Protect main branch (default)
gh vibe init --branch develop # Protect a different branch
```

### `gh vibe setup-ai`

Create a GitHub issue to configure Copilot instructions for the repository. The issue is automatically assigned to @copilot.

```bash
gh vibe setup-ai
```

### `gh vibe ready-merge`

Mark a pull request ready for review and merge it using squash merge. After merging, checks out the target branch and pulls the latest changes.

```bash
gh vibe ready-merge      # Use current branch's PR
gh vibe ready-merge 123  # Specify a PR number
```

### `gh vibe usage`

Show Copilot premium request usage summary for the authenticated user.

```bash
gh vibe usage            # Show full JSON response
gh vibe usage --summary  # Show only total gross quantity
```

## Development

### Build

```bash
go build .
```

### Test

```bash
go test -v ./...
```

### Install locally

```bash
gh extension install .
```
