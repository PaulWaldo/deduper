package main

// import (
// 	"os"
// 	"path/filepath"
// 	"testing"
// 	"testing/fstest"
// )

// // TestMatcherWithAbsolutePaths tests that the matcher fails when given absolute paths
// // but using a DirFS rooted at a different location
// func TestMatcherWithAbsolutePaths(t *testing.T) {
// 	root := "root"
// 	// Create an in-memory filesystem with specific files
// 	memFS := fstest.MapFS{
// 		filepath.Join(root, "file1.txt"): &fstest.MapFile{Data: []byte("test content")},
// 		filepath.Join(root, "file2.txt"): &fstest.MapFile{Data: []byte("test content")},
// 		filepath.Join(root, "file3.txt"): &fstest.MapFile{Data: []byte("different content")},
// 	}

// 	// Create a matcher with the in-memory filesystem
// 	matcher := NewMatcher(memFS, root)

// 	// Test with paths that exist in the filesystem
// 	isDuplicate, err := matcher.IsExactDuplicate("root/file1.txt", "root/file2.txt")
// 	if err != nil {
// 		t.Fatalf("IsExactDuplicate() returned error for valid paths: %v", err)
// 	}
// 	if !isDuplicate {
// 		t.Fatalf("IsExactDuplicate() = false, want true for exact duplicates")
// 	}

// 	// Test with absolute paths that don't exist in the filesystem
// 	// These paths simulate what would happen if we used absolute paths from the scanner
// 	// with a matcher that has a filesystem rooted elsewhere
// 	_, err = matcher.IsExactDuplicate("/absolute/path/to/file1.txt", "/absolute/path/to/file2.txt")

// 	// We MUST get an error here because the absolute paths don't exist in the in-memory filesystem
// 	if err == nil {
// 		t.Fatalf("CRITICAL ERROR: Expected error when using absolute paths with in-memory filesystem, but got nil")
// 	} else {
// 		t.Logf("Got expected error when using absolute paths: %v", err)
// 	}
// }

// // TestScannerMatcherIntegration simulates the integration between scanner and matcher
// // to verify the issue with absolute paths
// func TestScannerMatcherIntegration(t *testing.T) {
// 	// Create an in-memory filesystem
// 	memFS := fstest.MapFS{
// 		"artist/album/song.txt":   &fstest.MapFile{Data: []byte("test content")},
// 		"artist/album/song 1.txt": &fstest.MapFile{Data: []byte("test content")},
// 		"artist/album/song 2.txt": &fstest.MapFile{Data: []byte("test content")},
// 	}

// 	// Create a matcher with the in-memory filesystem
// 	matcher := NewMatcher(memFS,"/")

// 	// Create mock scan results with absolute paths
// 	// This simulates what the scanner would return in the real application
// 	absoluteBasePath := "/absolute/path/to"
// 	scanResults := []ScanResult{
// 		{
// 			Original: filepath.Join(absoluteBasePath, "artist/album/song.txt"),
// 			Duplicates: []string{
// 				filepath.Join(absoluteBasePath, "artist/album/song 1.txt"),
// 				filepath.Join(absoluteBasePath, "artist/album/song 2.txt"),
// 			},
// 		},
// 	}

// 	// Try to confirm duplicates with absolute paths
// 	// This should fail because the absolute paths don't exist in the in-memory filesystem
// 	_, err := matcher.ConfirmDuplicates(scanResults)
// 	if err == nil {
// 		t.Errorf("Expected error when using absolute paths with in-memory filesystem, but got nil")
// 	}

// 	// Now create mock scan results with paths relative to the filesystem root
// 	relativeResults := []ScanResult{
// 		{
// 			Original: "artist/album/song.txt",
// 			Duplicates: []string{
// 				"artist/album/song 1.txt",
// 				"artist/album/song 2.txt",
// 			},
// 		},
// 	}

