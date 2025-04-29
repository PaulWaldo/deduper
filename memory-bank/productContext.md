# Product Context: Deduper

## Why This Project Exists
Deduper was created to solve the common problem of duplicate music files accumulating in a music library. These duplicates often occur when the operating system appends numbers to filenames during file operations (e.g., copying, downloading). Managing these duplicates manually becomes impractical as the library grows, necessitating an automated solution.

## Problems It Solves
1. **Time Consumption**: Manually identifying and removing duplicate music files is time-consuming and error-prone.
2. **Storage Waste**: Duplicate files consume unnecessary storage space.
3. **Library Management**: Duplicates can clutter music applications and playlists, creating a suboptimal user experience.
4. **Risk of Data Loss**: Manual deletion risks removing the wrong files or versions.

## How It Should Work
1. **File Scanning**: The program traverses a music library organized in an artist/album/song hierarchy.
2. **Duplicate Identification**: It identifies duplicates based on exact file content matching, focusing on files with naming patterns like "song.txt", "song 1.txt", "song 2.txt".
3. **Conservative Approach**: The program assumes the file without a number suffix is the original and should be kept.
4. **Safe Handling**: Instead of deleting duplicates, it moves them to an archive directory while preserving the original folder structure.
5. **User Control**: The program offers options for dry-run (preview without action) and verbose output (detailed logging).
6. **Edge Case Handling**: Files with duplicate naming patterns but different content are flagged for manual review.

## User Experience Goals
1. **Confidence**: Users should feel confident that the program will only move files that are truly duplicates.
2. **Control**: Users should have control over the process through command-line options.
3. **Transparency**: The program should clearly communicate what it's doing, especially in verbose mode.
4. **Safety**: By moving files to an archive rather than deleting them, users have a safety net if mistakes occur.
5. **Efficiency**: The program should handle large music libraries efficiently.
6. **Simplicity**: The command-line interface should be straightforward and intuitive.

## Target Users
- Music enthusiasts with large digital music collections
- System administrators managing shared music libraries
- Anyone who values an organized, duplicate-free music collection
