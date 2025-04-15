package ghclient

import (
	"context"
	"fmt"
	"github.com/google/go-github/v71/github"
	"github.com/zisuu/github-actions-digest-pinner/pgk/types"
	"net/http"
)

type GitHubClient interface {
	ResolveActionSHA(context.Context, types.ActionRef) (string, error)
}

type DefaultGitHubClient struct {
	*github.Client
}

// NewDefaultClient creates a new GitHub client with the default HTTP client.
func NewDefaultClient() *DefaultGitHubClient {
	return &DefaultGitHubClient{
		Client: github.NewClient(http.DefaultClient),
	}
}

// ResolveActionSHA resolves the SHA of a GitHub Action reference.
func (c *DefaultGitHubClient) ResolveActionSHA(ctx context.Context, action types.ActionRef) (string, error) {
	commit, _, err := c.Repositories.GetCommitSHA1(ctx, action.Owner, action.Repo, action.Ref, "")
	if err != nil {
		return "", fmt.Errorf("failed to resolve SHA for %s/%s@%s: %w",
			action.Owner, action.Repo, action.Ref, err)
	}
	return commit, nil
}
