package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse command-line arguments
	var (
		dryRun     = flag.Bool("dry-run", false, "Preview actions without performing file moves")
		verbose    = flag.Bool("verbose", false, "Show detailed information about all actions")
		archiveDir = flag.String("archive-dir", "./archive", "Directory for archived duplicates")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <directory>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\nFlags:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Check if a directory was provided
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	// Get the directory to scan
	rootDir := flag.Arg(0)

	// Check if the directory exists
	info, err := os.Stat(rootDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: %s is not a directory\n", rootDir)
		os.Exit(1)
	}

	// Create a logger
	logger := NewLogger(*verbose)

	// Log the configuration
	logger.Info("Scanning directory: %s", rootDir)
	logger.Info("Archive directory: %s", *archiveDir)
	if *dryRun {
		logger.Info("Dry run mode: no files will be moved")
	}
	if *verbose {
		logger.Info("Verbose mode: showing detailed information")
	}

	// Create a scanner
	scanner := NewScanner(rootDir)

	// Scan for potential duplicates
	logger.LogVerbose("Scanning for potential duplicates...")
	scanResults, err := scanner.Scan()
	if err != nil {
		logger.Error("Failed to scan directory: %v", err)
		os.Exit(1)
	}

	// Log the scan results
	logger.LogVerbose("Found %d groups of potential duplicates", len(scanResults))

	// If no potential duplicates were found, exit
	if len(scanResults) == 0 {
		logger.Info("No potential duplicates found")
		return
	}

	// Create a matcher with the OS filesystem
	matcher := NewMatcher(os.DirFS(rootDir), rootDir)

	// Confirm which potential duplicates are exact duplicates
	logger.LogVerbose("Confirming exact duplicates...")
	matchResults, err := matcher.ConfirmDuplicates(scanResults)
	if err != nil {
		logger.Error("Failed to confirm duplicates: %v", err)
		os.Exit(1)
	}

	// Log the match results
	logger.LogVerbose("Found %d groups of confirmed duplicates", len(matchResults))

	// If no confirmed duplicates were found, exit
	if len(matchResults) == 0 {
		logger.Info("No confirmed duplicates found")
		return
	}

	// Count the total number of duplicates
	totalDuplicates := 0
	for _, result := range matchResults {
		totalDuplicates += len(result.Duplicates)
	}
	logger.Info("Found %d confirmed duplicate files", totalDuplicates)

	// Create a handler
	handler := NewHandler(*archiveDir)
	handler.SetVerbose(*verbose)

	// Archive the confirmed duplicates
	logger.LogVerbose("Archiving duplicates...")
	err = handler.ArchiveDuplicates(matchResults, *dryRun)
	if err != nil {
		logger.Error("Failed to archive duplicates: %v", err)
		os.Exit(1)
	}

	// Log the completion
	if *dryRun {
		logger.Info("Dry run completed. %d duplicate files would be moved to %s", totalDuplicates, *archiveDir)
	} else {
		logger.Info("Completed. %d duplicate files moved to %s", totalDuplicates, *archiveDir)
	}

	// Check if there are any non-duplicates that need manual review
	hasNonDuplicates := false
	for _, result := range matchResults {
		if len(result.NonDuplicates) > 0 {
			hasNonDuplicates = true
			break
		}
	}

	if hasNonDuplicates {
		logger.Info("\nSome files have duplicate naming patterns but different content.")
		logger.Info("These files were not moved and require manual review.")
		logger.Info("See above for details.")
	}
}
