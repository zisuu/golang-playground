package parser

import (
	"fmt"
	"github.com/zisuu/github-actions-digest-pinner/pgk/types"
	"strings"

	"gopkg.in/yaml.v3"
)

// ParseWorkflowActions parses a GitHub Actions workflow file and extracts action references.
func ParseWorkflowActions(content []byte) ([]types.ActionRef, error) {
	var workflow struct {
		Jobs map[string]struct {
			Steps []struct {
				Uses string `yaml:"uses"`
			} `yaml:"steps"`
		} `yaml:"jobs"`
	}

	if err := yaml.Unmarshal(content, &workflow); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	var actions []types.ActionRef
	for _, job := range workflow.Jobs {
		for _, step := range job.Steps {
			if step.Uses == "" {
				continue
			}

			if strings.HasPrefix(step.Uses, "./") ||
				strings.HasPrefix(step.Uses, "../") ||
				strings.HasPrefix(step.Uses, "docker://") {
				continue
			}

			action, err := parseActionString(step.Uses)
			if err != nil {
				return nil, fmt.Errorf("invalid action reference %q: %w", step.Uses, err)
			}
			actions = append(actions, *action)
		}
	}

	return actions, nil
}

// parseActionString parses a string in the format "owner/repo@ref" into an ActionRef struct.
func parseActionString(actionStr string) (*types.ActionRef, error) {
	// Split into repo@ref parts
	parts := strings.Split(actionStr, "@")
	if len(parts) != 2 {
		return nil, fmt.Errorf("missing @ symbol in action reference")
	}

	fullRepo := parts[0]
	ref := parts[1]

	// Split owner/repo
	repoParts := strings.Split(fullRepo, "/")
	if len(repoParts) != 2 {
		return nil, fmt.Errorf("invalid repository format, expected owner/repo")
	}

	owner := repoParts[0]
	repo := repoParts[1]

	if owner == "" || repo == "" || ref == "" {
		return nil, fmt.Errorf("empty component in action reference")
	}

	return &types.ActionRef{
		Owner: owner,
		Repo:  repo,
		Ref:   ref,
	}, nil
}
