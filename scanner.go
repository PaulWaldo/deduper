package main

import (
	"os"
	"path/filepath"
	"regexp"
)

// Scanner is responsible for traversing directories and identifying potential duplicate files
type Scanner struct {
	// Root directory to scan
	RootDir string
	// Regular expression to identify potential duplicates
	DuplicatePattern *regexp.Regexp
}

// NewScanner creates a new Scanner instance
func NewScanner(rootDir string) *Scanner {
	// Pattern to match files with a number appended before the extension
	// e.g., "song_a 1.txt" where "song_a.txt" is the original
	pattern := regexp.MustCompile(`^(.+) \d+(\..+)?$`)

	return &Scanner{
		RootDir:          rootDir,
		DuplicatePattern: pattern,
	}
}

// ScanResult represents a group of files that might contain duplicates
type ScanResult struct {
	// The original file (without a number appended)
	Original string
	// Potential duplicates of the original file
	Duplicates []string
}

// Scan traverses the directory structure and returns groups of potential duplicate files
func (s *Scanner) Scan() ([]ScanResult, error) {
	var results []ScanResult

	// Map to group files by their base name (without the number suffix)
	fileGroups := make(map[string][]string)

	// Walk through the directory structure
	err := filepath.Walk(s.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Get the base name of the file
		base := filepath.Base(path)

		// Extract the original name if this is a potential duplicate
		originalName := s.GetOriginalName(base)

		// Get the directory path
		dir := filepath.Dir(path)

		// Create a key that includes the directory to group files correctly
		key := filepath.Join(dir, originalName)

		// Add the file to its group
		fileGroups[key] = append(fileGroups[key], path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Process the file groups to create ScanResults
	for _, files := range fileGroups {
		// Skip groups with only one file (no duplicates)
		if len(files) <= 1 {
			continue
		}

		// Find the original file (the one without a number appended)
		var original string
		var duplicates []string

		for _, file := range files {
			base := filepath.Base(file)
			if !s.IsPotentialDuplicate(base) {
				original = file
			} else {
				duplicates = append(duplicates, file)
			}
		}

		// If we found an original and duplicates, add a ScanResult
		if original != "" && len(duplicates) > 0 {
			results = append(results, ScanResult{
				Original:   original,
				Duplicates: duplicates,
			})
		}
	}

	return results, nil
}

// IsPotentialDuplicate checks if a file name matches the duplicate pattern
func (s *Scanner) IsPotentialDuplicate(fileName string) bool {
	return s.DuplicatePattern.MatchString(fileName)
}

// GetOriginalName returns the original file name for a potential duplicate
func (s *Scanner) GetOriginalName(fileName string) string {
	if match := s.DuplicatePattern.FindStringSubmatch(fileName); match != nil {
		originalName := match[1]
		if match[2] != "" {
			originalName += match[2]
		}
		return originalName
	}
	return fileName
}
