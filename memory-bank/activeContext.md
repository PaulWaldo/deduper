# Active Context: Deduper

## Current Work Focus
- Setting up the initial project structure and memory bank
- Defining the core architecture and approach
- Creating comprehensive tests for all functionality
- Planning the implementation of the main functionality

## Recent Changes
- Created the memory bank with project documentation
- Decided on a single command structure with flags rather than subcommands
- Established a flat file structure with separate files for different functionality
- Prioritized test-driven development approach

## Next Steps
1. Create comprehensive test suite
   - Write tests for each component before implementation
   - Ensure tests are properly named for gotestdox compatibility
   - Set up test fixtures in test_data directory

2. Implement the basic CLI structure in main.go with tests
   - Test command-line argument parsing
   - Test flag handling
   - Set up the core workflow with test coverage

3. Create the file scanning functionality in scanner.go with tests
   - Test directory traversal
   - Test duplicate identification based on naming patterns
   - Ensure edge cases are covered

4. Implement content matching in matcher.go with tests
   - Test file content comparison
   - Test hash-based matching
   - Test edge cases like empty files or permission issues

5. Develop file handling in handler.go with tests
   - Test archive directory creation
   - Test file movement while preserving structure
   - Test error handling scenarios

6. Set up logging in logger.go with tests
   - Test different verbosity levels
   - Test error reporting

## Active Decisions and Considerations

### Testing Strategy (HIGH PRIORITY)
- Test-driven development approach is mandatory
- Tests must be created before or alongside implementation
- All tests must pass before considering any feature complete
- Tests should cover edge cases and error scenarios
- Test naming must be compatible with gotestdox for clear reporting

### Command-Line Interface
- Using a single command with flags for simplicity
- Flags include --dry-run, --verbose, and --archive-dir
- Target directory is provided as a positional argument

### File Structure
- Keeping a flat file structure for now
- Separating functionality into different files
- May refactor into packages later if complexity grows

### Duplicate Identification
- Focus on files with naming patterns like "song.txt", "song 1.txt", "song 2.txt"
- Using exact content matching to confirm duplicates
- Considering hash-based comparison for efficiency

### Error Handling
- Need to decide on error handling strategy
- Considering whether to fail fast or continue on errors
- Must provide clear error messages for troubleshooting

## Important Patterns and Preferences

### Testing Patterns
- Write tests first or alongside implementation
- Use table-driven tests for comprehensive coverage
- Test both happy paths and error scenarios
- Ensure tests are isolated and don't depend on each other
- Use meaningful test names that describe the behavior being tested

### Code Organization
- Clear separation of concerns between files
- Consistent error handling approach
- Comprehensive documentation with go-doc comments

### Logging Approach
- Structured logging for clarity
- Different verbosity levels based on --verbose flag
- Clear distinction between normal output and errors

## Learnings and Project Insights

### Test-Driven Development
- Tests drive the design and implementation
- Tests provide confidence in refactoring
- Tests document the expected behavior
- Tests catch regressions early

### Project Structure
- Starting with a simple, flat structure
- Can evolve as complexity grows
- Keeping related functionality in separate files for clarity

### Command-Line Design
- Single command with flags is simpler for users
- Provides flexibility through optional flags
- Aligns with common CLI patterns

### File System Operations
- Need to handle cross-platform path differences
- Must consider performance for large music libraries
- Error handling is critical for file operations
