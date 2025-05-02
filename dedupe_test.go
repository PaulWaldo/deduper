package main

import (
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestDedup(t *testing.T) {
	root := "blah"
	memFS := fstest.MapFS{
		filepath.Join(root, "artist1/album1/song_a.txt"):   &fstest.MapFile{Data: []byte("test content")},
		filepath.Join(root, "artist1/album1/song_a 1.txt"): &fstest.MapFile{Data: []byte("test content")},
		filepath.Join(root, "artist1/album1/song_a 2.txt"): &fstest.MapFile{Data: []byte("different content")},
	}

	scanner := NewScan(memFS, ".")
	matcher := NewMatch(memFS, ".")
	archiver := NewArchive("/tmp/archive", true)

	totalDuplicates, hasNonDuplicates := Dedup(scanner, matcher, archiver, false, false)
	if totalDuplicates != 1 {
		t.Errorf("Expected 1 duplicate, got %d", totalDuplicates)
	}
	if hasNonDuplicates {
		t.Error("Expected no non-duplicates")
	}
}

// func TestDedup_NoDuplicates(t *testing.T) {
// 	memFS := fstest.MapFS{
// 		"file1.txt": {Data: []byte("content1")},
// 		"file2.txt": {Data: []byte("content2")},
// 	}
// 	totalDuplicates, hasNonDuplicates := Dedup("", "archive", false, false, memFS)
// 	if totalDuplicates != 0 {
// 		t.Errorf("Expected 0 duplicates, got %d", totalDuplicates)
// 	}
// 	if hasNonDuplicates {
// 		t.Error("Expected no non-duplicates")
// 	}
// }

// func TestDedup_DryRun(t *testing.T) {
// 	memFS := fstest.MapFS{
// 		"file1.txt": {Data: []byte("content1")},
// 		"file2.txt": {Data: []byte("content1")},
// 	}
// 	totalDuplicates, hasNonDuplicates := Dedup("", "archive", true, false, memFS)
// 	if totalDuplicates != 1 {
// 		t.Errorf("Expected 1 duplicate, got %d", totalDuplicates)
// 	}
// 	if hasNonDuplicates {
// 		t.Error("Expected no non-duplicates")
// 	}
// }

// func TestDedup_Error(t *testing.T) {
// 	memFS := fstest.MapFS{
// 		"file1.txt": {Data: []byte("content1")},
// 		"file2.txt": {Data: []byte("content1")},
// 	}
// 	archiveHandler := &Archive{ArchiveDir: "./archive"}
// 	archiveHandler.ArchiveDuplicates = func(matchResults []MatchResult, dryRun bool) error {
// 		return fmt.Errorf("simulated error")
// 	}
// 	totalDuplicates, hasNonDuplicates := Dedup("", "./archive", false, false, memFS)
// 	if totalDuplicates != 0 {
// 		t.Errorf("Expected 0 duplicates, got %d", totalDuplicates)
// 	}
// 	if !hasNonDuplicates {
// 		t.Error("Expected non-duplicates")
// 	}
// }
