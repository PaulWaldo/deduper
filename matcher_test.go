package main

import (
	"testing"
	"testing/fstest"
)

func TestMatcherIdentifiesExactDuplicates(t *testing.T) {
	// Create an in-memory filesystem
	memFS := fstest.MapFS{
		"original.txt":      &fstest.MapFile{Data: []byte("test content")},
		"duplicate.txt":     &fstest.MapFile{Data: []byte("test content")},
		"non_duplicate.txt": &fstest.MapFile{Data: []byte("different content")},
	}

	// Create a matcher with the in-memory filesystem
	matcher := NewMatcher(memFS)

	// Test exact duplicate detection
	isDuplicate, err := matcher.IsExactDuplicate("original.txt", "duplicate.txt")
	if err != nil {
		t.Fatalf("IsExactDuplicate() returned error: %v", err)
	}
	if !isDuplicate {
		t.Errorf("IsExactDuplicate() = false, want true for exact duplicates")
	}

	// Test non-duplicate detection
	isDuplicate, err = matcher.IsExactDuplicate("original.txt", "non_duplicate.txt")
	if err != nil {
		t.Fatalf("IsExactDuplicate() returned error: %v", err)
	}
	if isDuplicate {
		t.Errorf("IsExactDuplicate() = true, want false for different content")
	}
}

func TestMatcherConfirmsDuplicatesInScanResults(t *testing.T) {
	// Create an in-memory filesystem
	memFS := fstest.MapFS{
		"artist/album/song.txt":   &fstest.MapFile{Data: []byte("test content")},
		"artist/album/song 1.txt": &fstest.MapFile{Data: []byte("test content")},
		"artist/album/song 2.txt": &fstest.MapFile{Data: []byte("test content")},
		"artist/album/song 3.txt": &fstest.MapFile{Data: []byte("different content")},
	}

	// Create mock scan results
	scanResults := []ScanResult{
		{
			Original:   "artist/album/song.txt",
			Duplicates: []string{"artist/album/song 1.txt", "artist/album/song 2.txt", "artist/album/song 3.txt"},
		},
	}

	// Create a matcher with the in-memory filesystem
	matcher := NewMatcher(memFS)

	// Confirm duplicates in scan results
	matchResults, err := matcher.ConfirmDuplicates(scanResults)
	if err != nil {
		t.Fatalf("ConfirmDuplicates() returned error: %v", err)
	}

	// We expect one match result
	if len(matchResults) != 1 {
		t.Fatalf("ConfirmDuplicates() returned %d results, want 1", len(matchResults))
	}

	// Check the match result
	matchResult := matchResults[0]

	// The original file should be "song.txt"
	if matchResult.Original != "artist/album/song.txt" {
		t.Errorf("Original = %q, want %q", matchResult.Original, "artist/album/song.txt")
	}

	// There should be 2 confirmed duplicates
	if len(matchResult.Duplicates) != 2 {
		t.Errorf("Found %d confirmed duplicates, want 2", len(matchResult.Duplicates))
	}

	// Check that the duplicates are what we expect
	expectedDuplicates := map[string]bool{
		"artist/album/song 1.txt": true,
		"artist/album/song 2.txt": true,
	}

	for _, duplicate := range matchResult.Duplicates {
		if !expectedDuplicates[duplicate] {
			t.Errorf("Unexpected duplicate: %q", duplicate)
		}
	}

	// Check that the non-duplicate is in the non-duplicates list
	if len(matchResult.NonDuplicates) != 1 {
		t.Errorf("Found %d non-duplicates, want 1", len(matchResult.NonDuplicates))
	}

	if matchResult.NonDuplicates[0] != "artist/album/song 3.txt" {
		t.Errorf("NonDuplicate = %q, want %q", matchResult.NonDuplicates[0], "artist/album/song 3.txt")
	}
}
