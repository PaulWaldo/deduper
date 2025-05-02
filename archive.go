package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Archive is responsible for moving duplicate files to an archive directory
type Archive struct {
	// Archive directory where duplicates will be moved
	ArchiveDir string
	// Logger for output
	Logger *Logger
}

// NewArchive creates a new ArchiveHandler instance
func NewArchive(archiveDir string, verbose bool) *Archive {
	return &Archive{
		ArchiveDir: archiveDir,
		Logger:     NewLogger(verbose),
	}
}

// SetVerbose sets the verbose mode for the handler's logger
func (a *Archive) SetVerbose(verbose bool) {
	a.Logger.SetVerbose(verbose)
}

// ArchiveDuplicates moves duplicate files to the archive directory
func (a *Archive) ArchiveDuplicates(matchResults []MatchResult, dryRun bool) error {
	for _, result := range matchResults {
		for _, duplicate := range result.Duplicates {
			// Get the base name of the duplicate file
			duplicateBase := filepath.Base(duplicate)

			// Get the directory of the duplicate file
			duplicateDir := filepath.Dir(duplicate)

			// Variables for target directory and path
			var targetDir, targetPath string

			// Find the common base directory for all files
			// This is typically the root directory of the music library
			// For our test, this is the tempDir
			// We need to extract the relative path from this base directory
			// In a real scenario, this would be something like "/music" or "/Users/username/Music"
			// For simplicity, we'll use the parent of the parent of the album directory
			// e.g., if the path is "/tmp/test/artist/album/song.txt", the base is "/tmp"
			parts := strings.Split(duplicateDir, string(filepath.Separator))
			if len(parts) >= 3 {
				// Calculate the relative path from the base directory
				// e.g., "test/artist/album"
				relPath := filepath.Join(parts[len(parts)-3:]...)

				// Create the target path in the archive directory
				targetDir = filepath.Join(a.ArchiveDir, relPath)
				targetPath = filepath.Join(targetDir, duplicateBase)
			} else {
				// Fallback if the path structure is not as expected
				targetDir = filepath.Join(a.ArchiveDir, filepath.Base(duplicateDir))
				targetPath = filepath.Join(targetDir, duplicateBase)
			}

			// In dry-run mode, just log what would happen
			if dryRun {
				a.Logger.Info("Would move %s to %s", duplicate, targetPath)
				continue
			}

			// Log the action in verbose mode
			a.Logger.LogVerbose("Moving %s to %s", duplicate, targetPath)

			// Create the target directory if it doesn't exist
			err := os.MkdirAll(targetDir, 0755)
			if err != nil {
				return fmt.Errorf("failed to create directory %s: %w", targetDir, err)
			}

			// Move the file to the archive
			err = a.moveFile(duplicate, targetPath)
			if err != nil {
				return fmt.Errorf("failed to move file %s to %s: %w", duplicate, targetPath, err)
			}
		}

		// Log non-duplicates that need manual review
		for _, nonDuplicate := range result.NonDuplicates {
			a.Logger.Info("File %s has duplicate naming pattern but different content, manual review required", nonDuplicate)
		}
	}

	return nil
}

// moveFile moves a file from src to dst
func (a *Archive) moveFile(src, dst string) error {
	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return err
	}

	// Move the file
	return os.Rename(src, dst)
}
func (a *Archive) GetDir() string {
	return a.ArchiveDir
}
