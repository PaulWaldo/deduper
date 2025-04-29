# System Patterns: Deduper

## System Architecture
Deduper follows a straightforward command-line utility architecture with the following components:

1. **Command-Line Interface**: Handles user input, flags, and configuration
2. **File Scanner**: Traverses the directory structure to identify potential duplicates
3. **Content Matcher**: Compares file contents to confirm exact duplicates
4. **File Handler**: Manages the moving of duplicate files to the archive directory
5. **Logger**: Provides feedback to the user based on verbosity settings

## Key Technical Decisions

### Single Command Structure
- Deduper uses a single command with flags rather than subcommands
- This simplifies the user experience and aligns with the "simple and intuitive" UX goal
- The program automatically performs scan → match → archive in one operation

### Standard Go Libraries
- Primarily using standard Go libraries for file operations and path handling
- Using the `filepath` package for cross-platform path handling
- Using `os` package for file operations
- Using `regexp` for pattern matching in file names
- Limited use of the `script` library, focusing on standard Go code for clarity

### Conservative Duplicate Handling
- Files are moved to an archive directory rather than deleted
- Original folder structure is preserved in the archive
- This approach provides a safety net for users

### Exact Content Matching
- Duplicates are identified by comparing the exact file contents
- This is more reliable than using file size or other metadata
- Using byte-by-byte comparison for exact matching
- SHA-256 hashing available for efficient comparison of larger files

## Design Patterns

### Component-Based Architecture
- Each major function is encapsulated in its own component
- Components have clear responsibilities and interfaces
- This allows for easy testing and maintenance

### Strategy Pattern
- Different strategies for file matching and handling can be implemented
- Currently focusing on exact content matching, but could be extended

### Observer Pattern
- Logger observes the operations and reports based on verbosity settings
- This decouples the logging from the core operations

## Component Relationships

```
[CLI] → [Scanner] → [Matcher] → [Handler]
   ↓        ↓          ↓          ↓
   └────────┴──────────┴──────────┘
                  ↓
              [Logger]
```

- **CLI** parses arguments and initializes the process
- **Scanner** identifies potential duplicates based on naming patterns
- **Matcher** confirms duplicates through content comparison
- **Handler** moves confirmed duplicates to the archive
- **Logger** provides feedback throughout the process

## Critical Implementation Paths

### Duplicate Identification Path
1. Scan directory structure following artist/album/song hierarchy
2. Identify groups of files with naming patterns suggesting duplicates
3. Compare file contents within each group
4. Flag confirmed duplicates for handling

### File Handling Path
1. Create archive directory structure mirroring the original
2. Move duplicate files to the corresponding archive location
3. Verify successful move before proceeding to the next file
4. Log actions based on verbosity settings

### Edge Case Handling Path
1. Identify files with duplicate naming patterns but different content
2. Flag these for user attention regardless of verbosity settings
3. Provide clear information about the discrepancy
4. Skip these files in the automated handling process

## Implementation Details

### Scanner Component
- Uses regular expressions to identify potential duplicates based on naming patterns
- Groups files by their base name (without the number suffix)
- Returns a list of scan results, each containing an original file and potential duplicates

### Matcher Component
- Compares file contents to confirm exact duplicates
- Uses byte-by-byte comparison for exact matching
- Provides SHA-256 hashing for efficient comparison of larger files
- Supports both real filesystem and in-memory filesystem for testing

### Handler Component
- Moves confirmed duplicates to the archive directory
- Preserves the original directory structure in the archive
- Supports dry-run mode to preview actions without performing file moves
- Logs all actions based on verbosity settings

### Logger Component
- Provides different logging levels (info, verbose, error)
- Controls output based on verbosity settings
- Ensures important messages are always shown
- Formats error messages for clarity
