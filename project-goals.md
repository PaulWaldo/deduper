# Deduper goals

## Problem Statement

My music library contains many duplicated songs.  There are too many to remove by hand so I want to create a program that will do this for me.  I want it to be very conservative and only remove the ones I am sure are *strict* duplicates

## General Approach

* Music is organized by artist and albums, with a structure similar to below.  Note this is a set of test data files we can use for testing the application

```shell
❯ tree test_data
test_data
└── artist1
    ├── album1
    │   ├── song_a 1.txt
    │   ├── song_a 2.txt
    │   └── song_a.txt
    └── album2

4 directories, 3 files
``

* `song_a` is the original and `song_a 1` and `song_a 2` are duplicates because the OS has appended a number.  `song_a` is the only song that should be in that folder
* A song is considered a duplicate if the file contents are *exactly* the same
* An album will have multiple songs
* To be conservative, files will not be deleted, but moved to an archive directory, keeping the same folder structure

## Implementation details
* The archive directory should be allowed to be specified when invoking the program
* A "verbose" option should be allowed to show all the actions and duplicate determination
* A "dry-run" option should be allowed that will show what would happen, but not perform any file moves
* If songs are discovered that have the duplicate naming scheme, but are not *exactly* identical, this should be shown to the user, regardless of verbosity mode.  This will indicate that these files need to be investigated manually.

## Technical Details
* Implement in go
* Create detailed unit tests and name them properly.  Names should be checked with `gotestdox`.  See [gotestdox](https://github.com/bitfield/gotestdox) for more information
* Document well the code, making sure to use go-doc comments
* I want to experiment with [GitHub - bitfield/script: Making it easy to write shell-like scripts in Go](https://github.com/bitfield/script).  Try using that to start, but if it becomes to unwieldy, we will switch back to standard go code
* If you have questions, ask me!
