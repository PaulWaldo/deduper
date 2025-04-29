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

### Use of the `script` Library
- Experimenting with the [script](https://github.com/bitfield/script) library for shell-like operations
- This library simplifies file operations, command execution, and output handling
- If it becomes unwieldy, we may revert to standard Go code

### Conservative Duplicate Handling
- Files are moved to an archive directory rather than deleted
- Original folder structure is preserved in the archive
- This approach provides a safety net for users

### Exact Content Matching
- Duplicates are identified by comparing the exact file contents
- This is more reliable than using file size or other metadata
- SHA-256 or similar hashing will be used for efficient comparison

## Design Patterns

### Command Pattern
- Different operations (scan, match, move) are encapsulated as commands
- This allows for dry-run mode to preview actions without execution

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
