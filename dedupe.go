package main

type Scanner interface {
	DirHolder
	Scan() ([]ScanResult, error)
}

type Matcher interface {
	ConfirmDuplicates(scanResults []ScanResult) ([]MatchResult, error)
}

type Archiver interface {
	DirHolder
	SetVerbose(verbose bool)
	ArchiveDuplicates(matchResults []MatchResult, dryRun bool) error
}

type DirHolder interface {
	GetDir() string
}

// Dedup performs duplicate file detection and archiving.
func Dedup(scanner Scanner, matcher Matcher, archiver Archiver, dryRun bool, verbose bool) (int, bool) {
	logger := NewLogger(verbose)

	logger.Info("Scanning directory: %s", scanner.GetDir())
	logger.Info("Archive directory: %s", archiver.GetDir())
	if dryRun {
		logger.Info("Dry run mode: no files will be moved")
	}
	if verbose {
		logger.Info("Verbose mode: showing detailed information")
	}

	logger.LogVerbose("Scanning for potential duplicates...")
	scanResults, err := scanner.Scan()
	if err != nil {
		logger.Error("Failed to scan directory: %v", err)
		return 0, false // Return error
	}

	logger.LogVerbose("Found %d groups of potential duplicates", len(scanResults))
	if len(scanResults) == 0 {
		logger.Info("No potential duplicates found")
		return 0, false
	}

	logger.LogVerbose("Confirming exact duplicates...")
	logger.Info("Scan Results: %+V", scanResults)
	matchResults, err := matcher.ConfirmDuplicates(scanResults)
	if err != nil {
		logger.Error("Failed to confirm duplicates: %v", err)
		return 0, false // Return error
	}

	logger.LogVerbose("Found %d groups of confirmed duplicates", len(matchResults))
	if len(matchResults) == 0 {
		logger.Info("No confirmed duplicates found")
		return 0, false
	}

	totalDuplicates := 0
	for _, result := range matchResults {
		totalDuplicates += len(result.Duplicates)
	}
	logger.Info("Found %d confirmed duplicate files", totalDuplicates)

	archiver.SetVerbose(verbose)
	logger.LogVerbose("Archiving duplicates...")
	err = archiver.ArchiveDuplicates(matchResults, dryRun)
	if err != nil {
		logger.Error("Failed to archive duplicates: %v", err)
		return 0, false // Return error
	}

	hasNonDuplicates := false
	for _, result := range matchResults {
		if len(result.NonDuplicates) > 0 {
			hasNonDuplicates = true
			break
		}
	}

	if dryRun {
		logger.Info("Dry run completed. %d duplicate files would be moved to %s", totalDuplicates, archiver.GetDir())
	} else {
		logger.Info("Completed. %d duplicate files moved to %s", totalDuplicates, archiver.GetDir())
	}

	if hasNonDuplicates {
		logger.Info("\nSome files have duplicate naming patterns but different content.")
		logger.Info("These files were not moved and require manual review.")
		logger.Info("See above for details.")
	}

	return totalDuplicates, hasNonDuplicates
}
