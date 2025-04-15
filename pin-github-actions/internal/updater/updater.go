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

// UpdateWorkflows scans for workflow files, parses them, and updates action references
func (u *Updater) UpdateWorkflows(ctx context.Context, fsys fs.FS) (int, error) {
	files, err := finder.FindWorkflowFiles(fsys)
	if err != nil {
		return 0, fmt.Errorf("failed to find workflow files: %w", err)
	}

	totalUpdates := 0
	for _, file := range files {
		updates, err := u.processWorkflowFile(ctx, fsys, file)
		if err != nil {
			return totalUpdates, err
		}
		totalUpdates += updates
	}
	return totalUpdates, nil
}

// processWorkflowFile reads a workflow file, parses it for action references
func (u *Updater) processWorkflowFile(ctx context.Context, fsys fs.FS, file string) (int, error) {
	log.Printf("Processing file: %s", file)

	content, err := fs.ReadFile(fsys, file)
	if err != nil {
		return 0, fmt.Errorf("failed to read file %s: %w", file, err)
	}

	actions, err := parser.ParseWorkflowActions(content)
	if err != nil {
		return 0, fmt.Errorf("failed to parse actions in file %s: %w", file, err)
	}

	debugActions(actions)
	log.Printf("Found %d actions in file %s", len(actions), file)

	updatedContent, fileUpdates, err := u.updateActionReferences(ctx, string(content), actions)
	if err != nil {
		return 0, err
	}

	if fileUpdates > 0 {
		return fileUpdates, u.writeUpdatedFile(fsys, file, updatedContent)
	}

	log.Printf("No changes made to file: %s", file)
	return 0, nil
}

// updateActionReferences updates action references in the content
func (u *Updater) updateActionReferences(ctx context.Context, content string, actions []types.ActionRef) (string, int, error) {
	updatedContent := content
	fileUpdates := 0

	for _, action := range actions {
		newContent, updated, err := u.updateSingleActionReference(ctx, updatedContent, action)
		if err != nil {
			return "", 0, err
		}
		if updated {
			updatedContent = newContent
			fileUpdates++
		}
	}

	return updatedContent, fileUpdates, nil
}

// updateSingleActionReference updates a single action reference in the content
func (u *Updater) updateSingleActionReference(ctx context.Context, content string, action types.ActionRef) (string, bool, error) {
	if isSHA(action.Ref) {
		log.Printf("Skipping %s/%s@%s (already a SHA)", action.Owner, action.Repo, action.Ref)
		return content, false, nil
	}

	log.Printf("Processing action: %s/%s@%s", action.Owner, action.Repo, action.Ref)
	sha, err := u.Client.ResolveActionSHA(ctx, action)
	if err != nil {
		return "", false, fmt.Errorf("failed to resolve SHA for action %s/%s@%s: %w",
			action.Owner, action.Repo, action.Ref, err)
	}
	log.Printf("Resolved SHA for %s/%s@%s: %s", action.Owner, action.Repo, action.Ref, sha)

	oldRef := fmt.Sprintf("%s/%s@%s", action.Owner, action.Repo, action.Ref)
	newRef := fmt.Sprintf("%s/%s@%s", action.Owner, action.Repo, sha)
	return strings.ReplaceAll(content, oldRef, newRef), true, nil
}

// writeUpdatedFile writes the updated content back to the file system
func (u *Updater) writeUpdatedFile(fsys fs.FS, file string, content string) error {
	if u.DryRun {
		log.Printf("[DRY RUN] Would update file: %s", file)
		return nil
	}

	if writeFS, ok := fsys.(interface {
		WriteFile(name string, data []byte, perm fs.FileMode) error
	}); ok {
		return writeFS.WriteFile(file, []byte(content), 0644)
	}

	absPath, err := filepath.Abs(file)
	if err != nil {
		return fmt.Errorf("failed to get absolute path for %s: %w", file, err)
	}
	return os.WriteFile(absPath, []byte(content), 0644)
}

// isSHA checks if a string is a valid Git SHA-1 hash (40 hex characters)
func isSHA(ref string) bool {
	return shaRegex.MatchString(ref)
}

// debugActions logs the action references for debugging purposes
func debugActions(actions []types.ActionRef) {
	for i, action := range actions {
		log.Printf("Action %d: %s/%s@%s", i+1, action.Owner, action.Repo, action.Ref)
	}
}
