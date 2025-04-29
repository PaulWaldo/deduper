# Technical Context: Deduper

## Technologies Used

### Go Programming Language
- Deduper is implemented in Go (Golang)
- Go was chosen for its performance, simplicity, and excellent standard library for file operations
- The project follows Go's idiomatic patterns and best practices

### Script Library
- Experimenting with the [script](https://github.com/bitfield/script) library for shell-like operations
- This library provides a fluent API for common operations like reading files, executing commands, and processing text
- Example usage: `script.File("file.txt").Match("pattern").CountLines()`

### Standard Library Components
- `os` and `io` packages for file operations
- `path/filepath` for path manipulation and directory traversal
- `flag` package for command-line argument parsing
- `crypto` packages for file content hashing and comparison
- `log` package for structured logging

## Development Setup

### Project Structure
```
deduper/
├── main.go           # Entry point and CLI handling
├── scanner.go        # Directory traversal and file discovery
├── matcher.go        # File content comparison logic
├── handler.go        # File movement and archiving
├── logger.go         # Logging and user feedback
├── test_data/        # Test fixtures
└── go.mod, go.sum    # Go module files
```

### Build Process
- Standard Go build process: `go build`
- No special build flags or configurations required
- Produces a single binary executable

### Testing Approach
- Unit tests with proper naming (compatible with gotestdox)
- Test fixtures in the test_data directory
- Table-driven tests for comprehensive coverage
- Integration tests for end-to-end validation

## Technical Constraints

### File System Compatibility
- Must work across different file systems (ext4, NTFS, APFS, etc.)
- Handle file path differences between operating systems
- Consider case sensitivity issues in file paths

### Performance Considerations
- Efficient handling of large music libraries
- Minimize memory usage when comparing file contents
- Use streaming approaches where possible to handle large files

### Error Handling
- Robust error handling for file system operations
- Graceful degradation when permissions or other issues arise
- Clear error messages for user troubleshooting

## Dependencies

### Direct Dependencies
- Go standard library
- [script](https://github.com/bitfield/script) library for shell-like operations

### Indirect Dependencies
- None currently identified

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
- Recursive directory traversal
- Focus on artist/album/song hierarchy
- Identify potential duplicates based on naming patterns

### Content Matching Pattern
- Hash-based comparison for efficiency
- Byte-by-byte comparison for confirmation when needed
- Conservative approach to duplicate identification

### File Handling Pattern
- Create mirror directory structure in archive
- Move (not copy) duplicate files to preserve space
- Maintain original metadata (timestamps, permissions)

## Code Organization

### Flat File Structure
- All Go files are in the root directory (no subfolders)
- Functionality is separated into different files for clarity
- Files are organized by functional area

### Functional Areas
- main.go: Entry point and CLI handling
- scanner.go: Directory traversal and file discovery
- matcher.go: File content comparison logic
- handler.go: File movement and archiving
- logger.go: Logging and user feedback

### Future Refactoring Considerations
- If the codebase grows significantly, it may be reorganized into packages
- For now, the flat structure keeps things simple while still maintaining separation of concerns
