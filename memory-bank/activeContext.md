# Active Context: Deduper

## Current Work Focus
- Implementing the core components of the deduper utility
- Following test-driven development approach for all components
- Connecting components together in the main.go file
- Finalizing the CLI structure with command-line flags

## Recent Changes
- Implemented the Scanner component for identifying potential duplicates
- Implemented the Matcher component for confirming exact duplicates
- Implemented the Handler component for moving duplicates to an archive
- Implemented the Logger component for user feedback
- Created comprehensive tests for all components
- Used in-memory filesystem (testing/fstest) for unit tests

## Next Steps
1. Complete the main.go implementation
   - Implement command-line argument parsing
   - Connect all components together
   - Set up the core workflow

2. Add comprehensive documentation
   - Add go-doc comments to all functions and types
   - Create usage examples
   - Create README with installation and usage instructions

3. Perform integration testing
   - Test the entire workflow with real files
   - Test edge cases and error scenarios
   - Ensure all requirements are met

## Active Decisions and Considerations

### Testing Strategy (HIGH PRIORITY)
- Test-driven development approach has been successful
- All components have comprehensive tests
- All tests are passing
- Used in-memory filesystem for testing file operations

### Command-Line Interface
- Using a single command with flags for simplicity
- Flags include --dry-run, --verbose, and --archive-dir
- Target directory is provided as a positional argument

### File Structure
- Keeping a flat file structure for now
- Separated functionality into different files:
  - scanner.go: Directory traversal and duplicate identification
  - matcher.go: File content comparison
  - handler.go: File movement and archiving
  - logger.go: User feedback and logging
  - main.go: CLI and workflow orchestration

### Duplicate Identification
- Implemented regex-based pattern matching for identifying potential duplicates
- Using exact content comparison to confirm duplicates
- Handling edge cases like files with duplicate naming patterns but different content

### Error Handling
- Using descriptive error messages with fmt.Errorf and %w for wrapping errors
- Providing context in error messages for troubleshooting
- Failing fast on critical errors, but logging non-critical issues

## Important Patterns and Preferences

### Testing Patterns
- Used table-driven tests for comprehensive coverage
- Tested both happy paths and error scenarios
- Used in-memory filesystem for file operation tests
- Created isolated tests that don't depend on each other

### Code Organization
- Clear separation of concerns between files
- Consistent error handling approach
- Comprehensive documentation with go-doc comments

### Logging Approach
- Implemented structured logging with different verbosity levels
- Info level for important information regardless of verbosity
- LogVerbose level for detailed information in verbose mode
- Error level for error messages

## Learnings and Project Insights

### Test-Driven Development
- Tests drove the design and implementation
- Tests provided confidence in refactoring
- Tests documented the expected behavior
- Tests caught regressions early

### Project Structure
- Flat structure worked well for this project
- Separation of concerns made implementation cleaner
- Each component has a clear responsibility

### File System Operations
- Used filepath package for cross-platform path handling
- Implemented careful error handling for file operations
- Used in-memory filesystem for testing to avoid disk I/O