// 	// Try to confirm duplicates with relative paths
// 	// This should succeed
// 	matchResults, err := matcher.ConfirmDuplicates(relativeResults)
// 	if err != nil {
// 		t.Errorf("Unexpected error when using relative paths with in-memory filesystem: %v", err)
// 	}
// 	if len(matchResults) != 1 {
// 		t.Errorf("ConfirmDuplicates() returned %d results, want 1", len(matchResults))
// 	}
// 	if len(matchResults[0].Duplicates) != 2 {
// 		t.Errorf("Found %d confirmed duplicates, want 2", len(matchResults[0].Duplicates))
// 	}
// }

// // TestMainIntegrationIssue simulates the exact scenario in main.go where we're using
// // os.DirFS(".") with absolute paths from the scanner. This test demonstrates the issue.
// func TestMainIntegrationIssue(t *testing.T) {
// 	// Skip this test when running with go test -short
// 	if testing.Short() {
// 		t.Skip("Skipping test in short mode")
// 	}

// 	// Create a temporary directory for testing
// 	tempDir := t.TempDir()

// 	// Create a test directory structure
// 	artistDir := filepath.Join(tempDir, "artist")
// 	albumDir := filepath.Join(artistDir, "album")
// 	err := os.MkdirAll(albumDir, 0755)
// 	if err != nil {
// 		t.Fatalf("Failed to create test directory: %v", err)
// 	}

// 	// Create original file
// 	originalPath := filepath.Join(albumDir, "song.txt")
// 	err = os.WriteFile(originalPath, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	// Create duplicate files
// 	duplicate1Path := filepath.Join(albumDir, "song 1.txt")
// 	err = os.WriteFile(duplicate1Path, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	duplicate2Path := filepath.Join(albumDir, "song 2.txt")
// 	err = os.WriteFile(duplicate2Path, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	// Create mock scan results with absolute paths
// 	// This simulates what the scanner would return in the real application
// 	scanResults := []ScanResult{
// 		{
// 			Original: originalPath,
// 			Duplicates: []string{
// 				duplicate1Path,
// 				duplicate2Path,
// 			},
// 		},
// 	}

// 	// Create a matcher with DirFS rooted at the current directory
// 	// This is exactly what we're doing in main.go
// 	matcher := NewMatcher(os.DirFS(tempDir))

// 	// Try to confirm duplicates
// 	// This should fail because the absolute paths don't exist in the DirFS
// 	_, err = matcher.ConfirmDuplicates(scanResults)

// 	// We expect an error because the absolute paths don't exist in the DirFS
// 	if err != nil {
// 		t.Fatalf("Error confirming duplicates: %s", err)
// 	}
// }

// // TestMainIntegrationFix demonstrates the fix for the issue in TestMainIntegrationIssue
// // by using os.DirFS(tempDir) instead of os.DirFS(".")
// func TestMainIntegrationFix(t *testing.T) {
// 	// Skip this test when running with go test -short
// 	if testing.Short() {
// 		t.Skip("Skipping test in short mode")
// 	}

// 	// Create a temporary directory for testing
// 	tempDir := t.TempDir()

// 	// Create a test directory structure
// 	artistDir := filepath.Join(tempDir, "artist")
// 	albumDir := filepath.Join(artistDir, "album")
// 	err := os.MkdirAll(albumDir, 0755)
// 	if err != nil {
// 		t.Fatalf("Failed to create test directory: %v", err)
// 	}

// 	// Create original file
// 	originalPath := filepath.Join(albumDir, "song.txt")
// 	err = os.WriteFile(originalPath, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	// Create duplicate files
// 	duplicate1Path := filepath.Join(albumDir, "song 1.txt")
// 	err = os.WriteFile(duplicate1Path, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	duplicate2Path := filepath.Join(albumDir, "song 2.txt")
// 	err = os.WriteFile(duplicate2Path, []byte("test content"), 0644)
// 	if err != nil {
// 		t.Fatalf("Failed to create test file: %v", err)
// 	}

// 	// Create mock scan results with absolute paths
// 	// This simulates what the scanner would return in the real application
// 	scanResults := []ScanResult{
// 		{
// 			Original: originalPath,
// 			Duplicates: []string{
// 				duplicate1Path,
// 				duplicate2Path,
// 			},
// 		},
// 	}

