package parser

import (
	"fmt"
	"github.com/zisuu/pin-github-actions/pgk/types"
	"regexp"
	"strings"
)

var (
	// Strict pattern for valid actions
	actionPattern = regexp.MustCompile(`(?m)^\s*-\s+uses:\s*([^\s@]+)@([^\s#]+)`)
	// Loose pattern to detect any uses line
	usesLinePattern = regexp.MustCompile(`(?m)^\s*-\s+uses:\s*([^\s#]+)`)
)

func ParseWorkflowActions(content []byte) ([]types.ActionRef, error) {
	var actions []types.ActionRef
	strContent := string(content)

	// First find all lines with uses:
	usesLines := usesLinePattern.FindAllStringSubmatch(strContent, -1)
	for _, line := range usesLines {
		if len(line) < 2 {
			continue
		}

		fullRef := strings.TrimSpace(line[1])

		// Skip local and docker actions
		if strings.HasPrefix(fullRef, ".") || strings.HasPrefix(fullRef, "docker://") {
			continue
		}

		// Check if it matches the strict pattern
		if !actionPattern.MatchString(line[0]) {
			return nil, fmt.Errorf("invalid action reference: %q", fullRef)
		}

		// Now extract parts using the strict pattern
		match := actionPattern.FindStringSubmatch(line[0])
		if len(match) != 3 {
			continue
		}

		fullRepo := match[1]
		ref := match[2]

		parts := strings.Split(fullRepo, "/")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid action reference: %s", fullRepo)
		}

		// Validate there are no empty parts
		if parts[0] == "" || parts[1] == "" || ref == "" {
			return nil, fmt.Errorf("invalid action reference: %s@%s", fullRepo, ref)
		}

		actions = append(actions, types.ActionRef{
			Owner: parts[0],
			Repo:  parts[1],
			Ref:   ref,
		})
	}

	return actions, nil
}
