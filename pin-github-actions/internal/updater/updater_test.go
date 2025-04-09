package updater

import (
	"testing"
)

func TestUpdateActionReferences(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		replacements map[string]string
		expected     string
		wantErr      bool
	}{
		{
			name: "basic replacement",
			input: `
steps:
  - uses: actions/checkout@v3
  - run: echo "Hello"
`,
			replacements: map[string]string{
				"actions/checkout@v3": "actions/checkout@a81bbbf",
			},
			expected: `
steps:
  - uses: actions/checkout@a81bbbf
  - run: echo "Hello"
`,
		},
		{
			name: "preserves formatting",
			input: `
  - uses:  actions/checkout@v3  # with comment
`,
			replacements: map[string]string{
				"actions/checkout@v3": "actions/checkout@a81bbbf",
			},
			expected: `
  - uses: actions/checkout@a81bbbf  # with comment
`,
		},
		{
			name: "multiple replacements",
			input: `
steps:
  - uses: actions/checkout@v3
  - uses: actions/setup-go@v4
`,
			replacements: map[string]string{
				"actions/checkout@v3": "actions/checkout@a81bbbf",
				"actions/setup-go@v4": "actions/setup-go@1b3a3f9",
			},
			expected: `
steps:
  - uses: actions/checkout@a81bbbf
  - uses: actions/setup-go@1b3a3f9
`,
		},
		{
			name: "no replacements found",
			input: `
steps:
  - run: echo "No actions here"
`,
			replacements: map[string]string{
				"actions/checkout@v3": "actions/checkout@a81bbbf",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updated, err := UpdateActionReferences([]byte(tt.input), tt.replacements)

			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(updated) != tt.expected {
				t.Errorf("content mismatch:\nexpected: %q\ngot:      %q",
					tt.expected, string(updated))
			}
		})
	}
}

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		changed  string
	}{
		{
			name:     "basic replacement",
			input:    "  - uses: actions/checkout@v3",
			expected: "  - uses: actions/checkout@a81bbbf",
			changed:  "actions/checkout@v3 -> actions/checkout@a81bbbf",
		},
		{
			name:     "preserves whitespace",
			input:    "    - uses:  actions/checkout@v3  ",
			expected: "    - uses: actions/checkout@a81bbbf  ",
			changed:  "actions/checkout@v3 -> actions/checkout@a81bbbf",
		},
		{
			name:     "ignores non-matching lines",
			input:    "  - run: echo 'hello'",
			expected: "  - run: echo 'hello'",
			changed:  "",
		},
	}

	replacements := map[string]string{
		"actions/checkout@v3": "actions/checkout@a81bbbf",
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updated, changed := processLine([]byte(tt.input), replacements)

			if string(updated) != tt.expected {
				t.Errorf("line mismatch:\nexpected: %q\ngot:      %q",
					tt.expected, string(updated))
			}

			if changed != tt.changed {
				t.Errorf("change message mismatch:\nexpected: %q\ngot:      %q",
					tt.changed, changed)
			}
		})
	}
}
