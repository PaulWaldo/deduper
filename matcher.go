package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io/fs"
	"path/filepath"
)

// Match is responsible for comparing file contents to confirm exact duplicates
type Match struct {
	// FileSystem is the filesystem to use for file operations
	FileSystem fs.FS
	root       string
}

// NewMatch creates a new Matcher instance
func NewMatch(fileSystem fs.FS, root string) *Match {
	return &Match{
		FileSystem: fileSystem,
		root:       root,
	}
}

// MatchResult represents the result of confirming duplicates in a scan result
type MatchResult struct {
	// The original file
	Original string
	// Confirmed duplicates of the original file
	Duplicates []string
	// Files that were potential duplicates but have different content
	NonDuplicates []string
}

// IsExactDuplicate checks if two files have exactly the same content
func (m *Match) IsExactDuplicate(file1, file2 string) (bool, error) {
	// Read the content of both files
	relFile1, err := filepath.Rel(m.root, file1)
	if err != nil {
		return false, err
	}
	content1, err := fs.ReadFile(m.FileSystem, relFile1)
	if err != nil {
		return false, fmt.Errorf("failed to read file %s: %w", file1, err)
	}

	relFile2, err := filepath.Rel(m.root, file2)
	if err != nil {
		return false, err
	}
	content2, err := fs.ReadFile(m.FileSystem, relFile2)
	if err != nil {
		return false, fmt.Errorf("failed to read file %s: %w", file2, err)
	}

	// Compare the content
	return bytes.Equal(content1, content2), nil
}

// ConfirmDuplicates confirms which potential duplicates are exact duplicates
func (m *Match) ConfirmDuplicates(scanResults []ScanResult) ([]MatchResult, error) {
	var results []MatchResult

	for _, scanResult := range scanResults {
		var duplicates []string
		var nonDuplicates []string

		for _, potentialDuplicate := range scanResult.Duplicates {
			isDuplicate, err := m.IsExactDuplicate(scanResult.Original, potentialDuplicate)
			if err != nil {
				return nil, fmt.Errorf("failed to check if %s is a duplicate of %s: %w", potentialDuplicate, scanResult.Original, err)
			}

			if isDuplicate {
				duplicates = append(duplicates, potentialDuplicate)
			} else {
				nonDuplicates = append(nonDuplicates, potentialDuplicate)
			}
		}

		// Only add a result if there are confirmed duplicates
		if len(duplicates) > 0 {
			results = append(results, MatchResult{
				Original:      scanResult.Original,
				Duplicates:    duplicates,
				NonDuplicates: nonDuplicates,
			})
		}
	}

	return results, nil
}

// GetFileHash calculates the SHA-256 hash of a file's content
func (m *Match) GetFileHash(filePath string) (string, error) {
	content, err := fs.ReadFile(m.FileSystem, filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read file %s: %w", filePath, err)
	}

	hash := sha256.Sum256(content)
	return fmt.Sprintf("%x", hash), nil
}
