# Deduper

A Go utility for identifying and safely handling duplicate music files in your music library.

## Overview

Deduper is designed to solve the common problem of duplicate music files accumulating in a music library. These duplicates often occur when the operating system appends numbers to filenames during file operations (e.g., copying, downloading). Deduper identifies these duplicates based on exact file content matching and moves them to an archive directory, preserving the original folder structure.

## Features

- **Exact Content Matching**: Identifies duplicates based on byte-by-byte comparison of file contents
- **Safe Handling**: Moves duplicates to an archive directory rather than deleting them
- **Preserves Structure**: Maintains the original directory structure in the archive
- **Dry Run Mode**: Preview actions without performing file moves
- **Verbose Output**: Detailed logging of all actions
- **Edge Case Detection**: Alerts about files with duplicate naming patterns that aren't exact content matches

## Installation

### Prerequisites

- Go 1.16 or higher

### Building from Source

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/deduper.git
   cd deduper
   ```

2. Build the binary:
   ```
   go build
   ```

3. (Optional) Install the binary to your GOPATH:
   ```
   go install
   ```

## Usage

```
deduper [flags] <directory>
```

### Flags

- `--dry-run`: Preview actions without performing file moves
- `--verbose`: Show detailed information about all actions
- `--archive-dir`: Specify the directory for archived duplicates (default: "./archive")

### Examples

#### Basic Usage

Scan a directory and move duplicates to the default archive directory:

```
deduper ~/Music
```

#### Dry Run

Preview what would happen without making any changes:

```
deduper --dry-run ~/Music
```

#### Verbose Output

Get detailed information about all actions:

```
deduper --verbose ~/Music
```

#### Custom Archive Directory

Specify a custom archive directory:

```
deduper --archive-dir=/path/to/archive ~/Music
```

#### Combining Flags

You can combine multiple flags:

```
deduper --dry-run --verbose --archive-dir=/path/to/archive ~/Music
```

## How It Works

1. **File Scanning**: The program traverses a music library organized in an artist/album/song hierarchy.
2. **Duplicate Identification**: It identifies duplicates based on exact file content matching, focusing on files with naming patterns like "song.txt", "song 1.txt", "song 2.txt".
3. **Conservative Approach**: The program assumes the file without a number suffix is the original and should be kept.
4. **Safe Handling**: Instead of deleting duplicates, it moves them to an archive directory while preserving the original folder structure.

## Project Structure

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
└── test_data/        # Test fixtures
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
