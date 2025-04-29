package main

import (
	"path/filepath"
	"testing"
)

func TestScannerIdentifiesPotentialDuplicates(t *testing.T) {
	// Create a scanner with the test_data directory
	scanner := NewScanner("test_data")

	// Test cases for potential duplicates
	testCases := []struct {
		fileName string
		isDuplicate bool
		originalName string
	}{
		{"song.txt", false, "song.txt"},
		{"song 1.txt", true, "song.txt"},
		{"song 2.txt", true, "song.txt"},
		{"song_a.txt", false, "song_a.txt"},
		{"song_a 1.txt", true, "song_a.txt"},
		{"song_a 2.txt", true, "song_a.txt"},
		{"song with spaces.txt", false, "song with spaces.txt"},
		{"song with spaces 1.txt", true, "song with spaces.txt"},
	}

	for _, tc := range testCases {
		t.Run(tc.fileName, func(t *testing.T) {
			// Check if the file is identified as a potential duplicate
			isDuplicate := scanner.IsPotentialDuplicate(tc.fileName)
			if isDuplicate != tc.isDuplicate {
				t.Errorf("IsPotentialDuplicate(%q) = %v, want %v", tc.fileName, isDuplicate, tc.isDuplicate)
			}

			// Check if the original name is correctly extracted
			originalName := scanner.GetOriginalName(tc.fileName)
			if originalName != tc.originalName {
				t.Errorf("GetOriginalName(%q) = %q, want %q", tc.fileName, originalName, tc.originalName)
			}
		})
	}
}

func TestScannerFindsGroupsOfDuplicates(t *testing.T) {
	// Use the test_data directory which contains our test files
	scanner := NewScanner("test_data")

	// Scan for potential duplicates
	results, err := scanner.Scan()
	if err != nil {
		t.Fatalf("Scan() returned error: %v", err)
	}

	// We expect to find one group of duplicates in the test data
	if len(results) != 1 {
		t.Fatalf("Scan() found %d groups, want 1", len(results))
	}

	// Check the first group
	result := results[0]

	// The original file should be "song_a.txt"
	expectedOriginal := filepath.Join("test_data", "artist1", "album1", "song_a.txt")
	if result.Original != expectedOriginal {
		t.Errorf("Original = %q, want %q", result.Original, expectedOriginal)
	}

	// There should be 2 duplicates
	if len(result.Duplicates) != 2 {
		t.Errorf("Found %d duplicates, want 2", len(result.Duplicates))
	}

	// Check that the duplicates are what we expect
	expectedDuplicates := map[string]bool{
		filepath.Join("test_data", "artist1", "album1", "song_a 1.txt"): true,
		filepath.Join("test_data", "artist1", "album1", "song_a 2.txt"): true,
	}

	for _, duplicate := range result.Duplicates {
		if !expectedDuplicates[duplicate] {
			t.Errorf("Unexpected duplicate: %q", duplicate)
		}
	}
}