// 	// Convert absolute paths to paths relative to the temp directory
// 	for i := range scanResults {
// 		relOriginal, err := filepath.Rel(tempDir, scanResults[i].Original)
// 		if err != nil {
// 			t.Fatalf("Failed to get relative path: %v", err)
// 		}
// 		scanResults[i].Original = relOriginal

// 		for j := range scanResults[i].Duplicates {
// 			relDuplicate, err := filepath.Rel(tempDir, scanResults[i].Duplicates[j])
// 			if err != nil {
// 				t.Fatalf("Failed to get relative path: %v", err)
// 			}
// 			scanResults[i].Duplicates[j] = relDuplicate
// 		}
// 	}

// 	// Create a matcher with DirFS rooted at the temp directory
// 	// This is the fix for the issue in main.go
// 	matcher := NewMatcher(os.DirFS(tempDir))

// 	// Try to confirm duplicates
// 	// This should succeed because we're using relative paths with a DirFS rooted at the temp directory
// 	matchResults, err := matcher.ConfirmDuplicates(scanResults)

// 	// We expect no error because we're using relative paths with a DirFS rooted at the temp directory
// 	if err != nil {
// 		t.Fatalf("Unexpected error when using relative paths with DirFS rooted at tempDir: %v", err)
// 	}

// 	// Verify that we found the expected duplicates
// 	if len(matchResults) != 1 {
// 		t.Fatalf("ConfirmDuplicates() returned %d results, want 1", len(matchResults))
// 	}
// 	if len(matchResults[0].Duplicates) != 2 {
// 		t.Fatalf("Found %d confirmed duplicates, want 2", len(matchResults[0].Duplicates))
// 	}

// 	t.Logf("Successfully confirmed duplicates using relative paths with DirFS rooted at tempDir")
// }

// // TestFixedScannerMatcherIntegration tests the proposed solution of using a filesystem
// // rooted at the same directory as the scanner
// func TestFixedScannerMatcherIntegration(t *testing.T) {
// 	// Create an in-memory filesystem for the root directory
// 	rootFS := fstest.MapFS{
// 		"artist/album/song.txt":   &fstest.MapFile{Data: []byte("test content")},
// 		"artist/album/song 1.txt": &fstest.MapFile{Data: []byte("test content")},
// 		"artist/album/song 2.txt": &fstest.MapFile{Data: []byte("test content")},
// 	}

// 	// Create mock scan results with absolute paths
// 	// This simulates what the scanner would return in the real application
// 	rootDir := "/absolute/path/to"
// 	scanResults := []ScanResult{
// 		{
// 			Original: filepath.Join(rootDir, "artist/album/song.txt"),
// 			Duplicates: []string{
// 				filepath.Join(rootDir, "artist/album/song 1.txt"),
// 				filepath.Join(rootDir, "artist/album/song 2.txt"),
// 			},
// 		},
// 	}

// 	// In the real application, we would create a matcher with os.DirFS(rootDir)
// 	// Here we simulate that by using the in-memory filesystem and modifying the paths

// 	// Convert absolute paths to paths relative to the root directory
// 	for i := range scanResults {
// 		// In a real application with os.DirFS(rootDir), we would use filepath.Rel(rootDir, path)
// 		// Here we simulate that by removing the rootDir prefix
// 		scanResults[i].Original = scanResults[i].Original[len(rootDir)+1:]
// 		for j := range scanResults[i].Duplicates {
// 			scanResults[i].Duplicates[j] = scanResults[i].Duplicates[j][len(rootDir)+1:]
// 		}
// 	}

// 	// Create a matcher with the in-memory filesystem
// 	matcher := NewMatcher(rootFS)

// 	// Try to confirm duplicates with the modified paths
// 	// This should succeed
// 	matchResults, err := matcher.ConfirmDuplicates(scanResults)
// 	if err != nil {
// 		t.Errorf("Unexpected error when using relative paths with in-memory filesystem: %v", err)
// 	}
// 	if len(matchResults) != 1 {
// 		t.Errorf("ConfirmDuplicates() returned %d results, want 1", len(matchResults))
// 	}
// 	if len(matchResults[0].Duplicates) != 2 {
// 		t.Errorf("Found %d confirmed duplicates, want 2", len(matchResults[0].Duplicates))
// 	}
// }
