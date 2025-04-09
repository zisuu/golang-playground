package updater

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

// UpdateActionReferences replaces action tags with SHAs in workflow content
func UpdateActionReferences(content []byte, replacements map[string]string) ([]byte, error) {
	// Convert replacements to a more efficient lookup format
	replacementMap := make(map[string]string)
	for oldRef, newRef := range replacements {
		normalized := normalizeActionReference(oldRef)
		replacementMap[normalized] = newRef
	}

	// Process each line individually to preserve formatting
	lines := bytes.Split(content, []byte("\n"))
	var updatedLines [][]byte
	var changes []string

	for _, line := range lines {
		updatedLine, changed := processLine(line, replacementMap)
		updatedLines = append(updatedLines, updatedLine)
		if changed != "" {
			changes = append(changes, changed)
		}
	}

	if len(changes) == 0 {
		return content, fmt.Errorf("no action references found to update")
	}

	return bytes.Join(updatedLines, []byte("\n")), nil
}

// processLine handles a single line of YAML
func processLine(line []byte, replacements map[string]string) ([]byte, string) {
	// Skip lines without uses:
	if !bytes.Contains(line, []byte("uses:")) {
		return line, ""
	}

	// Find action references in the line
	matches := actionReferencePattern.FindSubmatch(line)
	if len(matches) < 3 {
		return line, ""
	}

	//fullMatch := string(matches[0])
	fullRef := string(matches[1])
	//ref := string(matches[2])

	// Check if we have a replacement for this reference
	newRef, ok := replacements[normalizeActionReference(fullRef)]
	if !ok {
		return line, ""
	}

	// Preserve original whitespace and formatting
	updated := bytes.Replace(line, matches[0], []byte(fmt.Sprintf("uses: %s", newRef)), 1)
	return updated, fmt.Sprintf("%s -> %s", fullRef, newRef)
}

// Regular expression to match action references while preserving formatting
var actionReferencePattern = regexp.MustCompile(`uses:\s*([^\s@]+@[^\s#]+)`)

// normalizeActionReference standardizes references for comparison
func normalizeActionReference(ref string) string {
	return strings.TrimSpace(ref)
}
