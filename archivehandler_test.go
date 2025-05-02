package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestHandlerArchivesDuplicates(t *testing.T) {
	// Create temporary test directories
	tempDir := t.TempDir()
	archiveDir := filepath.Join(tempDir, "archive")

	// Create a test directory structure
	testDir := filepath.Join(tempDir, "test")
	artistDir := filepath.Join(testDir, "artist")
	albumDir := filepath.Join(artistDir, "album")

	// Create the directory structure
	err := os.MkdirAll(albumDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create original file
	originalPath := filepath.Join(albumDir, "song.txt")
	err = os.WriteFile(originalPath, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create duplicate files
	duplicate1Path := filepath.Join(albumDir, "song 1.txt")
	err = os.WriteFile(duplicate1Path, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	duplicate2Path := filepath.Join(albumDir, "song 2.txt")
	err = os.WriteFile(duplicate2Path, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create match results
	matchResults := []MatchResult{
		{
			Original:   originalPath,
			Duplicates: []string{duplicate1Path, duplicate2Path},
		},
	}

	// Create a handler
	handler := NewArchive(archiveDir)

	// Archive the duplicates
	err = handler.ArchiveDuplicates(matchResults, false)
	if err != nil {
		t.Fatalf("ArchiveDuplicates() returned error: %v", err)
	}

	// Check that the original file still exists
	if _, err := os.Stat(originalPath); os.IsNotExist(err) {
		t.Errorf("Original file %s was moved or deleted", originalPath)
	}

	// Check that the duplicates were moved to the archive
	archivedDuplicate1 := filepath.Join(archiveDir, "test", "artist", "album", "song 1.txt")
	if _, err := os.Stat(archivedDuplicate1); os.IsNotExist(err) {
		t.Errorf("Duplicate file %s was not moved to archive", archivedDuplicate1)
	}

	archivedDuplicate2 := filepath.Join(archiveDir, "test", "artist", "album", "song 2.txt")
	if _, err := os.Stat(archivedDuplicate2); os.IsNotExist(err) {
		t.Errorf("Duplicate file %s was not moved to archive", archivedDuplicate2)
	}

	// Check that the duplicates no longer exist in the original location
	if _, err := os.Stat(duplicate1Path); !os.IsNotExist(err) {
		t.Errorf("Duplicate file %s still exists in original location", duplicate1Path)
	}

	if _, err := os.Stat(duplicate2Path); !os.IsNotExist(err) {
		t.Errorf("Duplicate file %s still exists in original location", duplicate2Path)
	}
}

func TestHandlerDryRun(t *testing.T) {
	// Create temporary test directories
	tempDir := t.TempDir()
	archiveDir := filepath.Join(tempDir, "archive")

	// Create a test directory structure
	testDir := filepath.Join(tempDir, "test")
	artistDir := filepath.Join(testDir, "artist")
	albumDir := filepath.Join(artistDir, "album")

	// Create the directory structure
	err := os.MkdirAll(albumDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Create original file
	originalPath := filepath.Join(albumDir, "song.txt")
	err = os.WriteFile(originalPath, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create duplicate files
	duplicate1Path := filepath.Join(albumDir, "song 1.txt")
	err = os.WriteFile(duplicate1Path, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	duplicate2Path := filepath.Join(albumDir, "song 2.txt")
	err = os.WriteFile(duplicate2Path, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Create match results
	matchResults := []MatchResult{
		{
			Original:   originalPath,
			Duplicates: []string{duplicate1Path, duplicate2Path},
		},
	}

	// Create a handler
	handler := NewArchive(archiveDir)

	// Archive the duplicates in dry-run mode
	err = handler.ArchiveDuplicates(matchResults, true)
	if err != nil {
		t.Fatalf("ArchiveDuplicates() returned error: %v", err)
	}

	// Check that the original file still exists
	if _, err := os.Stat(originalPath); os.IsNotExist(err) {
		t.Errorf("Original file %s was moved or deleted", originalPath)
	}

	// Check that the duplicates still exist in the original location
	if _, err := os.Stat(duplicate1Path); os.IsNotExist(err) {
		t.Errorf("Duplicate file %s was moved in dry-run mode", duplicate1Path)
	}

	if _, err := os.Stat(duplicate2Path); os.IsNotExist(err) {
		t.Errorf("Duplicate file %s was moved in dry-run mode", duplicate2Path)
	}

	// Check that the archive directory was not created
	if _, err := os.Stat(archiveDir); !os.IsNotExist(err) {
		t.Errorf("Archive directory %s was created in dry-run mode", archiveDir)
	}
}
