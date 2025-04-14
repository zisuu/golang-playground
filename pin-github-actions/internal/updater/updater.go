package updater

import (
	"context"
	"fmt"
	"github.com/zisuu/pin-github-actions/internal/finder"
	"github.com/zisuu/pin-github-actions/internal/ghclient"
	"github.com/zisuu/pin-github-actions/internal/parser"
	"github.com/zisuu/pin-github-actions/pgk/types"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var shaRegex = regexp.MustCompile(`^[0-9a-fA-F]{40}$`)

type Updater struct {
	Client ghclient.GitHubClient
	DryRun bool // Add dry-run mode for testing/safety
}

func NewUpdater(client ghclient.GitHubClient, dryRun bool) *Updater {
	return &Updater{Client: client, DryRun: dryRun}
}

func (u *Updater) UpdateWorkflows(ctx context.Context, fsys fs.FS) (int, error) {
	files, err := finder.FindWorkflowFiles(fsys)
	if err != nil {
		return 0, fmt.Errorf("failed to find workflow files: %w", err)
	}

	totalUpdates := 0

	for _, file := range files {
		log.Printf("Processing file: %s", file)
		content, err := fs.ReadFile(fsys, file)
		if err != nil {
			return totalUpdates, fmt.Errorf("failed to read file %s: %w", file, err)
		}

		actions, err := parser.ParseWorkflowActions(content)
		if err != nil {
			return totalUpdates, fmt.Errorf("failed to parse actions in file %s: %w", file, err)
		}

		debugActions(actions)
		log.Printf("Found %d actions in file %s", len(actions), file)
		originalContent := string(content)
		updatedContent := originalContent
		fileUpdates := 0

		for _, action := range actions {
			if isSHA(action.Ref) {
				log.Printf("Skipping %s/%s@%s (already a SHA)", action.Owner, action.Repo, action.Ref)
				continue
			}

			log.Printf("Processing action: %s/%s@%s", action.Owner, action.Repo, action.Ref)
			sha, err := u.Client.ResolveActionSHA(ctx, action)
			if err != nil {
				return totalUpdates, fmt.Errorf("failed to resolve SHA for action %s/%s@%s: %w",
					action.Owner, action.Repo, action.Ref, err)
			}
			log.Printf("Resolved SHA for %s/%s@%s: %s", action.Owner, action.Repo, action.Ref, sha)

			oldRef := fmt.Sprintf("%s/%s@%s", action.Owner, action.Repo, action.Ref)
			newRef := fmt.Sprintf("%s/%s@%s", action.Owner, action.Repo, sha)
			updatedContent = strings.ReplaceAll(updatedContent, oldRef, newRef)
			fileUpdates++
		}

		if fileUpdates > 0 { // Only proceed if we actually made changes
			totalUpdates += fileUpdates
			if u.DryRun {
				log.Printf("[DRY RUN] Would update file: %s (%d changes)", file, fileUpdates)
				continue
			}

			if writeFS, ok := fsys.(interface {
				WriteFile(name string, data []byte, perm fs.FileMode) error
			}); ok {
				err = writeFS.WriteFile(file, []byte(updatedContent), 0644)
				if err != nil {
					return totalUpdates, fmt.Errorf("failed to write updated file %s: %w", file, err)
				}
			} else {
				absPath, err := filepath.Abs(file)
				if err != nil {
					return totalUpdates, fmt.Errorf("failed to get absolute path for %s: %w", file, err)
				}
				err = os.WriteFile(absPath, []byte(updatedContent), 0644)
				if err != nil {
					return totalUpdates, fmt.Errorf("failed to write updated file %s: %w", absPath, err)
				}
			}
			log.Printf("Updated file: %s (%d changes)", file, fileUpdates)
		} else {
			log.Printf("No changes made to file: %s", file)
		}
	}
	return totalUpdates, nil
}

// isSHA checks if a string is a valid Git SHA-1 hash (40 hex characters)
func isSHA(ref string) bool {
	return shaRegex.MatchString(ref)
}

func debugActions(actions []types.ActionRef) {
	for i, action := range actions {
		log.Printf("Action %d: %s/%s@%s", i+1, action.Owner, action.Repo, action.Ref)
	}
}
