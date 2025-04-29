# Project Brief: Deduper

## Overview
Deduper is a Go-based utility designed to identify and handle duplicate music files in a music library. The program aims to be conservative in its approach, only removing files that are confirmed to be strict duplicates.

## Core Requirements
1. Identify duplicate music files based on exact file content matching
2. Move duplicates to an archive directory (preserving folder structure) rather than deleting them
3. Support a verbose mode to show all actions and duplicate determinations
4. Support a dry-run mode to preview actions without performing file moves
5. Alert users about files with duplicate naming patterns that aren't exact content matches

## Technical Requirements
1. Implement in Go programming language
2. Create detailed unit tests with proper naming (compatible with gotestdox)
3. Document code thoroughly with go-doc comments
4. Experiment with the [script](https://github.com/bitfield/script) library for shell-like operations

## Project Scope
The scope is limited to identifying and handling duplicate music files where:
- Files are organized in an artist/album/song hierarchy
- Duplicates follow a naming pattern where the OS has appended a number (e.g., "song_a.txt", "song_a 1.txt", "song_a 2.txt")
- Only exact content matches are considered duplicates
- The original file (without the appended number) is kept, and duplicates are archived
