# Copilot Instructions for gh-vibe

## Project Overview

gh-vibe is a GitHub CLI extension written in Go that provides commonly used commands for vibe coding workflows. The extension is built using the Cobra CLI framework and integrates with the GitHub CLI.

## Tech Stack

- **Language**: Go 1.25+
- **CLI Framework**: [Cobra](https://github.com/spf13/cobra) v1.10+
- **GitHub CLI Integration**: [go-gh](https://github.com/cli/go-gh) v2.13+

## Project Structure

```
gh-vibe/
├── main.go           # Entry point
├── cmd/              # Command implementations
│   ├── root.go       # Root command definition
│   ├── init.go       # Repository initialization command
│   ├── ready_merge.go # PR ready and merge command
│   └── cmd_test.go   # Tests for commands
├── go.mod            # Go module definition
├── go.sum            # Dependency checksums
└── .github/
    └── workflows/    # CI/CD workflows
```

## Commands

### Build

```bash
go build .
```

### Test

```bash
go test -v ./...
```

### Run Locally

```bash
go run main.go [command]
```

### Install as GitHub CLI Extension

```bash
gh extension install .
```

## Code Style Guidelines

- Follow standard Go conventions and idioms
- Use `gofmt` for formatting
- Keep functions focused and concise
- Use meaningful variable and function names
- Add comments for exported functions and types
- Handle errors explicitly, don't ignore them
- Use the Cobra command pattern for new CLI commands

## Adding New Commands

1. Create a new file in the `cmd/` directory (e.g., `cmd/mycommand.go`)
2. Define the command using `&cobra.Command{}`
3. Register the command in `init()` using `rootCmd.AddCommand()`
4. Add corresponding tests in `cmd/cmd_test.go`

Example command structure:

```go
package cmd

import "github.com/spf13/cobra"

var myCmd = &cobra.Command{
    Use:   "mycommand",
    Short: "Brief description",
    Long:  `Detailed description`,
    RunE:  runMyCommand,
}

func init() {
    rootCmd.AddCommand(myCmd)
}

func runMyCommand(cmd *cobra.Command, args []string) error {
    // Implementation
    return nil
}
```

## Git Workflow

- Create feature branches from `main`
- Use descriptive commit messages
- Squash merge PRs to keep history clean
- Branch protection is enabled on `main`

## Testing Guidelines

- Write unit tests for all new commands
- Test command registration with the root command
- Test argument validation
- Use table-driven tests where appropriate
- Run all tests before submitting changes: `go test -v ./...`

## Boundaries and Restrictions

- Do not modify the release workflow (`.github/workflows/release.yml`) without explicit approval
- Do not commit secrets, API keys, or sensitive credentials
- Do not add dependencies without justification
- Do not change the module path in `go.mod`
- Keep the extension compatible with the latest GitHub CLI
