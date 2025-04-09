package parser

import (
	"github.com/zisuu/pin-github-actions/pgk/types"
	"testing"
)

func TestParseWorkflowActions(t *testing.T) {
	tests := []struct {
		name     string
		content  string
		expected []types.ActionRef
		wantErr  bool
	}{
		{
			name: "simple action reference",
			content: `
name: CI
on: push
jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
      - uses: actions/setup-java@v4.7
      - uses: actions/setup-node@v4.3.0
`,
			expected: []types.ActionRef{
				{Owner: "actions", Repo: "checkout", Ref: "v3"},
				{Owner: "actions", Repo: "setup-go", Ref: "v4"},
				{Owner: "actions", Repo: "setup-java", Ref: "v4.7"},
				{Owner: "actions", Repo: "setup-node", Ref: "v4.3.0"},
			},
		},
		{
			name: "with whitespace",
			content: `
steps:
  - uses:  actions/checkout  @v3
`,
			wantErr: true,
		},
		{
			name: "invalid action reference",
			content: `
steps:
  - uses: invalid-ref
`,
			wantErr: true,
		},
		{
			name: "missing version",
			content: `
steps:
  - uses: actions/checkout@
`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actions, err := ParseWorkflowActions([]byte(tt.content))

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if len(actions) != len(tt.expected) {
				t.Fatalf("expected %d actions, got %d", len(tt.expected), len(actions))
			}

			for i, action := range actions {
				if action != tt.expected[i] {
					t.Errorf("action %d mismatch:\nexpected: %+v\ngot:      %+v",
						i, tt.expected[i], action)
				}
			}
		})
	}
}
