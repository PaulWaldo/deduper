# Technical Context: Deduper

## Technologies Used

### Go Programming Language
- Deduper is implemented in Go (Golang)
- Go was chosen for its performance, simplicity, and excellent standard library for file operations
- The project follows Go's idiomatic patterns and best practices

### Standard Library Components
- `os` and `io/fs` packages for file operations
- `path/filepath` for path manipulation and directory traversal
- `flag` package for command-line argument parsing
- `crypto/sha256` for file content hashing
- `regexp` for pattern matching in file names
- `bytes` for efficient byte comparison
- `fmt` for formatted output and error messages
- `testing/fstest` for in-memory filesystem testing

### Script Library
- Limited use of the [script](https://github.com/bitfield/script) library
- Primarily relying on standard Go libraries for clarity and control

## Development Setup

### Project Structure
```
deduper/
├── main.go           # Entry point and CLI handling
├── scanner.go        # Directory traversal and file discovery
├── matcher.go        # File content comparison logic
├── handler.go        # File movement and archiving
├── logger.go         # Logging and user feedback
├── scanner_test.go   # Tests for scanner component
├── matcher_test.go   # Tests for matcher component
├── handler_test.go   # Tests for handler component
├── test_data/        # Test fixtures
└── go.mod, go.sum    # Go module files
```

### Build Process
- Standard Go build process: `go build`
- No special build flags or configurations required
- Produces a single binary executable

### Testing Approach
- Test-driven development with tests written before implementation
- Unit tests with proper naming (compatible with gotestdox)
- Test fixtures in the test_data directory
- Table-driven tests for comprehensive coverage
- In-memory filesystem testing using testing/fstest
- Integration tests for end-to-end validation

## Technical Constraints

### File System Compatibility
- Works across different file systems (ext4, NTFS, APFS, etc.)
- Handles file path differences between operating systems using filepath package
- Considers case sensitivity issues in file paths

### Performance Considerations
- Efficient handling of large music libraries
- Minimizes memory usage when comparing file contents
- Uses byte-by-byte comparison for small files and hashing for larger files

### Error Handling
- Robust error handling for file system operations
- Descriptive error messages with context using fmt.Errorf and %w
- Graceful degradation when permissions or other issues arise
- Clear error messages for user troubleshooting

## Dependencies

### Direct Dependencies
- Go standard library (primary dependency)
- [script](https://github.com/bitfield/script) library (limited use)

### Indirect Dependencies
- Dependencies of the script library

## Tool Usage Patterns

### Command-Line Interface
```
deduper [flags] <directory>
```

Flags:
- `--dry-run`: Preview actions without performing file moves
- `--verbose`: Show detailed information about all actions
- `--archive-dir`: Specify the directory for archived duplicates (default: "./archive")

### File Scanning Pattern
- Recursive directory traversal using filepath.Walk
- Focus on artist/album/song hierarchy
- Identify potential duplicates based on naming patterns using regular expressions
- Group files by their original name (without the number suffix)

### Content Matching Pattern
- Byte-by-byte comparison using bytes.Equal for exact matching
- SHA-256 hashing available for efficient comparison of larger files
- Conservative approach to duplicate identification

### File Handling Pattern
- Create mirror directory structure in archive
- Move (not copy) duplicate files to preserve space
- Maintain original metadata (timestamps, permissions)
- Support for dry-run mode to preview actions

### Logging Pattern
- Different verbosity levels (info, verbose, error)
- Always show important messages regardless of verbosity
- Show detailed information only in verbose mode
- Format error messages for clarity

## Code Organization

### Flat File Structure
- All Go files are in the root directory (no subfolders)
- Functionality is separated into different files for clarity
- Files are organized by functional area

### Component-Based Architecture
- Each major function is encapsulated in its own component
- Components have clear responsibilities and interfaces
- This allows for easy testing and maintenance

### Functional Areas
- main.go: Entry point and CLI handling
- scanner.go: Directory traversal and file discovery
- matcher.go: File content comparison logic
- handler.go: File movement and archiving
- logger.go: Logging and user feedback
- *_test.go: Tests for each component

### Future Refactoring Considerations
- If the codebase grows significantly, it may be reorganized into packages
- For now, the flat structure keeps things simple while still maintaining separation of concerns
