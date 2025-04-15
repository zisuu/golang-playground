package ghclient

import (
	"context"
	"fmt"
	"github.com/google/go-github/v71/github"
	"github.com/zisuu/github-actions-digest-pinner/pgk/types"
	"golang.org/x/oauth2"
	"os"
	"strings"
)

type GitHubClient interface {
	ResolveActionSHA(ctx context.Context, action types.ActionRef) (string, error)
}

type githubClient struct {
	client *github.Client
}

// NewGitHubClient creates a new GitHub client.
func NewGitHubClient() GitHubClient {
	// Get token from environment
	token := os.Getenv("GITHUB_TOKEN")

	var client *github.Client
	if token != "" {
		// Authenticated client
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client = github.NewClient(tc)
	} else {
		// Unauthenticated client
		client = github.NewClient(nil)
	}

	return &githubClient{client: client}
}

// ResolveActionSHA resolves the SHA of a GitHub Action reference.
func (g *githubClient) ResolveActionSHA(ctx context.Context, action types.ActionRef) (string, error) {
	if isSHA(action.Ref) {
		return action.Ref, nil
	}

	// For actions with paths, we need to get the repo's default branch first
	var ref *github.Reference
	var err error

	// First try as a tag
	ref, _, err = g.client.Git.GetRef(ctx, action.Owner, action.Repo, fmt.Sprintf("refs/tags/%s", action.Ref))
	if err != nil {
		// Try as a branch if tag not found
		ref, _, err = g.client.Git.GetRef(ctx, action.Owner, action.Repo, fmt.Sprintf("refs/heads/%s", action.Ref))
		if err != nil {
			// If we still can't find it, try getting the default branch
			repo, _, err := g.client.Repositories.Get(ctx, action.Owner, action.Repo)
			if err != nil {
				return "", fmt.Errorf("failed to resolve ref %s: %w", action.Ref, err)
			}
			defaultBranch := repo.GetDefaultBranch()
			if defaultBranch == "" {
				defaultBranch = "main" // fallback
			}
			ref, _, err = g.client.Git.GetRef(ctx, action.Owner, action.Repo,
				fmt.Sprintf("refs/heads/%s", defaultBranch))
			if err != nil {
				return "", fmt.Errorf("failed to resolve ref %s: %w", action.Ref, err)
			}
		}
	}

	sha := strings.TrimPrefix(ref.GetRef(), "refs/heads/")
	sha = strings.TrimPrefix(sha, "refs/tags/")
	return sha, nil
}

func isSHA(ref string) bool {
	// Your existing SHA validation logic
	return len(ref) == 40 && isHexString(ref)
}

func isHexString(s string) bool {
	for _, c := range s {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')) {
			return false
		}
	}
	return true
}
