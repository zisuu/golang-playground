package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zisuu/github-actions-digest-pinner/internal/finder"
	"github.com/zisuu/github-actions-digest-pinner/internal/ghclient"
	"github.com/zisuu/github-actions-digest-pinner/internal/updater"
)

func main() {
	// Command-line flags
	dryRun := flag.Bool("dry-run", false, "Preview changes without modifying files")
	dir := flag.String("dir", ".", "Directory containing GitHub workflows")
	timeout := flag.Int("timeout", 30, "API timeout in seconds")
	verbose := flag.Bool("v", false, "Verbose output")
	flag.Parse()

	// Setup
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeout)*time.Second)
	defer cancel()

	if *verbose {
		log.Println("Starting GitHub Actions pinning utility")
		log.Printf("Scanning directory: %s", *dir)
	}

	// 1. Initialize all components
	client := ghclient.NewDefaultClient()
	newupdater := updater.NewUpdater(client, *dryRun)
	fsys := os.DirFS(*dir)

	// 2. Find workflow files (finder module)
	if *verbose {
		log.Println("Finding workflow files...")
	}
	files, err := finder.FindWorkflowFiles(fsys)
	if err != nil {
		log.Fatalf("Failed to find workflow files: %v", err)
	}

	if *verbose {
		log.Printf("Found %d workflow files", len(files))
	}

	// 3. Parse and update files (parser + updater modules)
	totalUpdates, err := newupdater.UpdateWorkflows(ctx, fsys)
	if err != nil {
		log.Fatalf("Failed to update workflows: %v", err)
	}

	// 4. Results
	if *dryRun {
		log.Printf("[DRY RUN] Would update %d action references", totalUpdates)
	} else {
		log.Printf("Updated %d action references in %v", totalUpdates, time.Since(start).Round(time.Millisecond))
	}

	if *verbose {
		for _, file := range files {
			fmt.Printf("- Processed: %s\n", file)
		}
	}
}
