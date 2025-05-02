package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
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

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	rootDir := flag.Arg(0)

	info, err := os.Stat(rootDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: %s is not a directory\n", rootDir)
		os.Exit(1)
	}

	logger := NewLogger(*verbose)

	logger.Info("Scanning directory: %s", rootDir)
	logger.Info("Archive directory: %s", *archiveDir)
	if *dryRun {
		logger.Info("Dry run mode: no files will be moved")
	}
	if *verbose {
		logger.Info("Verbose mode: showing detailed information")
	}

	fs := os.DirFS(rootDir)
	scanner := NewScan(fs, rootDir)
	matcher := NewMatch(fs, rootDir)
	archiver := NewArchive(*archiveDir, *verbose)
	totalDuplicates, hasNonDuplicates := Dedup(scanner, matcher, archiver, *dryRun, *verbose)
	if err != nil {
		logger.Error("Failed to dedup files: %v", err)
		os.Exit(1)
	}

	logger.Info("Found %d confirmed duplicate files", totalDuplicates)

	if hasNonDuplicates {
		logger.Info("\nSome files have duplicate naming patterns but different content.")
		logger.Info("These files were not moved and require manual review.")
		logger.Info("See above for details.")
	}
}
