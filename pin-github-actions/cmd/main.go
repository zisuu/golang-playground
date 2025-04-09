package main

import (
	"fmt"
	"github.com/zisuu/pin-github-actions/internal/finder"
	"github.com/zisuu/pin-github-actions/internal/parser"
	"io/fs"
	"os"
)

func main() {
	// Example usage:
	fsys := os.DirFS(".")
	files, err := finder.FindWorkflowFiles(fsys)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error finding workflow files: %v\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		content, err := fs.ReadFile(fsys, file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", file, err)
			continue
		}

		actions, err := parser.ParseWorkflowActions(content)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", file, err)
			continue
		}

		fmt.Printf("Found %d actions in %s\n", len(actions), file)
	}
}
