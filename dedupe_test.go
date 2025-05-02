package main

import (
	"testing"
	"testing/fstest"
)

func TestDedup(t *testing.T) {
	memFS := fstest.MapFS{
		"./dir1/file1.txt":                {Data: []byte("content1")},
		"./dir1/subdir/file2.txt":         {Data: []byte("content2")},
		"./dir1/file3.txt":                {Data: []byte("content1")},
		"./archive/dir1/file1.txt":        {Data: []byte("content1")},
		"./archive/dir1/file3.txt":        {Data: []byte("content1")},
		"./archive/dir1/subdir":           {},
		"./archive/dir1/subdir/file2.txt": {Data: []byte("content2")},
	}
	scanner := NewScan(".")
	matcher := NewMatch(memFS, ".")
	archiver := NewArchive("./archive")

	totalDuplicates, hasNonDuplicates := Dedup(scanner, matcher, archiver, "dir1", "archive", false, false, memFS)
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
